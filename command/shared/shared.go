package shared

import (
	"context"
	"errors"

	app "github.com/piyuo/go-libsrv/app"
)

//ErrAccessTokenRequired mean service need access  token
var ErrAccessTokenRequired = errors.New("access token required")

//ErrAccessTokenExpired mean access token is expired, client need use refresh token to get new access token
var ErrAccessTokenExpired = errors.New("access token expired")

//ErrPaymentTokenRequired mean service need access toke that generate from user enter password in 5 min
var ErrPaymentTokenRequired = errors.New("payment token required")

// Token return token or ErrorResponse
//
// 	token, errResp := shared.NeedToken(ctx)
// 	if errResp != nil {
// 		return errResp, nil
// 	}
func Token(ctx context.Context) (app.Token, error) {
	token, err := app.TokenFromContext(ctx)
	if err != nil {
		return nil, ErrAccessTokenRequired
	}
	if token.Expired() {
		return nil, ErrAccessTokenExpired
	}
	return token, nil
}

//OK return code=0 no error response
//
//	return shared.OK(),nil
func OK() interface{} {
	return &Err{
		Code: 0,
	}
}

//Error return  error response with code
//
//	return shared.Error(shared.ErrorUnknown),nil
func Error(code int32, msg string) interface{} {
	return errorInt32(code, msg)
}

func errorInt32(code int32, tag string) interface{} {
	return &Err{
		Code: code,
		Msg:  tag,
	}
}

//String return string response
//
//	return shared.Text("hi"),nil
func String(text string) interface{} {
	return &Text{
		Value: text,
	}
}

//Number return number response
//
//	return shared.Number(101),nil
func Number(num int64) interface{} {
	return &Num{
		Value: num,
	}
}
