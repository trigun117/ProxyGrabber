package code

import (
	"sync"
	"testing"
)

func TestCheckProxySOCKS(t *testing.T) {

	var wg sync.WaitGroup

	ch := make(chan QR, 1)

	wg.Add(1)
	checkProxySOCKS(`000.00.000.000:0000`, ch, &wg)

	r := <-ch

	wg.Wait()

	if r.Res != false {
		t.Fail()
	}

	wg.Add(1)
	checkProxySOCKS(`95.110.227.5:43106`, ch, &wg)

	r = <-ch

	wg.Wait()

	if r.Res != true {
		t.Fail()
	}

	close(ch)
}
