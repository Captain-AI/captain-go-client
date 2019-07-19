package captain

import (
	"encoding/json"
	"time"
)

const timeFormat = "2006-01-02T15:04:05.999-07:00"

type Timestamp struct {
	time.Time
}

func (t *Timestamp) MarhsalText() ([]byte, error) {
	return json.Marshal(t.Truncate(time.Millisecond).Format(timeFormat))
}

func (t *Timestamp) UnmarshalText(data []byte) error {
	ts, err := time.Parse(timeFormat, string(data))
	if err != nil {
		return err
	}
	t.Time = ts
	return nil
}
