package errors

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

const (
	DDTagsErr = "dd_tags"
)

type HTTPErrResponse struct {
	Status int
	Body   interface{}
}

type errWrapper struct {
	error
	stack    *stack
	msg      string
	fields   map[string]interface{}
	httpResp HTTPErrResponse
	retry    bool
}

type WrapOption func(*errWrapper)

type ErrorWrapper interface {
	StackTrace() errors.StackTrace
	GetMessage() string
	GetFields() map[string]interface{}
	GetHTTPErrResp() HTTPErrResponse
}

//please send string only arguments, avoid using formatted string
//use `WithParam` if want to add unique value to message i.e user_id, topic, awb
func WithMessage(m string) WrapOption {
	return func(stack *errWrapper) {
		if stack.msg == "" {
			stack.msg = m
		} else {
			stack.msg = fmt.Sprintf("%s: %s", stack.msg, m)
		}
	}
}

func WithOrderID(orderID int64) WrapOption {
	return func(stack *errWrapper) {
		stack.fields["order_id"] = orderID
	}
}

func WithUserID(userID interface{}) WrapOption {
	return func(stack *errWrapper) {
		stack.fields["user_id"] = userID
	}
}

func WithDeliveryID(dID int64) WrapOption {
	return func(stack *errWrapper) {
		stack.fields["delivery_id"] = dID
	}
}

//use this on repo level only i.e (db, apicalls, redis)
func WithDDTags(dd DDTags) WrapOption {
	return func(stack *errWrapper) {
		if dd.funcName == "" {
			pc, _, _, ok := runtime.Caller(2)
			fn := runtime.FuncForPC(pc)
			if ok && fn != nil {
				dd.funcName = funcname(fn.Name())
			}
		}

		stack.fields[DDTagsErr] = dd.GetString()
	}
}

func WithParam(key string, val interface{}) WrapOption {
	return func(stack *errWrapper) {
		stack.fields[key] = val
	}
}

//WithRetry will flag this caller function to be retry-able on usecase `NotifyRetry`
func WithRetry() WrapOption {
	return func(stack *errWrapper) {
		stack.retry = true
	}
}

// WithHTTPErrResp - Set Http Error Response to response client
func WithHTTPErrResp(status int, body interface{}) WrapOption {
	return func(stack *errWrapper) {
		stack.httpResp.Status = status
		stack.httpResp.Body = body
	}
}

// Log an error or return the error, never both - Dave Channey
// In order to do that, we need a modified implementation of error. Put more context into error and print the context when we need to log. That way we don't need to log and return an error at the same time, just to put more context into the log.
// How to do that:
// 1. Wrap the error on the inner most function, i.e. on db repo(when error query), uicontext(when error parsing)
// 2. Log the error on the outer most function, i.e. on handler layer(when returning a response to outside world)
//    in order to log the error with all the context, use ../log/log.go[log.Errorw] for logging the error
//
// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied wrap option.
// If err is nil, Wrap returns nil.
func Wrap(err error, options ...WrapOption) error {
	if err == nil {
		return nil
	}

	errWrapper, ok := err.(*errWrapper)
	if !ok || errWrapper == nil {
		errWrapper = wrapErr(err, 1)
	}

	for _, opt := range options {
		opt(errWrapper)
	}

	return errWrapper
}

//Newr new wrapper, creating new error wrapper from message
func Newr(message string, options ...WrapOption) error {
	e := wrapErr(errors.New(message), 1)
	for _, opt := range options {
		opt(e)
	}
	return e
}

func IsEqual(err error, err2 error) bool {
	return err.Error() == err2.Error()
}

//IsRetry determine if this error is retry-able
func IsRetry(err error) bool {
	if err == nil {
		return false
	}

	errWrapper, ok := err.(*errWrapper)
	if !ok || errWrapper == nil {
		return false
	}

	return errWrapper.retry
}

func wrapErr(err error, caller int) *errWrapper {
	s := &errWrapper{}
	s.error = err
	s.stack = callers(caller)
	s.fields = map[string]interface{}{}
	return s
}

// CastToHTTPErrResp - cast to http err response for response to client
func CastToHTTPErrResp(error error) HTTPErrResponse {
	errWrapper, ok := error.(ErrorWrapper)
	if !ok || errWrapper == nil {
		return HTTPErrResponse{}
	}

	return errWrapper.GetHTTPErrResp()
}

func (e *errWrapper) StackTrace() errors.StackTrace {
	f := make([]errors.Frame, len(*e.stack))
	for i := 0; i < len(f); i++ {
		f[i] = errors.Frame((*e.stack)[i])
	}
	return f
}

func (e *errWrapper) GetMessage() string {
	return e.msg
}

func (e *errWrapper) GetFields() map[string]interface{} {
	return e.fields
}

func (e *errWrapper) GetHTTPErrResp() HTTPErrResponse {
	return e.httpResp
}

// stack represents a stack of program counters.
type stack []uintptr

func callers(pos int) *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[pos:n]
	return &st
}

func funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	n := strings.Split(name, ".")
	return n[len(n)-1]
}
