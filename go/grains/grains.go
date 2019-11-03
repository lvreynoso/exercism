// grains.go

package grains

import (
    "errors"
)

// Square takes in a square and returns the number of grains on that square
func Square(square int) (uint64, error) {
    if square < 1 || square > 64 {
        return 0, errors.New("there ain't no such square")
    }
    var total uint64
    total = 0
    for i := 1; i <= square; i++ {
        if i == 1 {
            total = 1
        } else {
            total = total * 2
        }
    }
    return total, nil
}

// Total returns the total number of grains on the board
func Total() uint64 {
    var sum uint64
    for i := 1; i <= 64; i++ {
        grains, _ := Square(i)
        sum += grains
    }
    return sum
}