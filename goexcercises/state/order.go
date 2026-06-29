package state

import (
	"fmt"
	"sync"
)

type PaymentStatus string

const (
	Paid          PaymentStatus = "paid"
	NotPaid       PaymentStatus = "notPaid"
	PaymentFailed PaymentStatus = "paymentFailed"
)

type OrderDetails struct {
	OrderId       string
	PaymentStatus PaymentStatus
}

type Context struct {
	mutex        sync.RWMutex
	state        State
	OrderDetails OrderDetails
}

func (c *Context) SetState(nextState State) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.state = nextState
}

func (c *Context) RequestForward() error {
	c.mutex.RLock()
	currentState := c.state
	c.mutex.RUnlock()
	return currentState.TransitionStateForward(c)
}

func (c *Context) RequestBackward() error {
	c.mutex.RLock()
	currentState := c.state
	c.mutex.RUnlock()
	return currentState.TransitionStateBackward(c)
}

type State interface {
	TransitionStateForward(ctx *Context) error
	TransitionStateBackward(ctx *Context) error
}

type CompletedState struct{}

func (c *CompletedState) TransitionStateForward(ctx *Context) error {
	fmt.Println("this is the final completed state...")
	return nil
}

func (c *CompletedState) TransitionStateBackward(ctx *Context) error {
	fmt.Println("rollback state called on completed state...going to progressing state")
	previousState := ProgressingState{}
	ctx.SetState(&previousState)
	return nil
}

type ProgressingState struct{}

func (p *ProgressingState) TransitionStateForward(ctx *Context) error {
	fmt.Println("forward on progressing state called")
	comletedState := CompletedState{}
	ctx.SetState(&comletedState)
	return nil
}

func (p *ProgressingState) TransitionStateBackward(ctx *Context) error {
	fmt.Println("rollback state called on progressing state...going to initial state")
	initializedState := InitializedState{}
	ctx.SetState(&initializedState)
	return nil
}

type InitializedState struct{}

func (i *InitializedState) TransitionStateForward(ctx *Context) error {
	progressingState := ProgressingState{}
	ctx.SetState(&progressingState)
	ctx.OrderDetails.PaymentStatus = "notPaid"
	return nil
}

func (i *InitializedState) TransitionStateBackward(ctx *Context) error {
	fmt.Println("rollback state called on initial state...nothing to do")
	return nil
}
