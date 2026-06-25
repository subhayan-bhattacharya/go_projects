package main

import "memento"

func main() {
	t := memento.NewTextEditor("hello world")
	snapshot := t.Save()
	h := memento.History{}
	h.Push(snapshot)
}
