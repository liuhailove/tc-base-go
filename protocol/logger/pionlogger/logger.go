package pionlogger

import (
	"github.com/pion/logging"

	"github.com/liuhailove/tc-base-go/protocol/logger"
)

var (
	pionIgnoredPrefixes = map[string][]string{
		"ice": {
			"pingAllCandidates called with no candidate pairs",
			"failed to send packet: io: read/write on closed pipe",
			"Ignoring remote candidate with tcpType active",
			"discard message from",
			"Failed to discover mDNS candidate",
			"Failed to read from candidate tcp",
			"remote mDNS candidate added, but mDNS is disabled",
		},
		"pc": {
			"Failed to accept RTCP stream is already closed",
			"Failed to accept RTP stream is already closed",
			"Incoming unhandled RTCP ssrc",
		},
		"tcp_mux": {
			"Error reading first packet from",
			"error closing connection",
		},
		"turn": {
			"error when handling datagram",
			"Failed to send ChannelData from allocation",
			"Failed to handle datagram",
		},
	}
)

// LoggerFactory implements webrtc.LoggerFactory interface
type LoggerFactory struct {
	logger logger.Logger
}

func NewLoggerFactory(logger logger.Logger) *LoggerFactory {
	return &LoggerFactory{
		logger: logger,
	}
}

func (f *LoggerFactory) NewLogger(scope string) logging.LeveledLogger {
	return &logAdapter{
		logger:          f.logger.WithComponent("pion." + scope),
		ignoredPrefixes: pionIgnoredPrefixes[scope],
	}
}
