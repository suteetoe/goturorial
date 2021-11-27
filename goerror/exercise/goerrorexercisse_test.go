package exercise

import "testing"

func TestValidateLength(t *testing.T) {
	given := "yourxxxxxx"
	want := true
	if get, err := validateLength(given); err != nil {
		if get != want {
			t.Errorf("Want Validate %v but %v", want, get)
		}
	}

	given = "name"
	want = false
	get, err := validateLength(given)
	if err == nil {
		t.Error("must error a too short string")
	}

	if get != want {
		t.Error("must error a too short string")
	}
}

func TestValidateAge(t *testing.T) {
	given := 16
	want := "16 is too young"

	if get := validateAge(given); get == nil || get.Error() != want {
		t.Errorf("given age %d get %s but %s",
			given, get.Error(), want)
	}
}
