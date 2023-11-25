// utils.go
package main

import (
	"strconv"
	"sync"
)

var counterMutex sync.Mutex
var counter int

// GetNextTradeID increments the counter and returns the new TradeID as a string.
func GetNextTradeID() string {
	counterMutex.Lock()
	defer counterMutex.Unlock()

	counter++
	return strconv.Itoa(counter)
}
