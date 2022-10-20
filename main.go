package main

import (
	"context"
	"fmt"
	"wb_l0/configs"
	"wb_l0/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

func initConnectionPool() *pgxpool.Pool {
	connectionStr := fmt.Sprintf("postgresql://%s:%s@%s/%s", configs.Name, configs.Pass, configs.Host, configs.Name)
	pool, err := pgxpool.New(context.Background(), connectionStr)

	if err != nil {
		fmt.Printf("ERROR: can't connect to db: %v", err)
	}

	return pool
}

func main() {
	// sc, err := stan.Connect("wb_l0", "publisher")

	// if err != nil {
	// 	fmt.Println("Error")
	// 	return
	// }

	// defer sc.Close()

	// for i := 0; ; i++ {
	// 	sc.Publish("hello", []byte("Hello, "+strconv.Itoa(i)))
	// 	time.Sleep(2 * time.Second)
	// }

	// order_one := domain.Order{
	// 	OrderUID:    "k",
	// 	TrackNumber: "hjk",
	// 	Entry:       "hjk",
	// 	Delivery: domain.Delivery{
	// 		Name:    "ghj",
	// 		Phone:   "+89999999999",
	// 		Zip:     "hj",
	// 		City:    "hj",
	// 		Address: "j",
	// 		Region:  "hjk",
	// 		Email:   "jh",
	// 	},
	// 	Payment: domain.Payment{
	// 		Transaction:  "hj",
	// 		RequestID:    "hjk",
	// 		Currency:     "hjk",
	// 		Provider:     "hjk",
	// 		Amount:       1,
	// 		PaymentDt:    12,
	// 		Bank:         "hj",
	// 		DeliveryCost: 123,
	// 		GoodsTotal:   123,
	// 		CustomFee:    121,
	// 	},
	// 	Items: []domain.Item{
	// 		{
	// 			ChrtID:      12,
	// 			TrackNumber: "hjk",
	// 			Price:       124,
	// 			Rid:         "hj",
	// 			Name:        "hjk",
	// 			Sale:        123,
	// 			Size:        "hjk",
	// 			TotalPrice:  124,
	// 			NmID:        2134,
	// 			Brand:       "hjk",
	// 			Status:      1234,
	// 		},
	// 	},
	// 	Locale:            "jbm",
	// 	InternalSignature: "bn",
	// 	CustomerID:        "hj",
	// 	DeliveryService:   "hj",
	// 	Shardkey:          "hjk",
	// 	SmID:              123,
	// 	DateCreated:       time.Now(),
	// 	OofShard:          "hjk",
	// }

	pool := initConnectionPool()

	order_repo := repository.NewOrderRepo(pool)

	// err = order_repo.AddOrder(&order_one)
	// if err != nil {
	// 	fmt.Println("ERROR: can't add order")
	// }

	first_key := "h"
	second_key := "g"
	third_key := "k"

	if val, err := order_repo.GetOrderByUID(&first_key); err != nil {
		fmt.Println("ERROR: can't add order")
	} else {
		fmt.Printf("Order: %v\n", *val)
	}

	if val, err := order_repo.GetOrderByUID(&second_key); err != nil {
		fmt.Println("ERROR: can't add order")
	} else {
		fmt.Printf("Order: %v\n", *val)
	}

	if val, err := order_repo.GetOrderByUID(&third_key); err != nil {
		fmt.Println("ERROR: can't add order")
	} else {
		fmt.Printf("Order: %v\n", *val)
	}
}
