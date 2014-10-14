package morris

import (
	"fmt"
	"testing"
)

func TestIncrement(t *testing.T) {
	c := Counter(0)
	i := 1
	e := uint(27)
	for ; i < 1<<e; i++ {
		c.increment()
	}
	fmt.Printf("Expected := %d\t\tActual := %d\n", e, c.get())
}
