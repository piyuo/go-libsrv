package server

import (
	"bytes"
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/piyuo/libsrv/command"
	"github.com/piyuo/libsrv/command/mock"
	"github.com/stretchr/testify/assert"
)

var textLong = `{
    "_id": "55d26da7c3f96f90aa005",
    "age": 20,
    "gender": "female",
    "company": "ZOGAK",
    "phone": "+1 (915) 479-2908"
   `

func BenchmarkServerBigArchive(b *testing.B) {
	handler := newTestServerHandler()
	actBytes := newTestAction(textLong)
	req1, _ := http.NewRequest("GET", "/", bytes.NewReader(actBytes))
	req1.Header.Set("Accept-Encoding", "gzip")
	resp1 := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(resp1, req1)
		_ = resp1.Result()
	}
}

func BenchmarkServerSmallAction(b *testing.B) {
	handler := newTestServerHandler()
	actBytes := newTestAction("Hi")
	req1, _ := http.NewRequest("GET", "/", bytes.NewReader(actBytes))
	req1.Header.Set("Accept-Encoding", "gzip")
	resp1 := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(resp1, req1)
		_ = resp1.Result()
	}
}

func newBigDataAction() (*mock.CmdBigData, []byte) {
	dispatch := &command.Dispatch{
		Map: &mock.MapXXX{},
	}
	act := &mock.CmdBigData{}
	actBytes, _ := dispatch.EncodeCommand(act.XXX_MapID(), act)
	return act, actBytes
}

func TestServerReady(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	httpHandler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		return nil
	}
	taskHandler := func(ctx context.Context, r *http.Request) error {
		return nil
	}

	server := &Server{
		CommandHandlers: map[string]command.IMap{"/cmd": &mock.MapXXX{}},
		HTTPHandlers:    map[string]HTTPHandler{"/http": httpHandler},
		TaskHandlers:    map[string]TaskHandler{"/task": taskHandler},
	}
	port := server.ready(context.Background())
	assert.Equal(":8080", port)

	//cleanup http.Handle mapping
	http.DefaultServeMux = new(http.ServeMux)

	//test empty PORT
	os.Setenv("PORT", "")
	port = server.ready(context.Background())
	assert.Equal(":8080", port)
	os.Setenv("PORT", "8080")
}

func TestServerArchive(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	handler := newTestServerHandler()

	act, actBytes := newBigDataAction()
	sampleBytes := []byte(act.GetSample())
	sampleLen := len(sampleBytes)
	req1, _ := http.NewRequest("GET", "/", bytes.NewReader(actBytes))
	req1.Header.Set("Accept-Encoding", "gzip")
	resp1 := httptest.NewRecorder()
	handler.ServeHTTP(resp1, req1)
	res1 := resp1.Result()
	returnBytes := resp1.Body.Bytes()
	returnLen := len(returnBytes)
	assert.Equal(200, res1.StatusCode)
	assert.Greater(returnLen, 10)
	assert.Greater(sampleLen, returnLen)
	assert.Equal("gzip", res1.Header.Get("Content-Encoding"))
	assert.Equal("application/octet-stream", res1.Header.Get("Content-Type"))
}

/*
func customHTTPHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	w.WriteHeader(http.StatusOK)
	WriteText(w, "hello")
	return true, nil
}*/

func TestServeOK(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	handler := newTestServerHandler()
	actBytes := newTestAction("Hi")
	req1, _ := http.NewRequest("GET", "/", bytes.NewReader(actBytes))
	resp1 := httptest.NewRecorder()
	handler.ServeHTTP(resp1, req1)
	res1 := resp1.Result()

	returnBytes := resp1.Body.Bytes()
	returnLen := len(returnBytes)
	resp := newTestActionResponse()
	okLen := len(resp)
	assert.Equal(200, res1.StatusCode)
	assert.Equal(okLen, returnLen)
	assert.Equal(resp[0], returnBytes[0])
	assert.Equal("application/octet-stream", res1.Header.Get("Content-Type"))
}

func TestServe404(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	handler := newTestServerHandler()
	req1, _ := http.NewRequest("GET", "/favicon.ico", nil)
	resp1 := httptest.NewRecorder()
	handler.ServeHTTP(resp1, req1)
	res1 := resp1.Result()
	assert.Equal(400, res1.StatusCode)
}

func newTestServerHandler() http.Handler {
	return CommandEntry(&mock.MapXXX{})
}

func newTestAction(text string) []byte {
	dispatch := &command.Dispatch{
		Map: &mock.MapXXX{},
	}
	act := &mock.CmdRespond{
		Text: text,
	}
	actBytes, _ := dispatch.EncodeCommand(act.XXX_MapID(), act)
	return actBytes
}

func newTestActionResponse() []byte {
	dispatch := command.Dispatch{}
	resp := &mock.CmdResponse{}
	bytes, _ := dispatch.EncodeCommand(resp.XXX_MapID(), resp)
	return bytes
}

func TestServerContextCanceled(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	dateline := time.Now().Add(-1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), dateline)
	defer cancel()
	err := ctx.Err()
	assert.True(errors.Is(err, context.DeadlineExceeded))
}

func newDeadlineAction() []byte {
	dispatch := &command.Dispatch{
		Map: &mock.MapXXX{},
	}
	act := &mock.CmdDeadline{}
	actBytes, _ := dispatch.EncodeCommand(act.XXX_MapID(), act)
	return actBytes
}

func TestServeWhenContextCanceled(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	handler := newTestServerHandler()
	actBytes := newDeadlineAction()
	req, _ := http.NewRequest("GET", "/", bytes.NewReader(actBytes))
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	res := resp.Result()
	assert.Equal(504, res.StatusCode)
}

func TestServerHandleRouteException(t *testing.T) {
	t.Parallel()
	//r, _ := http.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()
	handleRouteException(context.Background(), w, context.DeadlineExceeded)
	handleRouteException(context.Background(), w, errors.New(""))
}

func TestServer(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	server := &Server{}
	assert.Panics(server.Start)
}

func TestServerQuery(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	r, err := http.NewRequest("GET", "/?type=maintenance", nil)
	assert.Nil(err)

	// success
	value, ok := Query(r, "type")
	assert.True(ok)
	assert.Equal("maintenance", value)

	// failed
	value, ok = Query(r, "notExist")
	assert.False(ok)
	assert.Equal("", value)
}
