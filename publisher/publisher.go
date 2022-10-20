package publisher

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/stan.go"
)

func Publish() {
	sc, err := stan.Connect("wb_l0", "publisher")

	if err != nil {
		fmt.Printf("ERROR: publisher can't connect to nats: %v\n", err)
		return
	}

	defer sc.Close()

	for i := 0; ; i++ {
		new_order := generateOneOrder()
		json_order, err := json.Marshal(new_order)
		if err != nil {
			fmt.Printf("ERROR: can't marshal order to json: %v\n", err)
			return
		}

		err = sc.Publish("order", json_order)
		if err != nil {
			fmt.Printf("ERROR: can't publish message: %v\n", err)
			return
		}

		time.Sleep(10 * time.Second)
		if i%10 == 0 && i != 0 {
			time.Sleep(10 * time.Minute)
		}
	}
}
