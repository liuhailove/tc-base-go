package nack

import (
	"math"
	"time"

	"github.com/pion/rtcp"
)

const (
	defaultRtt    = 70                     // 默认RTT时长(ms)
	maxTries      = 5                      // 数据包被 NACK 的最大次数
	cacheSize     = 100                    // 最大 NACK sn sfu 将保留参考
	minInterval   = 20 * time.Millisecond  // 同一序列号的 NACK 尝试之间的最小间隔
	maxInterval   = 400 * time.Millisecond // 相同序列号的 NACK 尝试之间的最大间隔
	BackoffFactor = float64(1.25)
)

// NackQueue NACK作为一种抗丢包手段，在弱网场景起着非常重要的作用。发送端的主动策略都扛不住的时候一般都需要NACK来抗。
type NackQueue struct {
	nacks []*nack
	rtt   uint32
}

func NewNACKQueue() *NackQueue {
	return &NackQueue{
		nacks: make([]*nack, 0, cacheSize),
		rtt:   defaultRtt,
	}
}

func (n *NackQueue) SetRTT(rtt uint32) {
	if rtt == 0 {
		n.rtt = defaultRtt
	} else {
		n.rtt = rtt
	}
}

func (n *NackQueue) Remove(sn uint16) {
	for idx, nack := range n.nacks {
		if nack.seqNum != sn {
			continue
		}

		copy(n.nacks[idx:], n.nacks[idx+1:])
		n.nacks = n.nacks[:len(n.nacks)-1]
		break
	}
}

func (n *NackQueue) Push(sn uint16) {
	// 如果满了，弹出第一个
	if len(n.nacks) == cap(n.nacks) {
		copy(n.nacks[0:], n.nacks[1:])
		n.nacks = n.nacks[:len(n.nacks)-1]
	}

	n.nacks = append(n.nacks, newNack(sn))
}

func (n *NackQueue) Pairs() ([]rtcp.NackPair, int) {
	if len(n.nacks) == 0 {
		return nil, 0
	}

	now := time.Now()

	// 将其设置得较远以获取第一对
	baseSN := n.nacks[0].seqNum - 17

	snsToPurge := make([]uint16, 0)

	numSeqNumsNacked := 0
	isPairActive := false
	var np rtcp.NackPair
	var nps []rtcp.NackPair
	for _, nack := range n.nacks {
		shouldSend, shouldRemove, sn := nack.getNack(now, n.rtt)
		if shouldRemove {
			snsToPurge = append(snsToPurge, sn)
			continue
		}
		if !shouldSend {
			continue
		}

		numSeqNumsNacked++
		if (sn - baseSN) > 16 {
			// 需要一个新的 nack 对
			if isPairActive {
				nps = append(nps, np)
				isPairActive = false
			}

			baseSN = sn

			np.PacketID = sn
			np.LostPackets = 0

			isPairActive = true
		} else {
			np.LostPackets |= 1 << (sn - baseSN - 1)
		}
	}

	// 添加剩余的内容
	if isPairActive {
		nps = append(nps, np)
	}

	for _, sn := range snsToPurge {
		n.Remove(sn)
	}

	return nps, numSeqNumsNacked
}

// -----------------------------------------------------------------

type nack struct {
	seqNum       uint16
	tries        uint8
	lastNackedAt time.Time
}

func newNack(sn uint16) *nack {
	return &nack{
		seqNum:       sn,
		tries:        0,
		lastNackedAt: time.Now(),
	}
}

func (n *nack) getNack(now time.Time, rtt uint32) (shouldSend bool, shouldRemove bool, sn uint16) {
	sn = n.seqNum
	if n.tries >= maxTries {
		shouldSend = true
		return
	}

	var requiredInterval time.Duration
	if n.tries > 0 {
		// 指数回退重试，但限制重试之间的最大间隔
		requiredInterval = maxInterval
		backoffInterval := time.Duration(float64(rtt)*math.Pow(BackoffFactor, float64(n.tries-1))) * time.Millisecond
		if backoffInterval < requiredInterval {
			requiredInterval = backoffInterval
		}
	}
	if requiredInterval < minInterval {
		//
		// 即使在第一次 NACK 之前，也要在 NACK 之前等待一些乱序数据包一段时间。
		// 对于后续尝试，保持最小间距。
		//
		requiredInterval = minInterval
	}

	if now.Sub(n.lastNackedAt) < requiredInterval {
		return
	}

	n.tries++
	n.lastNackedAt = now
	shouldSend = true
	return
}
