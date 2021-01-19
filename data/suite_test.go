package data

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	//	setup()
	code := m.Run()
	//	shutdown()
	os.Exit(code)
}

func TestCleanTest(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	g, err := NewSampleGlobalDB(ctx)
	assert.Nil(err)
	defer g.Close()

	counters := g.Counters()
	counter := counters.SampleCounter()
	counter.Clear(ctx)

	coders := g.Coders()
	coder := coders.SampleCoder()
	coder.Clear(ctx)

	g.SampleTable().Clear(ctx)
}