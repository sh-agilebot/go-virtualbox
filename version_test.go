package virtualbox

import (
	"context"
	"errors"
	"testing"

	"github.com/go-test/deep"
)

func TestVersion(t *testing.T) {
	testCases := map[string]struct {
		want string
		err  error
	}{
		"good": {
			want: "7.0.8r156879",
			err:  nil,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := newTestManager()

			got, err := m.Version(context.Background())
			if diff := deep.Equal(got, tc.want); !errors.Is(err, tc.err) || diff != nil {
				t.Errorf("Version = %+v, %v; want %v, %v; diff = %v",
					got, err, tc.want, tc.err, diff)
			}
		})
	}
}
