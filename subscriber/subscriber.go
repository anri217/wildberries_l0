package subscriber

import (
	"encoding/json"
	"fmt"
	"sync"
	"wb_l0/domain"
	"wb_l0/service"

	"github.com/nats-io/stan.go"
)

func Subscribe(o_service *service.OrderService) {
	sub, err := stan.Connect("wb_l0", "sub")

	if err != nil {
		fmt.Printf("ERROR: subscriber can't connect to nats: %v\n", err)
		return
	}

	defer sub.Close()

	sub.Subscribe("order", func(m *stan.Msg) {
		var new_order domain.Order
		err := json.Unmarshal(m.Data, &new_order)
		if err != nil {
			fmt.Printf("ERROR: can't unmarshal json with new order in subscriber: %v\n", err)
			return
		}
		o_service.AddOrder(&new_order)
	})

	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}
