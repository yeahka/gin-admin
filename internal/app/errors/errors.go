package errors

import (
	"github.com/pkg/errors"
)

// 定义别名
var (
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)

// 定义错误
var (
	ErrBadRequest = New400Response("请求发生错误")

	ErrNoPerm          = NewResponse(401, "无访问权限", 401)
	ErrNotFound        = NewResponse(404, "资源不存在", 404)
	ErrMethodNotAllow  = NewResponse(405, "方法不被允许", 405)
	ErrTooManyRequests = NewResponse(429, "请求过于频繁", 429)
	ErrInternalServer  = NewResponse(500, "服务器发生错误", 500)
)
