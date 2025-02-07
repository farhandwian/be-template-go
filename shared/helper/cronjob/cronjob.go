package cronjob

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type FuncID string
type FuncType string
type CronJobStatus string

const (
	OnSchedule CronJobStatus = "on_schedule"
	Executed   CronJobStatus = "executed"
	Cancelled  CronJobStatus = "cancelled"
	Missed     CronJobStatus = "missed"
)

// CronJobState represents the persistent state of a scheduled job
type CronJobState struct {
	FuncID        FuncID        `json:"func_id"`
	TargetDate    time.Time     `json:"target_date"`
	CreatedAt     time.Time     `json:"created_at"`
	FunctionType  FuncType      `json:"function_type"`
	Status        CronJobStatus `json:"status"`
	ErrorMessage  string        `json:"error_message,omitempty"`
	Data          []byte        `json:"data,omitempty"`
	MissedDate    *time.Time    `json:"missed_date,omitempty"`
	CancelledDate *time.Time    `json:"cancelled_date,omitempty"`
}

type CronJobStorage interface {
	SaveJob(ctx context.Context, state CronJobState) error
	DeleteJob(ctx context.Context, funcID FuncID) error
	GetAllJobs(ctx context.Context) ([]CronJobState, error)
	GetJob(ctx context.Context, funcID FuncID) (*CronJobState, error)
	UpdateJobStatus(ctx context.Context, funcID FuncID, status CronJobStatus, errorMsg string, statusDate *time.Time) error
}

// FunctionRegistry is a map of function types to their implementations
type FunctionRegistry map[FuncType]func(data []byte) error

// CronJob handles scheduled tasks with thread-safe operations
type CronJob struct {
	cronjob    *cron.Cron
	mu         sync.RWMutex
	storage    map[FuncID]cron.EntryID
	persistent CronJobStorage
	functions  FunctionRegistry
}

// NewCronJob creates a new CronJob instance with the specified timezone and storage
func NewCronJob(location *time.Location, persistentStorage CronJobStorage) *CronJob {
	if location == nil {
		location = time.Local
	}
	return &CronJob{
		cronjob:    cron.New(cron.WithSeconds(), cron.WithLocation(location)),
		storage:    make(map[FuncID]cron.EntryID),
		persistent: persistentStorage,
		functions:  make(FunctionRegistry),
	}
}

// RegisterFunction adds a named function to the registry
func (c *CronJob) RegisterFunction(functionType FuncType, f func(data []byte) error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.functions[functionType] = f
}

// Start begins the cron scheduler and recovers persisted jobs
func (c *CronJob) Start(ctx context.Context) error {
	if c.cronjob == nil {
		return fmt.Errorf("cronjob not initialized")
	}

	// Recover persisted jobs
	states, err := c.persistent.GetAllJobs(ctx)
	if err != nil {
		return fmt.Errorf("failed to recover jobs: %v", err)
	}

	now := time.Now()
	for _, state := range states {
		// Skip jobs that are already completed (executed or cancelled)
		if state.Status == Executed || state.Status == Cancelled {
			continue
		}

		// Handle jobs that were supposed to run while system was down
		if state.TargetDate.Before(now) && state.Status == OnSchedule {
			missedDate := now
			// Mark job as missed
			c.persistent.UpdateJobStatus(ctx, state.FuncID, Missed, "Target date passed while system was down", &missedDate)
			continue
		}

		// Only recover jobs that are still scheduled and have future target dates
		if state.Status == OnSchedule && state.TargetDate.After(now) {
			// Get the function from registry
			f, exists := c.functions[state.FunctionType]
			if !exists {
				log.Printf("Function type %s not registered, skipping job %s\n", state.FunctionType, state.FuncID)
				continue
			}

			// Schedule the recovered job
			if err := c.addFunc(ctx, state.TargetDate, state.FuncID, f, state.Data); err != nil {
				log.Printf("Failed to recover job %s: %v\n", state.FuncID, err)
			} else {
				log.Printf("Successfully recovered job %s scheduled for %v\n", state.FuncID, state.TargetDate)
			}
		}
	}

	c.cronjob.Start()
	return nil
}

// ExecuteLater schedules a registered function by its type
func (c *CronJob) ExecuteLater(ctx context.Context, targetDate time.Time, funcID FuncID, funcType FuncType, data []byte) error {
	c.mu.RLock()
	f, exists := c.functions[funcType]
	c.mu.RUnlock()

	if !exists {
		return fmt.Errorf("function type %s not registered", funcType)
	}

	// Validate target date
	now := time.Now()
	if targetDate.Before(now) {
		return fmt.Errorf("target date must be in the future")
	}

	// Create state for persistence
	state := CronJobState{
		FuncID:       funcID,
		TargetDate:   targetDate,
		CreatedAt:    now,
		FunctionType: funcType,
		Status:       OnSchedule,
		Data:         data,
	}

	// Save to persistent storage first
	if err := c.persistent.SaveJob(ctx, state); err != nil {
		return fmt.Errorf("failed to persist job: %v", err)
	}

	// Then schedule the job
	return c.addFunc(ctx, targetDate, funcID, f, data)
}

func (c *CronJob) addFunc(ctx context.Context, targetDate time.Time, funcID FuncID, funcToRunLater func([]byte) error, data []byte) error {
	if funcID == "" {
		return fmt.Errorf("funcID cannot be empty")
	}
	if funcToRunLater == nil {
		return fmt.Errorf("function cannot be nil")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.storage[funcID]; exists {
		return fmt.Errorf("job with funcID %s already exists", funcID)
	}

	localLocation := c.cronjob.Location()
	targetDate = targetDate.In(localLocation)
	now := time.Now().In(localLocation)

	if targetDate.Before(now) {
		return fmt.Errorf("scheduled time must be in the future")
	}

	cronExpression := fmt.Sprintf("%d %d %d %d %d *",
		targetDate.Second(),
		targetDate.Minute(),
		targetDate.Hour(),
		targetDate.Day(),
		targetDate.Month(),
	)

	safeFunc := func() {
		executionTime := time.Now()
		defer func() {
			if r := recover(); r != nil {
				errMsg := fmt.Sprintf("Recovered from panic: %v", r)
				c.persistent.UpdateJobStatus(ctx, funcID, Executed, errMsg, &executionTime)
			}
			// Remove from internal storage after execution
			c.removeFromStorage(funcID)
		}()

		// Execute the function and handle any errors
		err := funcToRunLater(data)
		if err != nil {
			c.persistent.UpdateJobStatus(ctx, funcID, Executed, err.Error(), &executionTime)
		} else {
			c.persistent.UpdateJobStatus(ctx, funcID, Executed, "", &executionTime)
		}
	}

	cid, err := c.cronjob.AddFunc(cronExpression, safeFunc)
	if err != nil {
		return fmt.Errorf("error adding cron function: %v", err)
	}

	c.storage[funcID] = cid
	return nil
}

// removeFromStorage removes a job from internal storage without updating its status
func (c *CronJob) removeFromStorage(funcID FuncID) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if cs, exist := c.storage[funcID]; exist {
		c.cronjob.Remove(cs)
		delete(c.storage, funcID)
	}
}

// CancelFunc explicitly cancels a job and marks it as cancelled
func (c *CronJob) CancelFunc(ctx context.Context, funcID FuncID) error {
	if funcID == "" {
		return fmt.Errorf("funcID cannot be empty")
	}

	// First check if job exists and has correct status
	job, err := c.persistent.GetJob(ctx, funcID)
	if err != nil {
		return fmt.Errorf("job not found: %w", err)
	}

	// Only allow cancellation of jobs that are OnSchedule
	if job.Status != OnSchedule {
		return fmt.Errorf("cannot cancel job with status %s", job.Status)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if cs, exist := c.storage[funcID]; exist {
		c.cronjob.Remove(cs)
		delete(c.storage, funcID)

		cancelledDate := time.Now()
		// Update status to cancelled
		if err := c.persistent.UpdateJobStatus(ctx, funcID, Cancelled, "Job cancelled by user", &cancelledDate); err != nil {
			return fmt.Errorf("failed to update job status: %w", err)
		}
		return nil
	}
	return fmt.Errorf("job not found in scheduler")
}

// GetJobCount returns the number of currently scheduled jobs
func (c *CronJob) GetJobCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.storage)
}

// Stop gracefully stops the cron scheduler
func (c *CronJob) Stop() {
	if c.cronjob != nil {
		// Just stop the scheduler without cancelling jobs
		c.cronjob.Stop()

		// Clear internal storage without updating job statuses
		c.mu.Lock()
		c.storage = make(map[FuncID]cron.EntryID)
		c.mu.Unlock()
	}
}
