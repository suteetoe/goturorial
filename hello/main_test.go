package hello

import "testing"

func TestSumOfFirst(t *testing.T) {
	given := 3
	want := 6
	get := sumOfFirst(given)
	if get != want {
		t.Errorf("given %d, want %d get %d",
			given, want, get)
	}
}

func TestDouble(t *testing.T) {
	given := 2
	want := 4
	double(&given)
	if given != want {
		t.Errorf("given %d, want %d",
			given, want)
	}

}

func TestAppendGreeting(t *testing.T) {
	given := "Bob"
	want := "Hi, Bob"
	appendGreeting(&given)
	if given != want {
		t.Errorf("given %s, want %s",
			given, want)
	}
}
