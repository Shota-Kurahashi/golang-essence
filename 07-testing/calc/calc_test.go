package calc

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"1 + 1", 1, 1, 2},
		{"1 + 2", 1, 2, 3},

		// ...

		{"1 + 3", 1, 3, 4},
	}

	for _, test := range tests {
		// これがないと、参照先が一番最後になる可能性がある。goroutineのfor文と同じ
		test := test

		t.Run(test.name, func(t *testing.T) {
			// test.Parallel()は、テストを並列実行する
			t.Parallel()

			got := Add(test.a, test.b)

			if got != test.want {
				t.Errorf("Add(%d, %d): got %d, want %d", test.a, test.b, got, test.want)
			}
		})

	}
}

func TestDoSomething(t *testing.T) {
	fns, err := filepath.Glob("testdata/*.json")

	if err != nil {
		t.Fatalf("failed to glob: %s", err)
	}

	for _, fn := range fns {
		t.Log(fn)
		b, err := os.ReadFile(fn)

		if err != nil {
			t.Fatalf("failed to read file: %s", err)
		}

		want := string(b)

		// cmpを使用すると、Diffが見やすくなる
		if diff := cmp.Diff(want, ""); diff != "" {
			t.Errorf("DoSomething(%s): mismatch (-want +got):\n%s", fn, diff)
		}
	}

}
