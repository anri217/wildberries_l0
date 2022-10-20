package publisher

import (
	"time"
	"wb_l0/domain"

	"github.com/google/uuid"
)

func generateOneOrder() *domain.Order {
	id := uuid.New()

	order := domain.Order{
		OrderUID:    id.String(),
		TrackNumber: "789-789",
		Entry:       "hjk",
		Delivery: domain.Delivery{
			Name:    "ghj",
			Phone:   "+79999999999",
			Zip:     "hj",
			City:    "hj",
			Address: "j",
			Region:  "hjk",
			Email:   "jh",
		},
		Payment: domain.Payment{
			Transaction:  "hj",
			RequestID:    "hjk",
			Currency:     "hjk",
			Provider:     "hjk",
			Amount:       1,
			PaymentDt:    1,
			Bank:         "hj",
			DeliveryCost: 1,
			GoodsTotal:   1,
			CustomFee:    1,
		},
		Items: []domain.Item{
			{
				ChrtID:      1,
				TrackNumber: "hjk",
				Price:       1,
				Rid:         "hj",
				Name:        "hjk",
				Sale:        1,
				Size:        "hjk",
				TotalPrice:  1,
				NmID:        21454,
				Brand:       "hjk",
				Status:      1234,
			},
		},
		Locale:            "jbm",
		InternalSignature: "bn",
		CustomerID:        "hj",
		DeliveryService:   "hj",
		Shardkey:          "hjk",
		SmID:              123,
		DateCreated:       time.Now(),
		OofShard:          "hjk",
	}

	return &order
}
