package bloom

import (
	"github.com/cespare/xxhash"
	"github.com/steakknife/bloomfilter"
)

const (
	maxElements = 100000000
	probCollide = 0.00001
)

var BF *bloomfilter.Filter

func init() {
	bf, err := bloomfilter.NewOptimal(maxElements, probCollide)
	if err != nil {
		panic(err)
	}
	BF = bf
}

func Add(v string) {
	hash := xxhash.New()
	hash.Write([]byte(v))
	BF.Add(hash)
}

func Contain(v string) bool {
	hash := xxhash.New()
	hash.Write([]byte(v))
	return BF.Contains(hash)
}
