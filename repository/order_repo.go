package repository

import (
	"context"
	"fmt"
	"wb_l0/domain"
	"wb_l0/migrate"
	"wb_l0/repository/pg_model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepo struct {
	pool   *pgxpool.Pool
	orders map[string]*domain.Order
}

func NewOrderRepo(pool *pgxpool.Pool) *OrderRepo {
	err := migrate.TryInitTables(pool)
	if err != nil {
		fmt.Println(err)
	}

	all_orders, err := getAllOrders(pool)
	if err != nil {
		fmt.Printf("ERROR: can't get all orders from db: %v", err)
	}

	return &OrderRepo{
		pool:   pool,
		orders: all_orders,
	}
}

func getAllOrders(pool *pgxpool.Pool) (map[string]*domain.Order, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	order_slice, err := getAllRowsFromOrdersTable(conn)
	if err != nil {
		return nil, err
	}

	delivery_slice, err := getAllRowsFromDeliveriesTable(conn)
	if err != nil {
		return nil, err
	}

	payment_slice, err := getAllRowsFromPaymentsTable(conn)
	if err != nil {
		return nil, err
	}

	item_slice, err := getAllRowsFromItemsTable(conn)
	if err != nil {
		return nil, err
	}

	var all_orders = make(map[string]*domain.Order)

	for _, pg_order := range order_slice {
		cur_order := domain.Order{
			OrderUID:          pg_order.OrderUID,
			TrackNumber:       pg_order.TrackNumber,
			Entry:             pg_order.Entry,
			Delivery:          domain.Delivery{},
			Payment:           domain.Payment{},
			Items:             []domain.Item{},
			Locale:            pg_order.Locale,
			InternalSignature: pg_order.InternalSignature,
			CustomerID:        pg_order.CustomerID,
			DeliveryService:   pg_order.DeliveryService,
			Shardkey:          pg_order.Shardkey,
			SmID:              pg_order.SmID,
			DateCreated:       pg_order.DateCreated,
			OofShard:          pg_order.OofShard,
		}
		all_orders[pg_order.OrderUID] = &cur_order
	}

	for _, pg_delivery := range delivery_slice {
		cur_delivery := domain.Delivery{
			Name:    pg_delivery.Name,
			Phone:   pg_delivery.Phone,
			Zip:     pg_delivery.Zip,
			City:    pg_delivery.City,
			Address: pg_delivery.Address,
			Region:  pg_delivery.Region,
			Email:   pg_delivery.Email,
		}
		all_orders[pg_delivery.OrderUID].Delivery = cur_delivery
	}

	for _, pg_payment := range payment_slice {
		cur_payment := domain.Payment{
			Transaction:  pg_payment.Transaction,
			RequestID:    pg_payment.RequestID,
			Currency:     pg_payment.Currency,
			Provider:     pg_payment.Provider,
			Amount:       pg_payment.Amount,
			PaymentDt:    pg_payment.PaymentDt,
			Bank:         pg_payment.Bank,
			DeliveryCost: pg_payment.DeliveryCost,
			GoodsTotal:   pg_payment.GoodsTotal,
			CustomFee:    pg_payment.CustomFee,
		}
		all_orders[pg_payment.OrderUID].Payment = cur_payment
	}

	for _, pg_item := range item_slice {
		cur_item := domain.Item{
			ChrtID:      pg_item.ChrtID,
			TrackNumber: pg_item.TrackNumber,
			Price:       pg_item.Price,
			Rid:         pg_item.Rid,
			Name:        pg_item.Name,
			Sale:        pg_item.Sale,
			Size:        pg_item.Size,
			TotalPrice:  pg_item.TotalPrice,
			NmID:        pg_item.NmID,
			Brand:       pg_item.Brand,
			Status:      pg_item.Status,
		}
		all_orders[pg_item.OrderUID].Items = append(all_orders[pg_item.OrderUID].Items, cur_item)
	}

	return all_orders, nil
}

func getAllRowsFromItemsTable(conn *pgxpool.Conn) ([]*pg_model.PGItem, error) {
	item_query := "SELECT * FROM items"
	items, err := conn.Query(context.Background(), item_query)
	if err != nil {
		return nil, err
	}
	defer items.Close()

	var item_slice []*pg_model.PGItem
	for items.Next() {
		var cur_item pg_model.PGItem
		err = items.Scan(&cur_item.ItemID, &cur_item.OrderUID, &cur_item.ChrtID, &cur_item.TrackNumber, &cur_item.Price, &cur_item.Rid, &cur_item.Name,
			&cur_item.Sale, &cur_item.Size, &cur_item.TotalPrice, &cur_item.NmID, &cur_item.Brand, &cur_item.Status)
		if err != nil {
			return nil, err
		}
		item_slice = append(item_slice, &cur_item)
	}
	if err := items.Err(); err != nil {
		return nil, err
	}

	return item_slice, nil
}

func getAllRowsFromPaymentsTable(conn *pgxpool.Conn) ([]*pg_model.PGPayment, error) {
	payment_query := "SELECT * FROM payments"
	payments, err := conn.Query(context.Background(), payment_query)
	if err != nil {
		return nil, err
	}
	defer payments.Close()

	var payment_slice []*pg_model.PGPayment
	for payments.Next() {
		var cur_payment pg_model.PGPayment
		err = payments.Scan(&cur_payment.OrderUID, &cur_payment.Transaction, &cur_payment.RequestID, &cur_payment.Currency, &cur_payment.Provider,
			&cur_payment.Amount, &cur_payment.PaymentDt, &cur_payment.Bank, &cur_payment.DeliveryCost, &cur_payment.GoodsTotal, &cur_payment.CustomFee)
		if err != nil {
			return nil, err
		}
		payment_slice = append(payment_slice, &cur_payment)
	}
	if err := payments.Err(); err != nil {
		return nil, err
	}

	return payment_slice, nil
}

func getAllRowsFromDeliveriesTable(conn *pgxpool.Conn) ([]*pg_model.PGDelivery, error) {
	delivery_query := "SELECT * FROM deliveries"
	deliveries, err := conn.Query(context.Background(), delivery_query)
	if err != nil {
		return nil, err
	}
	defer deliveries.Close()

	var delivery_slice []*pg_model.PGDelivery
	for deliveries.Next() {
		var cur_delivery pg_model.PGDelivery
		err = deliveries.Scan(&cur_delivery.OrderUID, &cur_delivery.Name, &cur_delivery.Phone, &cur_delivery.Zip, &cur_delivery.City,
			&cur_delivery.Address, &cur_delivery.Region, &cur_delivery.Email)
		if err != nil {
			return nil, err
		}
		delivery_slice = append(delivery_slice, &cur_delivery)
	}
	if err := deliveries.Err(); err != nil {
		return nil, err
	}

	return delivery_slice, nil
}

func getAllRowsFromOrdersTable(conn *pgxpool.Conn) ([]*pg_model.PGOrder, error) {
	order_query := "SELECT * FROM orders"
	orders, err := conn.Query(context.Background(), order_query)
	if err != nil {
		return nil, err
	}
	defer orders.Close()

	var order_slice []*pg_model.PGOrder
	for orders.Next() {
		var cur_order pg_model.PGOrder
		err = orders.Scan(&cur_order.OrderUID, &cur_order.TrackNumber, &cur_order.Entry, &cur_order.Locale, &cur_order.InternalSignature, &cur_order.CustomerID,
			&cur_order.DeliveryService, &cur_order.Shardkey, &cur_order.SmID, &cur_order.DateCreated, &cur_order.OofShard)
		if err != nil {
			return nil, err
		}
		order_slice = append(order_slice, &cur_order)
	}
	if err := orders.Err(); err != nil {
		return nil, err
	}

	return order_slice, nil
}

func (r *OrderRepo) GetOrderByUID(uid *string) (*domain.Order, error) {
	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	if v, found := r.orders[*uid]; !found {
		return nil, fmt.Errorf("ERROR: can't find order with this UID")
	} else {
		return v, nil
	}
}

func (r *OrderRepo) AddOrder(order *domain.Order) error {
	if _, found := r.orders[order.OrderUID]; found {
		return fmt.Errorf("ERROR: order with same UID is already added")
	} else {
		r.orders[order.OrderUID] = order
	}

	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	delivery := &order.Delivery
	payment := &order.Payment
	items := &order.Items

	order_query := "INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, " +
		"customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"

	delivery_query := "INSERT INTO deliveries (order_uid, name, phone, zip, city, address, region, email) values ($1, $2, $3, $4, $5, $6, $7, $8)"

	payment_query := "INSERT INTO payments (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) " +
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"

	item_query := "INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) " +
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"

	_, err = conn.Exec(context.Background(), order_query, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID,
		order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard)

	if err != nil {
		return fmt.Errorf("ERROR when insert into orders: %v", err)
	}

	_, err = conn.Exec(context.Background(), delivery_query, order.OrderUID, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)

	if err != nil {
		return fmt.Errorf("ERROR when insert into deliveries: %v", err)
	}

	_, err = conn.Exec(context.Background(), payment_query, order.OrderUID, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDt,
		payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)

	if err != nil {
		return fmt.Errorf("ERROR when insert into payment: %v", err)
	}

	for _, item := range *items {
		_, err = conn.Exec(context.Background(), item_query, order.OrderUID, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice,
			item.NmID, item.Brand, item.Status)

		if err != nil {
			return fmt.Errorf("ERROR when insert into items: %v", err)
		}
	}

	return nil
}
