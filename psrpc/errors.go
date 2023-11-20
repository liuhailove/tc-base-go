package psrpc

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/twitchtv/twirp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrRequestCanceled = NewErrorf(Canceled, "request canceled")
	ErrRequestTimeOut  = NewErrorf(DeadlineExceeded, "request timed out")
	ErrNoResponse      = NewErrorf(Unavailable, "no response from servers")
	ErrStreamEOF       = NewError(Unavailable, io.EOF)
	ErrClientClosed    = NewErrorf(Canceled, "client is closed")
	ErrServerClosed    = NewErrorf(Canceled, "server is closed")
	ErrStreamClosed    = NewErrorf(Canceled, "stream closed")
	ErrSlowConsumer    = NewErrorf(Unavailable, "stream message discarded by slow consumer")
)

type Error interface {
	error
	Code() ErrorCode

	// convenience methods

	ToHttp() int
	GRPCStatus() *status.Status
}

type ErrorCode string

func NewError(code ErrorCode, err error) Error {
	return &psrpcError{
		error: err,
		code:  code,
	}
}

func NewErrorf(code ErrorCode, msg string, args ...interface{}) Error {
	return &psrpcError{
		error: fmt.Errorf(msg, args...),
		code:  code,
	}
}

func NewErrorFromResponse(code, err string) Error {
	if code == "" {
		code = string(Unknown)
	}

	return &psrpcError{
		error: errors.New(err),
		code:  ErrorCode(code),
	}
}

const (
	OK ErrorCode = ""

	// Canceled 请求被client取消
	Canceled ErrorCode = "canceled"
	// MalformedRequest 不能unmarshal请求
	MalformedRequest ErrorCode = "malformed_request"
	// MalformedResponse 不能unmarshall结果
	MalformedResponse ErrorCode = "malformed_result"
	// DeadlineExceeded 请求超时
	DeadlineExceeded ErrorCode = "deadline_exceeded"
	// Unavailable 服务由于负载或者亲和性限制而不可用
	Unavailable ErrorCode = "unavailable"
	// Unknown （server返回non-psrpc错误）
	Unknown ErrorCode = "unknown"

	// InvalidArgument 在请求中有无效参数
	InvalidArgument ErrorCode = "invalid_argument"
	// NotFound 实体没有找到
	NotFound ErrorCode = "not_found"
	// AlreadyExists 尝试重复创建
	AlreadyExists ErrorCode = "already_exists"
	// PermissionDenied 调用方没有获取到权限
	PermissionDenied ErrorCode = "permission_denied"
	// ResourceExhausted 资源被耗尽，如内存或者额度
	ResourceExhausted ErrorCode = "resource_exhausted"
	// FailedPrecondition 执行请求时不一致的状态
	FailedPrecondition ErrorCode = "failed_precondition"
	// Aborted 请求aborted
	Aborted ErrorCode = "aborted"
	// OutOfRange 操作在范围之外
	OutOfRange ErrorCode = "out_of_range"
	// Unimplemented 服务端没有实现的操作
	Unimplemented ErrorCode = "unimplemented"
	// Internal 由于内部操作导致失败
	Internal ErrorCode = "internal"
	// DataLoss 不可恢复的数据丢失或损坏
	DataLoss ErrorCode = "data_loss"
	// Unauthenticated 和 PermissionDenied类似，调用方使用时没有被授权
	Unauthenticated ErrorCode = "unauthenticated"
)

type psrpcError struct {
	error
	code ErrorCode
}

func (e psrpcError) Code() ErrorCode {
	return e.code
}

func (e psrpcError) Unwrap() error {
	return e.error
}

func (e psrpcError) ToHttp() int {
	switch e.code {
	case OK:
		return http.StatusOK
	case Canceled, DeadlineExceeded:
		return http.StatusRequestTimeout
	case Unknown, MalformedResponse, Internal, DataLoss:
		return http.StatusInternalServerError
	case InvalidArgument, MalformedRequest:
		return http.StatusBadRequest
	case NotFound:
		return http.StatusNotFound
	case AlreadyExists, Aborted:
		return http.StatusConflict
	case PermissionDenied:
		return http.StatusForbidden
	case ResourceExhausted:
		return http.StatusTooManyRequests
	case FailedPrecondition:
		return http.StatusPreconditionFailed
	case OutOfRange:
		return http.StatusRequestedRangeNotSatisfiable
	case Unimplemented:
		return http.StatusNotImplemented
	case Unavailable:
		return http.StatusServiceUnavailable
	case Unauthenticated:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func (e psrpcError) GRPCStatus() *status.Status {
	var c codes.Code
	switch e.code {
	case OK:
		c = codes.OK
	case Canceled:
		c = codes.Canceled
	case Unknown:
		c = codes.Unknown
	case InvalidArgument, MalformedRequest:
		c = codes.InvalidArgument
	case DeadlineExceeded:
		c = codes.DeadlineExceeded
	case NotFound:
		c = codes.NotFound
	case AlreadyExists:
		c = codes.AlreadyExists
	case PermissionDenied:
		c = codes.PermissionDenied
	case ResourceExhausted:
		c = codes.ResourceExhausted
	case FailedPrecondition:
		c = codes.FailedPrecondition
	case Aborted:
		c = codes.Aborted
	case OutOfRange:
		c = codes.OutOfRange
	case Unimplemented:
		c = codes.Unimplemented
	case MalformedResponse, Internal:
		c = codes.Internal
	case Unavailable:
		c = codes.Unavailable
	case DataLoss:
		c = codes.DataLoss
	case Unauthenticated:
		c = codes.Unauthenticated
	default:
		c = codes.Unknown
	}

	return status.New(c, e.Error())
}

func (e psrpcError) toTwirp() twirp.Error {
	var c twirp.ErrorCode
	switch e.code {
	case OK:
		c = twirp.NoError
	case Canceled:
		c = twirp.Canceled
	case Unknown:
		c = twirp.Unknown
	case InvalidArgument:
		c = twirp.InvalidArgument
	case MalformedRequest, MalformedResponse:
		c = twirp.Malformed
	case DeadlineExceeded:
		c = twirp.DeadlineExceeded
	case NotFound:
		c = twirp.NotFound
	case AlreadyExists:
		c = twirp.AlreadyExists
	case PermissionDenied:
		c = twirp.PermissionDenied
	case ResourceExhausted:
		c = twirp.ResourceExhausted
	case FailedPrecondition:
		c = twirp.FailedPrecondition
	case Aborted:
		c = twirp.FailedPrecondition
	case OutOfRange:
		c = twirp.OutOfRange
	case Unimplemented:
		c = twirp.Unimplemented
	case Internal:
		c = twirp.Internal
	case Unavailable:
		c = twirp.Unavailable
	case DataLoss:
		c = twirp.DataLoss
	case Unauthenticated:
		c = twirp.Unauthenticated
	default:
		c = twirp.Unknown
	}
	return twirp.NewErrorf(c, e.Error())
}

func (e psrpcError) As(target any) bool {
	switch te := target.(type) {
	case *twirp.Error:
		*te = e.toTwirp()
		return true
	}

	return false
}
