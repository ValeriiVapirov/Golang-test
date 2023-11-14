package main

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

func main() {
	http.HandleFunc("/fulfillOrder", fulfillOrderHandler)
	http.ListenAndServe(":8080", nil)
}

func fulfillOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderQuantity, err := strconv.Atoi(r.URL.Query().Get("orderQuantity"))
	if err != nil {
		http.Error(w, "Invalid value for orderQuantity", http.StatusBadRequest)
		return
	}

	numberOfPacks := calculateNumberOfPacks(orderQuantity)

	fmt.Fprintf(w, "%d", numberOfPacks)
}


func calculateNumberOfPacks(orderQuantity int) int {
	packSizes := []int{250, 500, 1000, 2000, 5000}
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	smallestPack := packSizes[len(packSizes)-1]
	numberOfPacks := 0

	for _, packSize := range packSizes {
		packCount := orderQuantity / packSize
		numberOfPacks += packCount
		orderQuantity -= packCount * packSize

		if packSize == smallestPack && packCount == 0 && orderQuantity > 0 && orderQuantity < smallestPack {
			numberOfPacks++
		}
	}

	return numberOfPacks
}