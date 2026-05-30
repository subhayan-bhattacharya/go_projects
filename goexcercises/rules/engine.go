package rules

import (
	"fmt"
	"strings"
)

type Rule[T any] interface {
	Evaluate(input T) bool
	Description() string
}

type Predicate[T any] func(T) bool

type ConditionRule[T any] struct {
	description string
	predicate   Predicate[T]
}

func Condition[T any](description string, predicate Predicate[T]) Rule[T] {
	return ConditionRule[T]{
		description: description,
		predicate:   predicate,
	}
}

func (c ConditionRule[T]) Evaluate(input T) bool {
	return c.predicate(input)
}

func (c ConditionRule[T]) Description() string {
	return c.description
}

type AndRule[T any] struct {
	children []Rule[T]
}

func And[T any](rules ...Rule[T]) Rule[T] {
	return AndRule[T]{
		children: rules,
	}
}

func (a AndRule[T]) Description() string {
	descriptions := make([]string, 0, len(a.children))
	for _, rule := range a.children {
		descriptions = append(descriptions, rule.Description())
	}
	joinedDescriptions := strings.Join(descriptions, "&&")
	return fmt.Sprintf("The rules which are evaluated using an AND condition are : \n %s", joinedDescriptions)
}

func (a AndRule[T]) Evaluate(input T) bool {
	for _, rule := range a.children {
		if !rule.Evaluate(input) {
			return false
		}
	}
	return true
}

type OrRule[T any] struct {
	children []Rule[T]
}

func Or[T any](rules ...Rule[T]) Rule[T] {
	return OrRule[T]{
		children: rules,
	}
}

func (o OrRule[T]) Evaluate(input T) bool {
	for _, rule := range o.children {
		if rule.Evaluate(input) {
			return true
		}
	}
	return false
}

func (o OrRule[T]) Description() string {
	descriptions := make([]string, 0, len(o.children))
	for _, rule := range o.children {
		descriptions = append(descriptions, rule.Description())
	}
	joinedDescriptions := strings.Join(descriptions, "||")
	return fmt.Sprintf("The rules which are evaluated using an AND condition are : \n %s", joinedDescriptions)
}
