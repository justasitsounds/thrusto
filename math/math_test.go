package math

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestRamp(t *testing.T) {
	is := is.New(t)

	testCases := []struct {
		s        float64
		a        float64
		b        float64
		rampwant float64
	}{
		{
			s:        0,
			a:        0,
			b:        1,
			rampwant: 0,
		},
		{
			s:        0.5,
			a:        0,
			b:        1,
			rampwant: 0.5,
		},
		{
			s:        -0.5,
			a:        0,
			b:        1,
			rampwant: -0.5,
		},
		{
			s:        3,
			a:        0,
			b:        1,
			rampwant: 3,
		},
		{
			s:        3,
			a:        0,
			b:        10,
			rampwant: 0.3,
		},
		{
			s:        -3,
			a:        0,
			b:        10,
			rampwant: -0.3,
		},
		{
			s:        0,
			a:        -10,
			b:        10,
			rampwant: 0.5,
		},
		{
			s:        -3,
			a:        -10,
			b:        10,
			rampwant: 0.35,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Ramp(%.1f,%.1f,%.1f)", tc.s, tc.a, tc.b), func(t *testing.T) {
			is := is.New(t)
			got := Ramp(tc.s, tc.a, tc.b)
			is.Equal(got, tc.rampwant)
		})
	}

}
