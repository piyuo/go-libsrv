package server

import (
	"os"
	"testing"

	"github.com/piyuo/libsrv/log"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	//	shutdown()
	os.Exit(code)
}

func setup() {
	log.TestMode = true
}

func TestClean(t *testing.T) {
}