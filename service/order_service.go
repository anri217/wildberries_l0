package service

import (
	"context"
	"fmt"
	"wb_l0/configs"
	"wb_l0/domain"
	"wb_l0/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderService struct {
	r *repository.OrderRepo
}

func initConnectionPool() *pgxpool.Pool {
	connectionStr := fmt.Sprintf("postgresql://%s:%s@%s/%s", configs.User, configs.Pass, configs.Host, configs.Name)
	pool, err := pgxpool.New(context.Background(), connectionStr)

	if err != nil {
		fmt.Printf("ERROR: can't connect to db: %v", err)
	}

	return pool
}

func NewOrderService() *OrderService {
	pool := initConnectionPool()
	r := repository.NewOrderRepo(pool)
	return &OrderService{r}
}

func (o_service *OrderService) GetOrderByUID(uid *string) (*domain.Order, error) {
	order, err := o_service.r.GetOrderByUID(uid)
	if err != nil {
		return nil, fmt.Errorf("ERROR: can't get order by uid: %v\n", err)
	}
	return order, nil
}

func (o_service *OrderService) AddOrder(order *domain.Order) {
	err := o_service.r.AddOrder(order)
	if err != nil {
		fmt.Printf("ERROR: can't add order: %v\n", err)
		return
	}
}
