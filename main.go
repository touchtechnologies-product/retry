package retry

import (
	"errors"
)

// Manager Manage retry
type Manager interface {
	GetRetryCount(key string) int
	AddRetryCount(key string) int
	ClearRetryCount(key string)
	DelayProcessFollowBackOffTime(key string)
	IsMaximumRetry(key string) bool
}

// InMemType In memory type
const InMemType = "inmem"

// NewManager make new manager
func NewManager(mtype string, backOffTime int, maxRetry int) (m Manager, err error) {
	switch mtype {
	case "inmem":
		m := InMemManager{
			data:             map[string]int{},
			retryBackOffTime: backOffTime,
			maximumRetyCount: maxRetry,
		}

		return &m, nil
	default:
		return m, errors.New("Invalid type")
	}
}
