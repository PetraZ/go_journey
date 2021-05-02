package main

import (
	"log"
	"time"
)

//Response is a type of http handler response
type Response struct {
	Status  string
	Message string
}

func handleCustomerTransaction(customerID int, isDone chan bool) {
	log.Println("ID:", customerID, "Received booking request")
	status := make(chan string)
	transactionSuccess := make(chan bool)

	timeout := time.After(15 * time.Second)

	go func() {
		status <- " Seat selection going on..."
		// Use customer details to make DB queries, third party API/Service call
		time.Sleep(5 * time.Second)
		status <- " Making payments from bank..."
		time.Sleep(5 * time.Second)
		// Everything looks good. Notify customer
		transactionSuccess <- true
		defer close(transactionSuccess)
	}()

	for {
		select {
		case update := <-status:
			log.Println("ID:", customerID, update)
		case <-timeout:
			close(status)
			log.Println("Operation timed out!")
			isDone <- false
			return
		case <-transactionSuccess:
			log.Println("ID:", customerID, "Successfully booked ticket!")
			isDone <- true
			return
		}
	}
}
