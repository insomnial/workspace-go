// Package name
package dast

import (
    "fmt"
)

// Creates a new node
type Node struct {
    Forw Node
    Back Node
    Value string
}
