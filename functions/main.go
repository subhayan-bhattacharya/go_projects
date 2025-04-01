// types and methods
// Wonderful example of how we can call by value and the internal attributes do not get updated
package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total           int
	lastUpdatedTime time.Time
}

func (c *Counter) Increment() {
	c.total += 1
	c.lastUpdatedTime = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("%d , %v", c.total, c.lastUpdatedTime)
}

type Stringer interface {
	String() string
}

type Incrementor interface {
	Increment()
}

func main() {
	var myStringer Stringer
	var myIncrementor Incrementor
	pointerCounter := &Counter{}
	valueCounter := Counter{}
	myStringer = pointerCounter
	myIncrementor = pointerCounter
	fmt.Println(myStringer.String())
	myIncrementor.Increment()
	fmt.Println(myIncrementor)
	fmt.Println(valueCounter)

}
