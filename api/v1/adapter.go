package api

import (
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func Time(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func AdapterTime(t *google_protobuf.Timestamp) *time.Time {
	if t == nil {
		return nil
	}
	t0 := time.Unix(t.GetSeconds(), int64(t.GetNanos())).In(time.Local)
	return &t0
}

func AdapterProtoTime(t *time.Time) *google_protobuf.Timestamp {
	if t == nil {
		return nil
	}
	return &google_protobuf.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
