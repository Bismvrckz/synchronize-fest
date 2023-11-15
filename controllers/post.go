package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Ticket struct {
	Amount         int64  `json:"amount"`
	InvoiceNumber  string `json:"invoice_number"`
	PaymentDueDate int64  `json:"payment_due_date"`
}

var BaseURL = "https://sandbox.doku.com/p-link/p/0QxPVRz"

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	// Contoh data ticket
	ticket := Ticket{
		Amount:         100000,         // Misalnya, Rp 100.000
		InvoiceNumber:  "INV123456789", // Nomor invoice
		PaymentDueDate: 1672531199,     // Contoh: Timestamp untuk 1 Januari 2024
	}

	// Mengubah data ticket ke format JSON
	jsonPayload, err := json.Marshal(ticket)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error preparing JSON payload", http.StatusInternalServerError)
		return
	}

	// Mengirim permintaan POST ke API DOKU
	response, err := http.Post(BaseURL+"/create-ticket", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Print(err)
		http.Error(w, "Error sending POST request", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Membaca respons dari API DOKU
	var result map[string]interface{}
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&result); err != nil {
		log.Print(err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Mengirim respons ke klien
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
