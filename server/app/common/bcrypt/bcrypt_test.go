package bcrypt

import "testing"

func TestBcrypt(t *testing.T) {

	actualValue := "something"
	hashed, err := Hash(actualValue)

	if err != nil {
		t.Error(err)
	}

	if !Compare(hashed, actualValue) {
		t.Error("error")
	}
}
