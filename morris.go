// A probabilistic Morris counter that uses only 1 byte and counts
// up to 2^(2^8) events.
package morris

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

type Counter byte

// Increments the counter by one.
func (c *Counter) increment() {
	r := uint64(random.Int63())
	mask := (uint64(1<<(*c%64)) - 1)

	incr := (mask&r == mask)

	repeat := int(*c / 64)
	for i := 0; incr && i < repeat; i++ {
		r = uint64(random.Int63())
		incr = incr && (r == ^uint64(0))
	}

	if incr {
		*c++
	}
}

// Returns log_2(counter).
func (c *Counter) get() byte {
	return byte(*c)
}
