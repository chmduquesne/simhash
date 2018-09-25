package charikar64

import (
	"github.com/chmduquesne/rollinghash/buzhash64"
)

const Size = 8
const ShingleSize = 8

type Charikar64 struct {
	counter [64]int32
	rollsum *buzhash64.Buzhash64

	n int64
}

func (d *Charikar64) Reset() {
	for i := 0; i < 64; i++ {
		d.counter[i] = 0
	}
	d.rollsum.Reset()
	s := d.rollsum.Sum64()
	for i := 0; i < 64; i++ {
		if s & (1 << uint(i)) > 0 {
			d.counter[i]++
		} else {
			d.counter[i]--
		}
	}
}

func New() (digest *Charikar64) {
	digest = &Charikar64{
		rollsum : buzhash64.New(),
		n : 0,
	}
	digest.Reset()
	return digest
}

func (d *Charikar64) Size() int { return Size }

func (d *Charikar64) BlockSize() int { return 1 }

func (d *Charikar64) Write(data []byte) (int, error) {
	for _, b := range(data) {
		d.n++
		if d.n < ShingleSize {
			s := make([]byte, 1)
			s[0] = b
			d.rollsum.Write(s)
		} else {
			d.rollsum.Roll(b)
			s := d.rollsum.Sum64()
			for i := 0; i < 64; i++ {
				if s & (1 << uint(i)) > 0 {
					d.counter[i]++
				} else {
					d.counter[i]--
				}
			}
		}
	}
	return len(data), nil
}

func (d *Charikar64) Sum64() uint64{
	if d.n <= ShingleSize {
		return d.rollsum.Sum64()
	}

	sum := uint64(0)
	for i := 0; i < 64; i++ {
		b := uint64(0)
		if d.counter[i] >= 0 {
			b = 1
		} else {
			b = 0
		}
		sum |= (b << uint(i))
	}
	return sum
}

func (d *Charikar64) Sum(b []byte) []byte {
	v := d.Sum64()
	return append(b, byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}
