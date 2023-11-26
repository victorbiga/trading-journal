package main

import (
	"fmt"
	"html/template"
	"net/http"
  "time"
)

// StartServer starts the HTTP server
func StartServer(port int, journal TradingJournal) {
	// Handle HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "templates/index.html", journal)
	})

	// Handle form submission to add trades
	http.HandleFunc("/add-trade", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Parse form values
			symbol := r.FormValue("symbol")
			// Parse other trade fields as needed

			// Create a new trade
			newTrade := Trades{
				TradeID: GetNextTradeID(),
				Date:    time.Now(),
				Symbol:  symbol,
				// Set other trade fields based on form values
			}

			// Add the new trade to the trading journal
			journal.Trades = append(journal.Trades, newTrade)
		}

		// Redirect back to the home page after adding the trade
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Server running on http://localhost:%d\n", port)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

// renderTemplate renders the HTML template with the provided data
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
