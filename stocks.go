package main

import "time"

// GetStockTrades returns example stock trades
func GetStockTrades() []Trades {
	return []Trades{
		{
			TradeID:    "123456",
			Date:       time.Now(),
			Symbol:     "AAPL",
			Side:       "Long",
			EntryTime:  "09:30 AM",
			EntryPrice: 150.0,
			Size:       10,
			StopLoss:   145.0,
			ExitPrice:  160.0,
			ExitTime:   "04:00 PM",
			ExitDate:   "2023-11-25",
			Commission: 5.0,
			PL:         1000.0,
			WinLoss:    "Win",
			RRR:        "2:1",
			Setup:      "Bullish Setup",
			Mistakes:   "None",
			Lessons:    "Learned to set better stop losses",
		},
		// Add more stock trades as needed
	}
}
