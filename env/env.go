package env

import (
	"context"
	"net/http"
	"os"

	"github.com/piyuo/libsrv/util"
)

var AppName = os.Getenv("NAME")

var Debug = os.Getenv("DEBUG") != ""

var Region = os.Getenv("REGION")

// KeyContext define key used in ctx
//
type KeyContext int

const (
	// KeyContextRequest is context key name for request
	//
	KeyContextRequest KeyContext = iota

	// KeyContextUserID is context key name for user id
	//
	KeyContextUserID

	// KeyContextAccountID is context key name for account id
	//
	KeyContextAccountID

	// KeyContextLocale used in i18n to mock locale
	//
	KeyContextLocale
)

// Mock define key test flag
//
type Mock int8

const (
	// MockIP provide mock ip 127.0.0.1 on GetIP()
	//
	MockIP Mock = iota
)

// GetRequest get current request from context
//
//	request := GetRequest(ctx)
//
func GetRequest(ctx context.Context) *http.Request {
	iRequest := ctx.Value(KeyContextRequest)
	if iRequest != nil {
		return iRequest.(*http.Request)
	}
	return nil
}

// SetRequest set request into ctx, this may used in log and data package
//
//	ctx = SetRequest(ctx,request)
//
func SetRequest(ctx context.Context, request *http.Request) context.Context {
	return context.WithValue(ctx, KeyContextRequest, request)
}

// GetIP return ip from current request, return empty if anything wrong
//
//	ip := GetIP(ctx)
//
func GetIP(ctx context.Context) string {
	if ctx.Value(MockIP) != nil {
		return "127.0.0.1"
	}

	value := ctx.Value(KeyContextRequest)
	if value == nil {
		return ""
	}
	req := value.(*http.Request)
	return util.GetIP(req)
}

// GetUserAgentID return short id from user agent. no version in here cause we used this for refresh token
//
//	ua := GetUserAgentID(ctx) // "iPhone,iOS,Safari"
//
func GetUserAgentID(ctx context.Context) string {
	value := ctx.Value(KeyContextRequest)
	if value == nil {
		return ""
	}
	req := value.(*http.Request)
	return util.GetUserAgentID(req)
}

// GetUserAgentString return short string with version info from user agent
//
//	ua := GetUserAgentString(ctx) // "iPhone,iOS 7.0,Safari 6.0"
//
func GetUserAgentString(ctx context.Context) string {
	value := ctx.Value(KeyContextRequest)
	if value == nil {
		return ""
	}
	req := value.(*http.Request)
	return util.GetUserAgentString(req)
}

// GetUserAgent return user agent from current request, return empty if anything wrong
//
//	ua := GetUserAgent(ctx) //"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25"
//
func GetUserAgent(ctx context.Context) string {
	value := ctx.Value(KeyContextRequest)
	if value == nil {
		return ""
	}
	req := value.(*http.Request)
	return util.GetUserAgent(req)
}

// SetUserID set UserID into ctx, this may used in log and data package
//
//	ctx = SetUserID(ctx,"user id")
//
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, KeyContextUserID, userID)
}

// GetUserID return current user id from context
//
//	userID := GetUserID(ctx)
//
func GetUserID(ctx context.Context) string {
	iUserID := ctx.Value(KeyContextUserID)
	if iUserID != nil {
		return iUserID.(string)
	}
	return ""
}

// SetAccountID set AccountID into ctx, this may used in log and data package
//
//	ctx = SetAccountID(ctx,"account id")
//
func SetAccountID(ctx context.Context, accountID string) context.Context {
	return context.WithValue(ctx, KeyContextAccountID, accountID)
}

// GetAccountID return current user id from context
//
//	accountID := GetAccountID(ctx)
//
func GetAccountID(ctx context.Context) string {
	iAccountID := ctx.Value(KeyContextAccountID)
	if iAccountID != nil {
		return iAccountID.(string)
	}
	return ""
}
