package main

import (
  "fmt"
	"html/template"
	"net/http"
)

// Trade represents a trading entry.
type Trade struct {
	Symbol      string
	EntryPrice  float64
	Quantity    uint
	Date        string
	Information string
	ExitPrice   float64
	Fees        float64
	PL          float64
	Realized    float64
}

// IndexPageData represents the data to be rendered in the index template.
type IndexPageData struct {
	Trades []Trade
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    trades := []Trade{
        {
            Symbol:      "AAPL",
            EntryPrice:  150.0,
            Quantity:    10,
            Date:        "2023-11-25",
            Information: "Long",
            ExitPrice:   160.0,
            Fees:        5.0,
            PL:          0.0,
            Realized:    0.0,
        },
        // Add more trades as needed
    }

    data := IndexPageData{Trades: trades}

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/", indexHandler)

    // Serve static files from the "static" directory
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Start the server
    port := 8080
    addr := fmt.Sprintf(":%d", port)

    fmt.Printf("Server running on http://localhost:%d\n", port)

    err := http.ListenAndServe(addr, nil)
    if err != nil {
        fmt.Println("Error starting the server:", err)
    }
}
