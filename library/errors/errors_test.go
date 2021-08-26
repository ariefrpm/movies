package errors

import (
	"errors"
	"testing"

	pkgErr "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	Wrap(nil)
	err := errors.New("error")

	error := Wrap(err, WithMessage("error add msg"))
	assertMessage(error, "error add msg", t)

	error = Wrap(err, WithParam("param", "value"))
	assertField(error, "param", "value", t)

	error = Wrap(err, WithDDTags(ErrDBDeliveries()))
	assertField(error, "dd_tags", "func:TestWrap|pic:mp-logistic|repo:db", t)

	error = Wrap(err, WithOrderID(111))
	assertField(error, "order_id", int64(111), t)

	error = Wrap(err, WithUserID(1))
	assertField(error, "user_id", 1, t)

	error = Wrap(err, WithDeliveryID(1))
	assertField(error, "delivery_id", int64(1), t)

	error = Newr("error", WithOrderID(1))
	assertField(error, "order_id", int64(1), t)

	error = Wrap(err, WithHTTPErrResp(404, "Not Found"))
	errWrapper, _ := error.(ErrorWrapper)
	assert.Equal(t, HTTPErrResponse{Status: 404, Body: "Not Found"}, errWrapper.GetHTTPErrResp())
}

func assertMessage(error error, message string, t *testing.T) {
	errWrapper, _ := error.(ErrorWrapper)
	assert.Equal(t, message, errWrapper.GetMessage())
}
func assertField(error error, field string, value interface{}, t *testing.T) {
	errWrapper, _ := error.(ErrorWrapper)
	assert.Equal(t, value, errWrapper.GetFields()[field])
}

func TestIsEqual(t *testing.T) {
	var eq bool
	var err1 error
	var err2 error

	// equal
	err1 = errors.New("errA")
	err2 = errors.New("errA")
	eq = IsEqual(err1, err2)
	assert.Equal(t, true, eq)

	// not equal
	err1 = errors.New("errA")
	err2 = errors.New("errB")
	eq = IsEqual(err1, err2)
	assert.Equal(t, false, eq)
}

func TestIsRetry(t *testing.T) {
	var err error

	// test err nil
	err = nil
	rt := IsRetry(err)
	assert.Equal(t, false, rt)

	// test errWrapper nil
	Wrap(nil)
	err = errors.New("error")
	rt = IsRetry(err)
	assert.Equal(t, false, rt)

	// test retry able
	err = Wrap(err, WithRetry())
	rt = IsRetry(err)
	assert.Equal(t, true, rt)
}

func TestStackTrace(t *testing.T) {
	err := Newr("error coy")
	err = Wrap(err, WithDeliveryID(1), WithMessage("error 2 coy"))
	err = Wrap(err, WithOrderID(1), WithMessage("error 3 coy"))

	errWrapper := err.(ErrorWrapper)
	stackTrace := errWrapper.StackTrace()
	assert.IsType(t, pkgErr.StackTrace{}, stackTrace)
}

func TestCastToHTTPErrResp(t *testing.T) {
	err := Newr("error", WithHTTPErrResp(503, "service unavailable"))
	casted := CastToHTTPErrResp(err)
	assert.Equal(t, HTTPErrResponse{Status: 503, Body: "service unavailable"}, casted)

	// test wrap empty
	err = Wrap(nil, WithHTTPErrResp(404, "not found"))
	casted = CastToHTTPErrResp(err)
	assert.Equal(t, HTTPErrResponse{}, casted)
}

func TestWrapDDTags(t *testing.T) {
	err := errors.New("error")

	errWrap := Wrap(err, WithDDTags(ErrDBDeliveries()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:mp-logistic|repo:db", t)

	errWrap = Wrap(err, WithDDTags(ErrDBOrder()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:order|repo:db", t)

	errWrap = Wrap(err, WithDDTags(ErrDBCore()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:core|repo:db", t)

	errWrap = Wrap(err, WithDDTags(ErrDBCore()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:core|repo:db", t)

	errWrap = Wrap(err, WithDDTags(ErrDBLogistic()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:logistic|repo:db", t)

	errWrap = Wrap(err, WithDDTags(ErrAPIKero()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:kero|repo:api", t)

	errWrap = Wrap(err, WithDDTags(ErrAPIKrab()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:mrkrab|repo:api", t)

	errWrap = Wrap(err, WithDDTags(ErrAPISonic()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:sonic|repo:api", t)

	errWrap = Wrap(err, WithDDTags(ErrAPIOrder()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:order|repo:api", t)

	errWrap = Wrap(err, WithDDTags(ErrAPIElastic()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:elastic|repo:api", t)

	errWrap = Wrap(err, WithDDTags(ErrAPITome()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:tome|repo:api", t)

	errWrap = Wrap(err, WithDDTags(ErrAPIPayment()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:payment|repo:api", t)

	errWrap = Wrap(err, WithDDTags(ErrAPITxV2()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:txv2|repo:api", t)

	errWrap = Wrap(err, WithDDTags(ErrAPINSQ()))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:nsq|repo:api", t)

	errWrap = Wrap(err, WithDDTags(ErrRedis("redis_err")))
	assertField(errWrap, "dd_tags", "func:TestWrapDDTags|pic:redis_err|repo:redis", t)

	ddTags := DDTags{
		repo: "testRepo",
		pic:  "testPic",
	}
	ddTags = ddTags.SetFuncName("testCustomFunc")
	errWrap = Wrap(err, WithDDTags(ddTags))
	assertField(errWrap, "dd_tags", "func:testCustomFunc|pic:testPic|repo:testRepo", t)
}
