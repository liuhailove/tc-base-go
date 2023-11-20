package utils

import (
	"crypto/sha1"
	"fmt"
	"github.com/liuhailove/tc-base-go/protocol/tc"
	"os"

	"github.com/jxskiss/base62"
	"github.com/lithammer/shortuuid/v4"
)

const guidSize = 12

const (
	RoomPrefix         = "RM_"
	NodePrefix         = "ND_"
	ParticipantPrefix  = "PA_"
	TrackPrefix        = "RT_"
	APIKeyPrefix       = "API"
	EgressPrefix       = "EG_"
	IngressPrefix      = "IN_"
	RPCPrefix          = "RPC_"
	WHIPResourcePrefix = "WH_"
)

func NewGuid(prefix string) string {
	return prefix + shortuuid.New()[:guidSize]
}

// HashedID 从唯一字符串创建哈希 ID
func HashedID(id string) string {
	h := sha1.New()
	h.Write([]byte(id))
	val := h.Sum(nil)

	return base62.EncodeToString(val)
}

func LocalNodeID() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", NodePrefix, HashedID(hostname)[:8]), nil
}

var b62Index = newB62Index()
var b62Chars = []byte(shortuuid.DefaultAlphabet)

func newB62Index() [256]byte {
	var index [256]byte
	for i := 0; i < len(shortuuid.DefaultAlphabet); i++ {
		index[shortuuid.DefaultAlphabet[i]] = byte(i)
	}
	return index
}

func guidPrefix[T tc.Guid]() string {
	var id T
	switch any(id).(type) {
	case tc.TrackID:
		return TrackPrefix
	case tc.ParticipantID:
		return ParticipantPrefix
	case tc.RoomID:
		return RoomPrefix
	default:
		panic("unreachable")
	}
}

func MarshalGuild[T tc.Guid](id T) tc.GuidBlock {
	var b tc.GuidBlock
	idb := []byte(id)[len(guidPrefix[T]()):]
	for i := 0; i < 3; i++ {
		j := i * 3
		k := i * 4
		b[j] = b62Index[idb[k]]<<2 | b62Index[idb[k+1]]>>4
		b[j+1] = b62Index[idb[k+1]]<<4 | b62Index[idb[k+2]]>>2
		b[j+2] = b62Index[idb[k+2]]<<6 | b62Index[idb[k+3]]
	}
	return b
}

func UnmarshalGuid[T tc.Guid](b tc.GuidBlock) T {
	prefix := guidPrefix[T]()
	id := make([]byte, len(prefix)+guidSize)
	copy(id, []byte(prefix))
	idb := id[len(prefix):]
	for i := 0; i < 3; i++ {
		j := i * 3
		k := i * 4
		idb[k] = b62Chars[b[j]>>2]
		idb[k+1] = b62Chars[(b[j]&3)<<4|b[j+1]>>4]
		idb[k+2] = b62Chars[(b[j+1]&15)<<2|b[j+2]>>6]
		idb[k+3] = b62Chars[b[j+2]&63]
	}
	return T(id)
}
