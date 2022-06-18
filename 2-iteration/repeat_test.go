package iteration

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	assertCorrectMessage(t, repeated, expected)
}

func TestRepeat2(t *testing.T) {
	repeated := Repeat2("a", 3)
	expected := "aaa"
	assertCorrectMessage(t, repeated, expected)
}

// to run benchmarks: go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func BenchmarkRepeat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat2("a", 5)  // takes 1/3rd the time
	}
}
