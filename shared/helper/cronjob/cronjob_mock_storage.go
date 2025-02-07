package cronjob

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// MockStorage implements CronJobStorage interface for testing
type MockStorage struct {
	mu    sync.RWMutex
	jobs  map[FuncID]CronJobState
	error error // Used to simulate storage errors
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		jobs: make(map[FuncID]CronJobState),
	}
}

func (m *MockStorage) SetError(err error) {
	m.error = err
}

func (m *MockStorage) SaveJob(ctx context.Context, state CronJobState) error {
	if m.error != nil {
		return m.error
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.jobs[state.FuncID] = state
	return nil
}

func (m *MockStorage) DeleteJob(ctx context.Context, funcID FuncID) error {
	if m.error != nil {
		return m.error
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.jobs, funcID)
	return nil
}

func (m *MockStorage) GetAllJobs(ctx context.Context) ([]CronJobState, error) {
	if m.error != nil {
		return nil, m.error
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	jobs := make([]CronJobState, 0, len(m.jobs))
	for _, job := range m.jobs {
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (m *MockStorage) GetJob(ctx context.Context, funcID FuncID) (*CronJobState, error) {
	if m.error != nil {
		return nil, m.error
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	if job, exists := m.jobs[funcID]; exists {
		return &job, nil
	}
	return nil, fmt.Errorf("job not found")
}

func (m *MockStorage) UpdateJobStatus(ctx context.Context, funcID FuncID, status CronJobStatus, errorMsg string, statusDate *time.Time) error {
	if m.error != nil {
		return m.error
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if job, exists := m.jobs[funcID]; exists {
		job.Status = status
		job.ErrorMessage = errorMsg
		switch status {
		case Missed:
			job.MissedDate = statusDate
		case Cancelled:
			job.CancelledDate = statusDate
		}
		m.jobs[funcID] = job
		return nil
	}
	return fmt.Errorf("job not found")
}
