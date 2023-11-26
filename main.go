package main

import (
)

func main() {
	// Example trades for stocks from the stocks.go file
	stockTrades := GetStockTrades()

	// Example trades for forex from the forex.go file
	forexTrades := GetForexTrades()

	// Create a trading journal and add the trades
	journal := TradingJournal{
		Trades: append(stockTrades, forexTrades...),
	}

  // Connect to postgres
  Postgres()
  
	// Start the server
	StartServer(8080, journal)
}
