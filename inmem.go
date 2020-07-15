package retry

import "time"

type InMemManager struct {
	data             map[string]int
	retryBackOffTime int
	maximumRetryCount int
}

func (m *InMemManager) GetRetryCount(key string) int {
	if value, isExist := m.data[key]; isExist {
		return value
	}

	return 0
}

func (m *InMemManager) AddRetryCount(key string) int {
	value, isExist := m.data[key]
	if !isExist {
		value = 0
	}

	m.data[key] = value + 1
	return value + 1
}

func (m *InMemManager) ClearRetryCount(key string) {
	delete(m.data, key)
}

func (m *InMemManager) DelayProcessFollowBackOffTime(key string) {
	retryCount := m.GetRetryCount(key)
	time.Sleep(time.Duration(m.retryBackOffTime*retryCount) * time.Second)
}

func (m *InMemManager) IsMaximumRetry(key string) bool {
	return m.maximumRetryCount == m.GetRetryCount(key)
}
