package davinci

import (
	"encoding/json"
	"fmt"
	"time"
)

type EpochTime struct {
	time.Time
}

func NewEpochTime(v int64) *EpochTime {
	this := EpochTime{
		Time: time.UnixMilli(v),
	}
	return &this
}

func (u EpochTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", (u.Time.UnixMilli()))), nil
}

func (u *EpochTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	u.Time = time.UnixMilli(timestamp)
	return nil
}
