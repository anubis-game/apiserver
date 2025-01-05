package stream

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func Test_Stream_Verify_verTim(t *testing.T) {
	testCases := []struct {
		now time.Time
		mes time.Time
		err error
	}{
		// Case 000
		{
			now: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			mes: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			err: nil,
		},
		// Case 001
		{
			now: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			mes: time.Date(2024, 12, 11, 9, 59, 30, 0, time.UTC), // -30 seconds
			err: nil,
		},
		// Case 002
		{
			now: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			mes: time.Date(2024, 12, 11, 9, 59, 0, 0, time.UTC), // -60 seconds
			err: nil,
		},
		// Case 003
		{
			now: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			mes: time.Date(2024, 12, 11, 10, 0, 30, 0, time.UTC), // +30 seconds
			err: nil,
		},
		// Case 004
		{
			now: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			mes: time.Date(2024, 12, 11, 10, 1, 00, 0, time.UTC), // +60 seconds
			err: nil,
		},
		// Case 005
		{
			now: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			mes: time.Date(2024, 12, 11, 9, 58, 59, 0, time.UTC), // -61 seconds
			err: signatureTimeInvalidError,
		},
		// Case 006
		{
			now: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			mes: time.Date(2024, 12, 11, 10, 1, 01, 0, time.UTC), // +61 seconds
			err: signatureTimeInvalidError,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			err := verTim(tc.now, tc.mes)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error %v, got %v", tc.err, err)
			}
		})
	}
}
