package log

import (
	"context"
	"testing"
	"time"

	"github.com/piyuo/libsrv/src/env"
	"github.com/piyuo/libsrv/src/identifier"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var here = "log_test"

func TestLogGetHeader(t *testing.T) {
	assert := assert.New(t)
	appName = "test"
	ctx := context.Background()
	ctx = env.SetUserID(ctx, "user1")
	header, id := getHeader(ctx, here)
	assert.Equal("user1@test/log_test: ", header)
	assert.Equal("user1", id)
}

func TestLogPrint(t *testing.T) {
	ctx := context.Background()
	Print(ctx, here, "info ...")
}

//TestLog is a production test, it will write log to google cloud platform under log viewer "Google Project, project name"
func TestLog(t *testing.T) {
	ctx := context.Background()
	Info(ctx, here, "my info")
	Warning(ctx, here, "my warning")
	Alert(ctx, here, "my alert")
	TestModeAlwaySuccess()
	defer TestModeBackNormal()
	Info(ctx, here, "my info")
	Warning(ctx, here, "my warning")
	Alert(ctx, here, "my alert")
}

func TestLogWhenContextCanceled(t *testing.T) {
	assert := assert.New(t)
	dateline := time.Now().Add(time.Duration(1) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), dateline)
	defer cancel()
	time.Sleep(time.Duration(2) * time.Millisecond)
	logger, err := NewLogger(ctx)
	assert.NotNil(err)
	assert.Nil(logger)

	Log(ctx, DEBUG, here, "")
	WriteError(ctx, here, "", "", "")
	Error(ctx, here, nil)
	Info(ctx, here, "my info log canceled")

	errorer, err := NewErrorer(ctx)
	assert.NotNil(err)
	assert.Nil(errorer)
	logger, err = NewLogger(ctx)
	assert.NotNil(err)
	assert.Nil(logger)
}

func TestBeautyStack(t *testing.T) {
	assert := assert.New(t)
	err := errors.New("beauty stack")
	stack := beautyStack(err)
	assert.NotEmpty(stack)
}

func TestExtractFilename(t *testing.T) {
	assert := assert.New(t)
	path := "/convey/doc.go:75"
	filename := extractFilename(path)
	assert.Equal("doc.go:75", filename)
	path = "doc.go:75"
	filename = extractFilename(path)
	assert.Equal("doc.go:75", filename)
}

func TestIsLineUsable(t *testing.T) {
	assert := assert.New(t)
	line := "/smartystreets/convey/doc.go:75"
	assert.False(isLineUsable(line))
}

func TestIsLineDuplicate(t *testing.T) {
	assert := assert.New(t)
	list := []string{"/doc.go:75", "/doc.go:75"}
	assert.False(isLineDuplicate(list, 0))
	assert.True(isLineDuplicate(list, 1))
}

func TestError(t *testing.T) {
	ctx := context.Background()
	err := errors.New("mock error happening in go")
	Error(ctx, here, err)

	TestModeAlwaySuccess()
	defer TestModeBackNormal()
	Error(ctx, here, nil)
}

func TestErrorWithRequest(t *testing.T) {
	ctx := context.Background()
	err := errors.New("mock error happening in go with request")
	Error(ctx, here, err)
}

func TestCustomError(t *testing.T) {
	ctx := context.Background()
	message := "mock error happening in flutter"
	stack := "at firstLine (a.js:3)\nat secondLine (b.js:3)"
	id := identifier.UUID()
	WriteError(ctx, here, message, stack, id)
}
