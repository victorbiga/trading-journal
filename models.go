package main

import "time"

// Trades represents a trading entry in the journal
type Trades struct {
	TradeID     string
	Date        time.Time
	Symbol      string
	Side        string
	EntryTime   string
	EntryPrice  float64
	Size        int
	StopLoss    float64
	ExitPrice   float64
	ExitTime    string
	ExitDate    string
	Commission  float64
	PL          float64
	WinLoss     string
	RRR         string
	Setup       string
	Mistakes    string
	Lessons     string
}

// TradingJournal represents the overall trading journal
type TradingJournal struct {
	Trades []Trades
}
