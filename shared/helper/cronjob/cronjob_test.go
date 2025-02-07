package cronjob

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"
)

type TestJob struct {
	Message string
	Done    chan bool
}

// Custom error check function
func errorContains(err error, want string) bool {
	if err == nil {
		return want == ""
	}
	if want == "" {
		return err == nil
	}
	return strings.Contains(err.Error(), want)
}

func TestCronJob_NewCronJob(t *testing.T) {
	storage := NewMockStorage()

	tests := []struct {
		name     string
		location *time.Location
		wantErr  bool
	}{
		{
			name:     "with UTC location",
			location: time.UTC,
			wantErr:  false,
		},
		{
			name:     "with nil location",
			location: nil,
			wantErr:  false,
		},
		{
			name:     "with local location",
			location: time.Local,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cron := NewCronJob(tt.location, storage)
			if cron == nil {
				t.Error("expected non-nil CronJob")
			}
			if tt.location == nil && cron.cronjob.Location() != time.UTC {
				t.Error("expected UTC location when nil provided")
			}
		})
	}
}

func TestCronJob_RegisterFunction(t *testing.T) {
	storage := NewMockStorage()
	cron := NewCronJob(time.UTC, storage)

	funcType := FuncType("test_func")
	testFunc := func(data []byte) error { return nil }

	// Register function
	cron.RegisterFunction(funcType, testFunc)

	// Verify registration
	if _, exists := cron.functions[funcType]; !exists {
		t.Error("function was not registered")
	}
}

func TestCronJob_ExecuteLater(t *testing.T) {
	storage := NewMockStorage()
	cron := NewCronJob(time.UTC, storage)

	// Register test function
	funcType := FuncType("test_func")
	testFunc := func(data []byte) error {
		var job TestJob
		if err := json.Unmarshal(data, &job); err != nil {
			return err
		}
		return nil
	}
	cron.RegisterFunction(funcType, testFunc)

	tests := []struct {
		name        string
		targetDate  time.Time
		funcID      FuncID
		funcType    FuncType
		data        []byte
		setupError  error
		wantErr     bool
		errContains string
	}{
		{
			name:       "valid future date",
			targetDate: time.Now().Add(time.Hour),
			funcID:     "test1",
			funcType:   funcType,
			data:       []byte(`{"Message": "test"}`),
			wantErr:    false,
		},
		{
			name:        "past date",
			targetDate:  time.Now().Add(-time.Hour),
			funcID:      "test2",
			funcType:    funcType,
			data:        []byte(`{"Message": "test"}`),
			wantErr:     true,
			errContains: "must be in the future",
		},
		{
			name:        "storage error",
			targetDate:  time.Now().Add(time.Hour),
			funcID:      "test3",
			funcType:    funcType,
			data:        []byte(`{"Message": "test"}`),
			setupError:  errors.New("storage error"),
			wantErr:     true,
			errContains: "failed to persist",
		},
		{
			name:        "unregistered function",
			targetDate:  time.Now().Add(time.Hour),
			funcID:      "test4",
			funcType:    "unknown_func",
			data:        []byte(`{"Message": "test"}`),
			wantErr:     true,
			errContains: "not registered",
		},
		{
			name:       "duplicate job ID",
			targetDate: time.Now().Add(time.Hour),
			funcID:     "test5",
			funcType:   funcType,
			data:       []byte(`{"Message": "test"}`),
			wantErr:    false, // First attempt should succeed
		},
		{
			name:        "duplicate job ID second attempt",
			targetDate:  time.Now().Add(time.Hour),
			funcID:      "test5", // Same ID as previous test
			funcType:    funcType,
			data:        []byte(`{"Message": "test"}`),
			wantErr:     true,
			errContains: "already exists",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupError != nil {
				storage.SetError(tt.setupError)
			} else {
				storage.SetError(nil)
			}

			err := cron.ExecuteLater(context.TODO(), tt.targetDate, tt.funcID, tt.funcType, tt.data)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ExecuteLater() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if tt.errContains != "" && !errorContains(err, tt.errContains) {
					t.Errorf("ExecuteLater() error = %v, want error containing %q", err, tt.errContains)
				}
			} else if err != nil {
				t.Errorf("ExecuteLater() unexpected error = %v", err)
			}
		})
	}
}

func TestCronJob_Start(t *testing.T) {
	tests := []struct {
		name       string
		setupJobs  []CronJobState
		setupError error
		wantErr    bool
	}{
		{
			name: "recover future jobs",
			setupJobs: []CronJobState{
				{
					FuncID:       "future1",
					TargetDate:   time.Now().Add(time.Hour),
					CreatedAt:    time.Now(),
					FunctionType: "test_func",
					Status:       OnSchedule,
				},
			},
			wantErr: false,
		},
		{
			name: "skip past jobs",
			setupJobs: []CronJobState{
				{
					FuncID:       "past1",
					TargetDate:   time.Now().Add(-time.Hour),
					CreatedAt:    time.Now().Add(-2 * time.Hour),
					FunctionType: "test_func",
					Status:       OnSchedule,
				},
			},
			wantErr: false,
		},
		{
			name:       "storage error",
			setupError: errors.New("storage error"),
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := NewMockStorage()
			cron := NewCronJob(time.UTC, storage)

			// Register test function
			funcType := FuncType("test_func")
			testFunc := func(data []byte) error { return nil }
			cron.RegisterFunction(funcType, testFunc)

			// Setup test state
			for _, job := range tt.setupJobs {
				if err := storage.SaveJob(context.TODO(), job); err != nil {
					t.Fatalf("Failed to setup test: %v", err)
				}
			}

			if tt.setupError != nil {
				storage.SetError(tt.setupError)
			}

			err := cron.Start(context.TODO())

			if tt.wantErr {
				if err == nil {
					t.Error("Start() error = nil, wantErr true")
				}
			} else if err != nil {
				t.Errorf("Start() unexpected error = %v", err)
			}

			// Cleanup
			cron.Stop()
		})
	}
}

func TestCronJob_CancelFunc(t *testing.T) {
	storage := NewMockStorage()
	cron := NewCronJob(time.UTC, storage)

	// Register test function
	funcType := FuncType("test_func")
	testFunc := func(data []byte) error { return nil }
	cron.RegisterFunction(funcType, testFunc)

	// Helper function to create a job with specific status
	createJob := func(id FuncID, status CronJobStatus) error {
		state := CronJobState{
			FuncID:       id,
			TargetDate:   time.Now().Add(time.Hour),
			CreatedAt:    time.Now(),
			FunctionType: funcType,
			Status:       status,
			Data:         []byte(`{"message": "test"}`),
		}
		return storage.SaveJob(context.TODO(), state)
	}

	// Setup various jobs with different statuses
	// Active job that should be cancellable
	activeJobID := FuncID("active_job")
	err := cron.ExecuteLater(context.TODO(), time.Now().Add(time.Hour), activeJobID, funcType, []byte(`{"message": "test"}`))
	if err != nil {
		t.Fatalf("Failed to schedule active job: %v", err)
	}

	// Create jobs with different statuses
	if err := createJob("executed_job", Executed); err != nil {
		t.Fatalf("Failed to create executed job: %v", err)
	}
	if err := createJob("cancelled_job", Cancelled); err != nil {
		t.Fatalf("Failed to create cancelled job: %v", err)
	}
	if err := createJob("missed_job", Missed); err != nil {
		t.Fatalf("Failed to create missed job: %v", err)
	}

	tests := []struct {
		name           string
		funcID         FuncID
		wantErr        bool
		errMsgContains string
		preCheck       func() error
	}{
		{
			name:    "cancel active job",
			funcID:  activeJobID,
			wantErr: false,
		},
		{
			name:           "cancel already executed job",
			funcID:         "executed_job",
			wantErr:        true,
			errMsgContains: "cannot cancel job with status executed",
		},
		{
			name:           "cancel already cancelled job",
			funcID:         "cancelled_job",
			wantErr:        true,
			errMsgContains: "cannot cancel job with status cancelled",
		},
		{
			name:           "cancel missed job",
			funcID:         "missed_job",
			wantErr:        true,
			errMsgContains: "cannot cancel job with status missed",
		},
		{
			name:           "cancel non-existent job",
			funcID:         "nonexistent",
			wantErr:        true,
			errMsgContains: "job not found",
		},
		{
			name:           "cancel with empty funcID",
			funcID:         "",
			wantErr:        true,
			errMsgContains: "funcID cannot be empty",
		},
		{
			name:           "cancel with storage error",
			funcID:         activeJobID,
			wantErr:        true,
			errMsgContains: "storage error",
			preCheck: func() error {
				storage.SetError(errors.New("storage error"))
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset storage error before each test
			storage.SetError(nil)

			if tt.preCheck != nil {
				if err := tt.preCheck(); err != nil {
					t.Fatalf("Pre-check failed: %v", err)
				}
			}

			// Perform cancellation
			err := cron.CancelFunc(context.TODO(), tt.funcID)

			// Check error expectations
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error, got nil")
				} else if tt.errMsgContains != "" && !strings.Contains(err.Error(), tt.errMsgContains) {
					t.Errorf("Error message should contain %q, got %v", tt.errMsgContains, err)
				}

				// For error cases, verify the job status hasn't changed
				if tt.funcID != "" {
					job, getErr := storage.GetJob(context.TODO(), tt.funcID)
					if getErr == nil {
						// Check that status hasn't changed for error cases
						originalStatus := job.Status
						if job.Status != originalStatus {
							t.Errorf("Job status should not have changed from %v", originalStatus)
						}
						// Check that CancelledDate is still nil for error cases
						if job.CancelledDate != nil {
							t.Error("CancelledDate should be nil for failed cancellation")
						}
					}
				}
			} else {
				// Success case
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}

				// Verify successful cancellation
				job, getErr := storage.GetJob(context.TODO(), tt.funcID)
				if getErr != nil {
					t.Errorf("Failed to get job after cancellation: %v", getErr)
				} else {
					if job.Status != Cancelled {
						t.Errorf("Job status = %v, want Cancelled", job.Status)
					}
					if job.CancelledDate == nil {
						t.Error("CancelledDate should not be nil")
					}
				}

				// Verify job is removed from internal storage
				cron.mu.RLock()
				if _, exists := cron.storage[tt.funcID]; exists {
					t.Error("Job should be removed from internal storage")
				}
				cron.mu.RUnlock()
			}
		})
	}
}

func TestCronJob_Stop(t *testing.T) {
	storage := NewMockStorage()
	cron := NewCronJob(time.UTC, storage)

	// Register test function
	funcType := FuncType("test_func")
	testFunc := func(data []byte) error { return nil }
	cron.RegisterFunction(funcType, testFunc)

	// Schedule some jobs
	futureTime := time.Now().Add(time.Hour)
	jobs := []struct {
		id   FuncID
		data []byte
	}{
		{"test1", []byte(`{"message": "test1"}`)},
		{"test2", []byte(`{"message": "test2"}`)},
	}

	// Create jobs and store their original states
	originalStates := make(map[FuncID]CronJobState)
	for _, job := range jobs {
		err := cron.ExecuteLater(context.TODO(), futureTime, job.id, funcType, job.data)
		if err != nil {
			t.Fatalf("Failed to schedule job: %v", err)
		}

		// Store original state
		jobState, err := storage.GetJob(context.TODO(), job.id)
		if err != nil {
			t.Fatalf("Failed to get job state: %v", err)
		}
		originalStates[job.id] = *jobState
	}

	// Start the scheduler
	if err := cron.Start(context.TODO()); err != nil {
		t.Fatalf("Failed to start scheduler: %v", err)
	}

	// Stop the scheduler
	cron.Stop()

	// Verify:
	// 1. Jobs still exist in storage with original status
	// 2. Internal storage is cleared

	// Check internal storage is empty
	if count := cron.GetJobCount(); count != 0 {
		t.Errorf("Job count = %d, want 0", count)
	}

	// Verify jobs in persistent storage remain unchanged
	for _, job := range jobs {
		currentState, err := storage.GetJob(context.TODO(), job.id)
		if err != nil {
			t.Errorf("Failed to get job %s after stop: %v", job.id, err)
			continue
		}

		originalState := originalStates[job.id]

		// Check status hasn't changed
		if currentState.Status != originalState.Status {
			t.Errorf("Job %s status changed from %v to %v",
				job.id, originalState.Status, currentState.Status)
		}

		// Check other fields remain unchanged
		if currentState.CancelledDate != nil {
			t.Errorf("Job %s shouldn't have CancelledDate set", job.id)
		}

		if currentState.ErrorMessage != originalState.ErrorMessage {
			t.Errorf("Job %s error message changed from %q to %q",
				job.id, originalState.ErrorMessage, currentState.ErrorMessage)
		}
	}
}

func TestCronJob_StopAndRestart(t *testing.T) {
	storage := NewMockStorage()
	cron := NewCronJob(time.UTC, storage)

	// Register test function
	funcType := FuncType("test_func")
	testFunc := func(data []byte) error { return nil }
	cron.RegisterFunction(funcType, testFunc)

	// Schedule a job
	futureTime := time.Now().Add(time.Hour)
	jobID := FuncID("test1")

	err := cron.ExecuteLater(context.TODO(), futureTime, jobID, funcType, []byte(`{"message": "test"}`))
	if err != nil {
		t.Fatalf("Failed to schedule job: %v", err)
	}

	// Start the scheduler
	if err := cron.Start(context.TODO()); err != nil {
		t.Fatalf("Failed to start scheduler: %v", err)
	}

	// Stop the scheduler
	cron.Stop()

	// Verify job status is unchanged
	jobState, err := storage.GetJob(context.TODO(), jobID)
	if err != nil {
		t.Fatalf("Failed to get job state: %v", err)
	}
	if jobState.Status != OnSchedule {
		t.Errorf("Job status = %v, want OnSchedule", jobState.Status)
	}

	// Restart the scheduler
	if err := cron.Start(context.TODO()); err != nil {
		t.Fatalf("Failed to restart scheduler: %v", err)
	}

	// Verify job is recovered
	jobState, err = storage.GetJob(context.TODO(), jobID)
	if err != nil {
		t.Fatalf("Failed to get job state after restart: %v", err)
	}
	if jobState.Status != OnSchedule {
		t.Errorf("Job status after restart = %v, want OnSchedule", jobState.Status)
	}

	// Verify job is in internal storage
	if count := cron.GetJobCount(); count != 1 {
		t.Errorf("Job count after restart = %d, want 1", count)
	}
}
