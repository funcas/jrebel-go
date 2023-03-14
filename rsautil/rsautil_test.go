package rsautil

import "testing"

func TestPrivateKeyFrom64(t *testing.T) {
	from64, err := PrivateKeyFrom64(key)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(from64)
}
