package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var prevOrderID int

// Item represents the model for an item in the order
type Item struct {
	ItemID      string `json:"itemID"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

// Order represents the model for an order
type Order struct {
	OrderID      int       `json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}

var OrderList map[int]Order

func init() {
	OrderList = make(map[int]Order, 10)
}

func main() {

	//API that does orders

	router := mux.NewRouter()

	router.HandleFunc("/orders", createOrder).Methods("POST")

	router.HandleFunc("/order/{orderID}", getOrder).Methods("GET")

	router.HandleFunc("/order/orders", getAllOrders).Methods("GET")

	// router.HandleFunc("/order/{orderID}",updateOrder).Methods("UPDATE")

	router.HandleFunc("/order/{orderID}", deleteOrder).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8089", router))

}

func createOrder(Resp http.ResponseWriter, req *http.Request) {

	var order Order

	json.NewDecoder(req.Body).Decode(&order)

	prevOrderID++

	order.OrderID = prevOrderID
	// if order.OrderID not in OrderList{
	OrderList[order.OrderID] = order

	// }

	Resp.Header().Set("Content-Type", "application/json")

	json.NewEncoder(Resp).Encode(order)
	Resp.Write([]byte("Your Order num is " + strconv.Itoa(order.OrderID)))

}

// func updateOrder(Resp http.ResponseWriter,req *http.Request){

// }

func getOrder(Resp http.ResponseWriter, req *http.Request) {
	Resp.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	orderId := params["orderID"]

	orderID, _ := strconv.Atoi(orderId)

	json.NewEncoder(Resp).Encode(OrderList[orderID])

}

func getAllOrders(Resp http.ResponseWriter, req *http.Request) {

	Resp.Header().Set("Content-Type", "application/json")
	for i := range OrderList {
		log.Println(i)
		json.NewEncoder(Resp).Encode(i)

	}

}

func deleteOrder(Resp http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	orderId := params["orderID"]
	orderID, _ := strconv.Atoi(orderId)
	delete(OrderList, orderID)
	Resp.WriteHeader(http.StatusNoContent)

}
