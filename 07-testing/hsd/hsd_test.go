package hsd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestStringDistance(t *testing.T) {
	got := StringDistance("foo", "foh")

	want := 1

	if got != want {
		t.Fatalf("StringDistance: got %d, want %d", got, want)
	}
}

func TestStringDistanceWithTableDrivenTests(t *testing.T) {
	tests := []struct {
		name string
		lhs  string
		rhs  string
		want int
	}{
		{"foo and foh", "foo", "foh", 1},
		{"foo and foo", "foo", "foo", 0},
		{"foo and bar", "foo", "bar", 3},
		{"foo and foobar", "foo", "foobar", -1},
		// ...
		{"foo and empty", "foo", "", -1},
	}

	for _, tc := range tests {
		got := StringDistance(tc.lhs, tc.rhs)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("StringDistance(%s, %s): got %d, want %d", tc.lhs, tc.rhs, got, tc.want)
		}
	}
}

func TestReadData(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skipf("only works on linux %s", runtime.GOOS)
	}
}

func setup() error {
	log.Println("setup ... ")
	return nil
}

func teardown() error {
	log.Println("teardown ... ")
	return nil
}

func TestMain(m *testing.M) {
	log.Println("Before tests")

	if err := setup(); err != nil {
		log.Fatalf("setup failed: %s", err)
	}

	ret := m.Run()

	if err := teardown(); err != nil {
		log.Fatalf("teardown failed: %s", err)
	}

	log.Println("After tests")
	os.Exit(ret)
}

func TestShort(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	t.Log("Running short test")
}

func TestCreateProfile(t *testing.T) {
	dir := t.TempDir()
	// 自動で作成されて、テスト終了後に削除される
	filename := filepath.Join(dir, "profile.json")
	got, err := CreateProfile(filename)

	if err != nil {
		t.Fatalf("CreateProfile(%s): %s", filename, err)
	}
	want := false

	if got != want {
		t.Fatalf("CreateProfile(%s): got %t, want %t", filename, got, want)
	}
}

func TestEnv(t *testing.T) {
	t.Setenv("FOO", "foo")
}

func FuzzDoSomething(f *testing.F) {
	// 想定外の入力を与える

	f.Add("foo&&&")
	f.Fuzz(func(t *testing.T, s string) {
		// 何かしらの処理
		fmt.Println(s)
	})
}
