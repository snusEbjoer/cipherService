package encodeutils_test

import (
	"fmt"
	"testing"

	"github.com/snusEbjoer/cipherService/internal/utils/encodeutils"
)

type Case struct {
	in  string
	out string
}

func TestNewSHA256(t *testing.T) {
	cases := [...]Case{
		{
			in:  "tati",
			out: "e1b22fabee80ded9b5fad0d0366d9f9caee9f26bc02db007daab17ea43fe4483",
		},
		{
			in:  "shawty",
			out: "b4decc098c0c7aa9f9e0edcad9876b15eb49ff1e292364032d854edecd7112fb",
		},
		{
			in:  "hood",
			out: "2314555c88c801b12cde83c91a706a3f925526711f6bc5bd77dd04119f28a84e",
		},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprintf("SHA256, case: %d", i), func(t *testing.T) {
			result := encodeutils.NewSHA256(tt.in).ToString()

			if result != tt.out {
				t.Errorf("want %s, got %s", tt.out, result)
			}
		})
	}
}

func TestNewMD5(t *testing.T) {
	cases := [...]Case{
		{
			in:  "tati",
			out: "d6a9b920af25b1d240105bec4efe9c81",
		},
		{
			in:  "shawty",
			out: "bbed54b0a12ce0a65d9e2c5534b40e6f",
		},
		{
			in:  "hood",
			out: "14da543530c3e30213c24915b346caa3",
		},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprintf("SHA256, case: %d", i), func(t *testing.T) {
			result := encodeutils.NewMD5(tt.in).ToString()

			if result != tt.out {
				t.Errorf("want %s, got %s", tt.out, result)
			}
		})
	}
}
