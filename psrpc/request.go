package psrpc

import "time"

type RequestOption func(*RequestOpts)

type RequestOpts struct {
	Timeout       time.Duration
	SelectionOpts SelectionOpts
}

func WithRequestTimeout(timeout time.Duration) RequestOption {
	return func(o *RequestOpts) {
		o.Timeout = timeout
	}
}

type SelectionOpts struct {
	MinimumAffinity      float32       // 服务器被视为有效处理程序的最小关联性
	AcceptFirstAvailable bool          // go 快点
	AffinityTimeout      time.Duration // 服务器选择截止日期
	ShortCircuitTimeout  time.Duration // 收到第一个响应后施加的截止日期
}

func WithSelectionOpts(opts SelectionOpts) RequestOption {
	return func(o *RequestOpts) {
		o.SelectionOpts = opts
	}
}
