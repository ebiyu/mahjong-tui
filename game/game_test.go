package game

import (
	"reflect"
	"testing"
)

func TestSortPai(t *testing.T) {
	testcases := []struct {
		pai  []string
		want []string
	}{
		{
			[]string{"1m", "2m"},
			[]string{"1m", "2m"},
		},
		{
			[]string{"2m", "1m", "3s", "5m", "1p", "1m", "4s", "8p"},
			[]string{"1m", "1m", "2m", "5m", "1p", "8p", "3s", "4s"},
		},
		{
			[]string{"2m", "7m", "0m"},
			[]string{"2m", "0m", "7m"},
		},
		{
			[]string{"2m", "5m", "0m"},
			[]string{"2m", "5m", "0m"},
		},
		{
			[]string{"2m", "0m", "5m"},
			[]string{"2m", "5m", "0m"},
		},
	}
	for _, test := range testcases {
		if got := sortPai(test.pai); !reflect.DeepEqual(got, test.want) {
			t.Errorf("want = %v, got = %v", test.want, got)
		}
	}
}
