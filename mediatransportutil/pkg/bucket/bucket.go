package bucket

import (
	"encoding/binary"
	"fmt"
	"math"
)

const (
	MaxPktSize     = 1500
	pktSizeHeader  = 2
	seqNumOffset   = 2
	seqNumSize     = 2
	invalidPktSize = uint16(65535)
)

type Bucket struct {
	buf []byte
	src *[]byte

	init               bool
	resyncOnNextPacket bool
	step               int
	headSN             uint16
	maxSteps           int
}

func NewBucket(buf *[]byte) *Bucket {
	b := &Bucket{
		src:      buf,
		buf:      *buf,
		maxSteps: int(math.Floor(float64(len(*buf)) / float64(MaxPktSize))),
	}

	b.invalidate(0, b.maxSteps)
	return b
}

func (b *Bucket) ResyncOnNextPacket() {
	b.resyncOnNextPacket = true
}

func (b *Bucket) Src() *[]byte {
	return b.src
}

func (b *Bucket) HeadSequenceNumber() uint16 {
	return b.headSN
}

func (b *Bucket) AddPacket(pkt []byte) ([]byte, error) {
	sn := binary.BigEndian.Uint16(pkt[seqNumOffset : seqNumSize+seqNumSize])
	if !b.init {
		b.headSN = sn - 1
		b.init = true
	}

	if b.resyncOnNextPacket {
		b.resyncOnNextPacket = false

		b.headSN = sn - 1
		b.invalidate(0, b.maxSteps)
	}

	diff := sn - b.headSN
	if diff == 0 || diff > (1<<15) {
		// duplicate of last packet or out-of-order
		return b.set(sn, pkt)
	}

	return b.push(sn, pkt)
}

func (b *Bucket) GetPacket(buf []byte, sn uint16) (int, error) {
	p, err := b.get(sn)
	if err != nil {
		return 0, err
	}
	n := len(p)
	if cap(buf) < n {
		return 0, ErrBufferTooSmall
	}
	if len(buf) < n {
		buf = buf[:n]
	}
	copy(buf, p)
	return n, nil
}

func (b *Bucket) push(sn uint16, pkt []byte) ([]byte, error) {
	diff := int(sn-b.headSN) - 1
	b.headSN = sn

	// 如果序列号有间隙，则使插槽无效
	b.invalidate(b.step, diff)

	// 存储头SN数据包
	off := b.offset(b.step + diff)
	storedPkt := b.store(off, pkt)

	// 用于下一个数据包
	b.step = b.wrap(b.step + diff + 1)

	return storedPkt, nil
}

func (b *Bucket) get(sn uint16) ([]byte, error) {
	diff := int(int16(b.headSN - sn))
	if diff < 0 {
		// 在headSN之前要求一些东西
		return nil, fmt.Errorf("%w, headSN %d, sn %d", ErrPacketTooNew, b.headSN, sn)
	}
	if diff >= b.maxSteps {
		// too old
		return nil, fmt.Errorf("%w, headSN %d, sn %d", ErrPacketTooOld, b.headSN, sn)
	}

	off := b.offset(b.step - diff - 1)
	cacheSN := binary.BigEndian.Uint16(b.buf[off+pktSizeHeader+seqNumOffset : off+pktSizeHeader+seqNumOffset+seqNumSize])
	if cacheSN != sn {
		return nil, fmt.Errorf("%w, headSN %d, sn %d, cacheSN %d", ErrPacketMismatch, b.headSN, sn, cacheSN)
	}

	sz := binary.BigEndian.Uint16(b.buf[off : off+pktSizeHeader])
	if sz == invalidPktSize {
		return nil, fmt.Errorf("%w, headSN %d, sn %d, size %d", ErrPacketSizeInvalid, b.headSN, sn, sz)
	}

	off += pktSizeHeader
	return b.buf[off : off+int(sz)], nil
}

func (b *Bucket) set(sn uint16, pkt []byte) ([]byte, error) {
	diff := int(b.headSN - sn)
	if diff >= b.maxSteps {
		return nil, fmt.Errorf("%w, headSN %d, sn %d", ErrPacketTooOld, b.headSN, sn)
	}

	off := b.offset(b.step - diff - 1)

	// 重复时不覆盖
	if binary.BigEndian.Uint16(b.buf[off+pktSizeHeader+seqNumOffset:off+pktSizeHeader+seqNumOffset+seqNumSize]) == sn {
		return nil, ErrRTXPacket
	}

	return b.store(off, pkt), nil
}

func (b *Bucket) store(off int, pkt []byte) []byte {
	// 存储数据包大小
	binary.BigEndian.PutUint16(b.buf[off:], uint16(len(pkt)))

	// 存储包
	off += pktSizeHeader
	copy(b.buf[off:], pkt)

	return b.buf[off : off+len(pkt)]
}

func (b *Bucket) wrap(slot int) int {
	for slot < 0 {
		slot += b.maxSteps
	}

	for slot >= b.maxSteps {
		slot -= b.maxSteps
	}

	return slot
}

func (b *Bucket) offset(slot int) int {
	return b.wrap(slot) * MaxPktSize
}

func (b *Bucket) invalidate(startSlot int, numSlots int) {
	if numSlots > b.maxSteps {
		numSlots = b.maxSteps
	}

	for i := 0; i < numSlots; i++ {
		off := b.offset(startSlot + 1)
		binary.BigEndian.PutUint16(b.buf[off:], invalidPktSize)
	}
}
