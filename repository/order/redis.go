package order

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Lutefd/gorders-api/model"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}

func orderIDKey(orderID uint64) string {
	return fmt.Sprintf("order:%d", orderID)
}
func (r *RedisRepo) Insert(ctx context.Context, order model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}
	key := orderIDKey(order.OrderID)
	txn := r.Client.TxPipeline()
	res := txn.SetNX(ctx, key, string(data), 0)
	if res.Err() != nil {
		return fmt.Errorf("failed to insert order: %w", res.Err())
	}
	if err := txn.SAdd(ctx, "orders", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to add order to set: %w", err)
	}
	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to execute transaction: %w", err)
	}
	return nil
}

func (r *RedisRepo) FindByID(ctx context.Context, id uint64) (model.Order, error) {
	key := orderIDKey(id)
	res := r.Client.Get(ctx, key)
	if res.Err() != nil {
		return model.Order{}, fmt.Errorf("failed to find order: %w", res.Err())
	}
	var order model.Order
	err := json.Unmarshal([]byte(res.Val()), &order)
	if err != nil {
		return model.Order{}, fmt.Errorf("failed to decode order: %w", err)
	}
	return order, nil
}

func (r *RedisRepo) Update(ctx context.Context, order model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}
	key := orderIDKey(order.OrderID)
	res := r.Client.SetXX(ctx, key, string(data), 0)
	if res.Err() != nil {
		return fmt.Errorf("failed to update order: %w", res.Err())
	}
	return nil
}

func (r *RedisRepo) Delete(ctx context.Context, id uint64) error {
	key := orderIDKey(id)
	txn := r.Client.TxPipeline()
	res := txn.Del(ctx, key)
	if res.Err() != nil {
		return fmt.Errorf("failed to delete order: %w", res.Err())
	}
	if err := txn.SRem(ctx, "orders", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to remove order from set: %w", err)
	}
	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to execute transaction: %w", err)
	}
	return nil
}

type FindAllPage struct {
	Size   uint64
	Offset uint64
}
type FindAllResult struct {
	Orders []model.Order
	Cursor uint64
}

func (r *RedisRepo) List(ctx context.Context, page FindAllPage) (FindAllResult, error) {
	res := r.Client.SScan(ctx, "orders", page.Offset, "*", int64(page.Size))
	keys, cursor, err := res.Result()
	if err != nil {
		return FindAllResult{}, fmt.Errorf("failed to scan orders: %w", err)
	}
	if len(keys) == 0 {
		return FindAllResult{Orders: []model.Order{}, Cursor: uint64(cursor)}, nil
	}
	xs, err := r.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return FindAllResult{}, fmt.Errorf("failed to get orders: %w", err)
	}
	orders := make([]model.Order, 0, len(xs))
	for _, x := range xs {
		var order model.Order
		err := json.Unmarshal([]byte(x.(string)), &order)
		if err != nil {
			return FindAllResult{}, fmt.Errorf("failed to decode order: %w", err)
		}
		orders = append(orders, order)
	}
	return FindAllResult{Orders: orders, Cursor: uint64(cursor)}, nil
}
