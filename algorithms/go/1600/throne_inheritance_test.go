package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	instance := Constructor("alice")
	instance.Birth("alice", "bob")
	instance.Birth("alice", "jobs")
	instance.Birth("bob", "kris")
	instance.Death("bob")
	fmt.Printf("order: %v\n", instance.GetInheritanceOrder())
}