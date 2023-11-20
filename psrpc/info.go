package psrpc

// RPCInfo rpc信息
type RPCInfo struct {
	Service string
	Method  string
	Topic   []string
	Multi   bool
}
