package survive

import "testing"

func TestYes(t *testing.T) {
	got := Yes("Avi")
	if got != "Avi" {
		t.Error("got", got, "want", "Avi")
	}
}
func BenchmarkYes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Yes("Avi")
	}
}
