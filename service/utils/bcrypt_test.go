package utils_test

import (
	"testing"

	"github.com/Fachrulmustofa20/bank-example.git/service/utils"
)

func TestHasPass(t *testing.T) {
	type args struct {
		pass string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				pass: "test",
			},
			want: "test",
		},
		{
			name: "Test 2",
			args: args{
				pass: "test",
			},
			want: "test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := utils.HashPass(test.args.pass); result == test.args.pass {
				t.Errorf("result = %v, expected result %v", result, test.want)
			}
		})
	}
}

func TestComparePassword(t *testing.T) {
	type args struct {
		h []byte
		p []byte
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				h: []byte("$2a$08$8bzmv47RMjAuiJeSyumq5.iTfGKelWeBzsV..hkDcRRl6p9AnZixm"),
				p: []byte("test123"),
			},
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := utils.ComparePassword(test.args.h, test.args.p); result != test.want {
				t.Errorf("result = %v, expected result %v", result, test.want)
			}
		})
	}
}
