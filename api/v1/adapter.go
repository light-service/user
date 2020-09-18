package api

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func Time(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func AdapterTime(t *timestamp.Timestamp) *time.Time {
	if t == nil {
		return nil
	}
	t0 := time.Unix(t.GetSeconds(), int64(t.GetNanos())).In(time.Local)
	return &t0
}

func AdapterProtoTime(t *time.Time) *timestamp.Timestamp {
	if t == nil {
		return nil
	}
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
