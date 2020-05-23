package mock

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSlowAction(t *testing.T) {
	Convey("should execute the action and get response", t, func() {
		action := &SlowAction{}
		// action.Name = "hello"

		response, err := action.Main(context.Background())
		So(err, ShouldBeNil)
		So(response, ShouldNotBeNil)
		//sr := response.(*StringResponse)
		//So(sr.Text, ShouldEqual, "hi")
	})

}
