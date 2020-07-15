package retry

import "time"

// InMemManager Retry manager which record retry count in memory
type InMemManager struct {
	data             map[string]int
	retryBackOffTime int
	maximumRetyCount int
}

// GetRetryCount Get retry count
func (m *InMemManager) GetRetryCount(key string) int {
	if value, isExist := m.data[key]; isExist {
		return value
	}

	return 0
}

// AddRetryCount Add retry count
func (m *InMemManager) AddRetryCount(key string) int {
	value, isExist := m.data[key]
	if !isExist {
		value = 0
	}

	m.data[key] = value + 1
	return value + 1
}

// ClearRetryCount Clear retry count
func (m *InMemManager) ClearRetryCount(key string) {
	delete(m.data, key)
}

// DelayProcessFollowBackOffTime Sleep for (back off time * retry count) seconds
func (m *InMemManager) DelayProcessFollowBackOffTime(key string) {
	retryCount := m.GetRetryCount(key)
	time.Sleep(time.Duration(m.retryBackOffTime*retryCount) * time.Second)
}

// IsMaximumRetry return true if retry count is equal to maximum, false otherwise
func (m *InMemManager) IsMaximumRetry(key string) bool {
	return m.maximumRetyCount == m.GetRetryCount(key)
}
