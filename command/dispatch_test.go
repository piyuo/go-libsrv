package command

import (
	"context"
	"os"
	"strconv"
	"testing"

	mock "github.com/piyuo/libsrv/command/mock"
	shared "github.com/piyuo/libsrv/command/shared"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEncodeDecodeCommand(t *testing.T) {
	act := &mock.RespondAction{
		Text: "Hi",
	}
	dispatch := &Dispatch{
		Map: &mock.MapXXX{},
	}
	Convey("test decode command is right", t, func() {
		actBytes, err := dispatch.encodeCommand(act.XXX_MapID(), act)
		actID, iAct2, err2 := dispatch.decodeCommand(actBytes)
		So(err2, ShouldBeNil)
		act2 := iAct2.(*mock.RespondAction)
		So(err, ShouldBeNil)
		So(actID, ShouldEqual, act.XXX_MapID())
		So(act2.Text, ShouldEqual, act.Text)
	})
}

func TestBetterResponseName(t *testing.T) {
	Convey("should get better response name", t, func() {
		errOK := OK().(*shared.PbOK)
		result := betterResponseName(errOK.XXX_MapID(), errOK)
		So(result, ShouldEqual, "OK")

		err := Error("failed").(*shared.PbError)
		result = betterResponseName(err.XXX_MapID(), err)
		So(result, ShouldEqual, "failed")

		errText := &shared.PbString{}
		result = betterResponseName(errText.XXX_MapID(), errText)
		So(result, ShouldEqual, "PbString")
	})
}

func TestActionNoRespose(t *testing.T) {
	act := &mock.NoRespondAction{
		Text: "Hi",
	}
	dispatch := &Dispatch{
		Map: &mock.MapXXX{},
	}
	//no response action,will cause &shared.Err{}
	Convey("test dispatch route", t, func() {
		actBytes, err := dispatch.encodeCommand(act.XXX_MapID(), act)
		So(err, ShouldBeNil)
		resultBytes, err2 := dispatch.Route(context.Background(), actBytes)
		So(err2, ShouldNotBeNil)
		So(resultBytes, ShouldBeNil)
	})
}

func TestRoute(t *testing.T) {
	act := &mock.RespondAction{
		Text: "Hi",
	}
	dispatch := &Dispatch{
		Map: &mock.MapXXX{},
	}
	actBytes, err := dispatch.encodeCommand(act.XXX_MapID(), act)
	resultBytes, err2 := dispatch.Route(context.Background(), actBytes)
	_, resp, err3 := dispatch.decodeCommand(resultBytes)
	actualResponse := resp.(*shared.PbOK)
	Convey("test dispatch route", t, func() {
		So(err, ShouldBeNil)
		So(err2, ShouldBeNil)
		So(err3, ShouldBeNil)
		So(actualResponse, ShouldNotBeNil)
	})
}

func TestHandle(t *testing.T) {

	//create sample data
	act := &mock.RespondAction{
		Text: "Hi",
	}
	//create dispatch and register
	dispatch := &Dispatch{
		Map: &mock.MapXXX{},
	}
	Convey("should run action", t, func() {
		_, respInterface, err := dispatch.runAction(context.Background(), act)
		response := respInterface.(*shared.PbOK)
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)
	})
}

func TestTimeExecuteAction(t *testing.T) {
	Convey("should warn slow action", t, func() {
		os.Setenv("PIYUO_SLOW", "1")
		act := &mock.SlowAction{}
		dispatch := &Dispatch{
			Map: &mock.MapXXX{},
		}
		_, respInterface, err := dispatch.timeExecuteAction(context.Background(), act)
		So(err, ShouldBeNil)
		So(respInterface, ShouldNotBeNil)
	})
}

var benchmarkResult string

func BenchmarkStringMapSpeed(b *testing.B) {
	var list map[string]string
	list = make(map[string]string, 100)
	for x := 0; x < 100; x++ {
		list[strconv.Itoa(x)] = strconv.Itoa(x)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for x := 0; x < 100; x++ {
			benchmarkResult = list[strconv.Itoa(x)]
		}
	}
}

func BenchmarkIntMapSpeed(b *testing.B) {
	var list map[int]string
	list = make(map[int]string, 100)
	for x := 0; x < 100; x++ {
		list[x] = strconv.Itoa(x)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for x := 0; x < 100; x++ {
			benchmarkResult = list[x]
		}
	}
}

var tmp []byte

func BenchmarkAppend(b *testing.B) {
	bytes1 := []byte("my first slice")
	bytes2 := []byte("second slice")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for x := 0; x < 100; x++ {
			tmp = append(bytes1[:], bytes2[:]...)
		}
	}
}

func BenchmarkCopyPreAllocate(b *testing.B) {
	bytes1 := []byte("my first slice")
	bytes2 := []byte("second slice")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for x := 0; x < 100; x++ {
			totalLen := len(bytes1) + len(bytes2)
			tmp := make([]byte, totalLen)
			var i int
			i += copy(tmp, bytes1)
			i += copy(tmp[i:], bytes2)
		}
	}
}

func BenchmarkDispatch(b *testing.B) {
	act := &mock.RespondAction{
		Text: "Hi",
	}
	dispatch := &Dispatch{
		Map: &mock.MapXXX{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actBytes, _ := dispatch.encodeCommand(act.XXX_MapID(), act)
		resultBytes, _ := dispatch.Route(context.Background(), actBytes)
		_, _, _ = dispatch.decodeCommand(resultBytes)
	}
}
