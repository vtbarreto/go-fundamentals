package arquitetura

import (
	"runtime"
	"testing"
)

func TestDepedente(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Nao funciona em arquitetura amd64")
	}

	t.Fail()
}
