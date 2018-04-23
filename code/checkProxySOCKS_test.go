package code

import (
	"testing"
)

func TestCheckProxySOCKS(t *testing.T) {

	ch := make(chan QR, 1)

	CheckProxySOCKS(`000.00.000.000:0000`, ch)

	r := <-ch

	if r.Res != false {
		t.Fail()
	}

	CheckProxySOCKS(`159.65.197.7:1080`, ch)

	r = <-ch

	if r.Res != true {
		t.Fail()
	}
}