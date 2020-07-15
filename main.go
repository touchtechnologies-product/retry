package retry

import (
	"errors"
)

type Manager interface {
	GetRetryCount(key string) int
	AddRetryCount(key string) int
	ClearRetryCount(key string)
	DelayProcessFollowBackOffTime(key string)
	IsMaximumRetry(key string) bool
}

const InMemType = "inMem"

func NewManager(mType string, backOffTime int, maxRetry int) (m Manager, err error) {
	switch mType {
	case "inMem":
		m := InMemManager{
			data:             map[string]int{},
			retryBackOffTime: backOffTime,
			maximumRetryCount: maxRetry,
		}

		return &m, nil
	default:
		return m, errors.New("invalid type")
	}
}
