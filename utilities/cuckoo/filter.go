package cuckoo

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	_ "errors"
	"strings"

	"math"
	"math/rand"
)

type bucket []byte
type fingerprint []byte

var hasher = sha256.New()

const retries = 500

type Cuckoo struct {
	buckets []bucket
	m       uint // buckets
	b       uint // entries per bucket
	f       uint // fingerprint length
	n       uint // filter capacity (rename cap?)
}

func NewCuckoo(n uint, fp float64) *Cuckoo {
	b := uint(4)
	f := fingerprintLength(b, fp)
	m := nextPower(n / f * 8)
	buckets := make([]bucket, m)
	//for i := uint(0); i < m; i++ {
	//	buckets[i] = make(bucket, b)
	//}
	return &Cuckoo{
		buckets: buckets,
		m:       m,
		b:       b,
		f:       f,
		n:       n,
	}
}

func (c *Cuckoo) delete(needle *id) {
	i1, i2, f := c.hashes(needle.IDString)
	// try to remove from f1
	b1 := c.buckets[i1%c.m]

	if _, ok := b1.contains(f); ok {
		b1 = []byte(strings.ReplaceAll(string(b1), string(f), ""))
		return
	}

	b2 := c.buckets[i2%c.m]
	if _, ok := b2.contains(f); ok {
		b1 = []byte(strings.ReplaceAll(string(b1), string(f), ""))
		return
	}
}

// lookup needle in the cuckoo filter
func (c *Cuckoo) lookup(needle *id) bool {
	i1, i2, f := c.hashes(needle.IDString)
	_, b1 := c.buckets[i1%c.m].contains(f)
	_, b2 := c.buckets[i2%c.m].contains(f)
	return b1 || b2
}

func (b bucket) contains(f fingerprint) (int, bool) {
	for i, x := range strings.Split(string(b), "|") {
		if bytes.Equal([]byte(x), f) {
			return i, true
		}
	}
	return -1, false
}

func (c *Cuckoo) insert(input *id) {
	i1, i2, f := c.hashes(input.IDString)
	// first try bucket one
	b1 := c.buckets[i1%c.m]

	if len(b1) < int(c.b*c.f) {
		b1 = append(b1, []byte(f)...)
		b1 = append(b1, '|')
		return
	}

	b2 := c.buckets[i2%c.m]

	if len(b2) < int(c.b*c.f) {
		b2 = append(b2, []byte(f)...)
		b2 = append(b2, '|')
		return
	}

	// else we need to start relocating items
	i := i1
	for r := 0; r < retries; r++ {
		index := i % c.m
		entryIndex := rand.Intn(int(c.b))
		// swap
		b := c.buckets[index]

		f, b = fingerprint(strings.Split(string(b), "|")[entryIndex]),
			append(b, []byte(f)...)
		f1 := append(f, '|')
		b = []byte(strings.ReplaceAll(string(b), string(f1), ""))

		i = i ^ uint(binary.BigEndian.Uint32(hash(f)))
		b = c.buckets[i%c.m]

		if len(b) < int(c.b*c.f) {
			b = append(b1, []byte(f1)...)
			//b = append(b1, '|')
			return
		}
	}
	panic("cuckoo filter full")
}

func (c *Cuckoo) hashes(data string) (uint, uint, fingerprint) {
	h := hash([]byte(data))
	f := h[0:c.f]
	i1 := uint(binary.BigEndian.Uint32(h))
	i2 := i1 ^ uint(binary.BigEndian.Uint32(hash(f)))
	return i1, i2, f
}

func hash(data []byte) []byte {
	hasher.Write([]byte(data))
	hash := hasher.Sum(nil)
	hasher.Reset()
	return hash
}

func fingerprintLength(b uint, e float64) uint {
	f := uint(math.Ceil(math.Log(2 * float64(b) / e)))
	f /= 8
	if f < 1 {
		return 1
	}
	return f
}

func nextPower(i uint) uint {
	i--
	i |= i >> 1
	i |= i >> 2
	i |= i >> 4
	i |= i >> 8
	i |= i >> 16
	i |= i >> 32
	i++
	return i
}
