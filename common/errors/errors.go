package errors

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

// HTTPErrResponse structure
type HTTPErrResponse struct {
	Status int
	Body   interface{}
}

// ErrorWrapper interface definition
type ErrorWrapper interface {
	StackTrace() errors.StackTrace
	GetMessage() string
	GetFields() map[string]interface{}
	GetHTTPErrResp() HTTPErrResponse
}

// WrapOption function definition
type WrapOption func(*errWrapper)

// WithMessage create new errors wrap option with given message
func WithMessage(msg string) WrapOption {
	return func(stack *errWrapper) {
		if stack.msg == "" {
			stack.msg = msg
		} else {
			stack.msg = fmt.Sprintf("%s: %s", stack.msg, msg)
		}
	}
}

// WithParam create new errors wrap option with given parameter
func WithParam(key string, val interface{}) WrapOption {
	return func(stack *errWrapper) {
		stack.fields[key] = val
	}
}

// WithHTTPErrResp create new errors wrap option with HTTP response status and body
func WithHTTPErrResp(status int, body interface{}) WrapOption {
	return func(stack *errWrapper) {
		stack.httpResp.Status = status
		stack.httpResp.Body = body
	}
}

// Wrap return an error annotated with stack trace and wrap options
func Wrap(err error, options ...WrapOption) error {
	if err == nil {
		return nil
	}

	wrapper, ok := err.(*errWrapper)
	if !ok || wrapper == nil {
		wrapper = wrapErr(err, 1)
	}

	for _, opt := range options {
		opt(wrapper)
	}

	return wrapper
}

// Newr create new wrapped error with stack trace and wrap options
func Newr(msg string, options ...WrapOption) error {
	err := wrapErr(errors.New(msg), 1)
	for _, opt := range options {
		opt(err)
	}

	return err
}

// IsEqual compare two error for equality
func IsEqual(err1, err2 error) bool {
	return err1.Error() == err2.Error()
}

// CastToHTTPErrResp convert given error into HTTPErrResponse
func CastToHTTPErrResp(err error) HTTPErrResponse {
	errWrapper, ok := err.(ErrorWrapper)
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

func wrapErr(err error, caller int) *errWrapper {
	s := &errWrapper{}
	s.error = err
	s.stack = callers(caller)
	s.fields = map[string]interface{}{}
	return s
}

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

type errWrapper struct {
	error
	stack    *stack
	msg      string
	fields   map[string]interface{}
	httpResp HTTPErrResponse
}
