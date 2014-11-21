// Package name
package dast

import (
    "fmt"
    "strconv"
)

type LL struct {
    size int
    head Node
    Iter Node
}

func getSize() (int) {
    return size
}

func build() {
    head := Node{ nil, nil, "one" }
    currentNode := *head
    for i := 0; i < 10; i++ {
        head.Forw := Node{ nil, currentNode, Itoa( i ) }
        currentNode = head.Forw
    }
    
     
}