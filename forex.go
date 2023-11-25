package main

import "time"

// GetForexTrades returns example forex trades
func GetForexTrades() []Trades {
	return []Trades{
		{
			TradeID:    "789012",
			Date:       time.Now(),
			Symbol:     "EURUSD",
			Side:       "Short",
			EntryTime:  "12:00 PM",
			EntryPrice: 1.1200,
			Size:       100000,
			StopLoss:   1.1250,
			ExitPrice:  1.1100,
			ExitTime:   "03:30 PM",
			ExitDate:   "2023-11-25",
			Commission: 10.0,
			PL:         500.0,
			WinLoss:    "Win",
			RRR:        "3:1",
			Setup:      "Bearish Setup",
			Mistakes:   "Entered too early",
			Lessons:    "Wait for confirmation",
		},
		// Add more forex trades as needed
	}
}
