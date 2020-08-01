package log

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGCPLogger(t *testing.T) {
	Convey("should write log", t, func() {
		appName = "log-gcp_test"
		ctx := context.Background()
		ctx = context.WithValue(ctx, keyToken, map[string]string{"id": "user1"})
		logger, err := NewGCPLogger(ctx)
		So(err, ShouldBeNil)
		So(logger, ShouldNotBeNil)
		defer logger.Close()

		logger.Write(ctx, DEBUG, here, "TestGCPLogger")

		shouldPrint = false
		//empty message will not log
		logger.Write(ctx, DEBUG, here, "no print console")

		//empty message will not log
		logger.Write(ctx, DEBUG, here, "")

	})
}