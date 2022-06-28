package repository

import (
	"context"
	"encoding/json"
	"errors"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

func (r *repository) DeleteProduct(ctx context.Context, pr *pb.BasketProductRequest) (err error) {
	const queryCheck = `select id, products from baskets where user_id = $1`
	row := r.pool.QueryRow(ctx, queryCheck, pr.User.ID)
	var basket models.DBBasket
	err = row.Scan(&basket.ID, &basket.Products)
	var existProducts []models.DBBasketProduct
	if err != nil {
		if err != nil {
			return errors.New("basket not found")
		}

	}
	if &basket.Products != nil {
		err = json.Unmarshal([]byte(basket.Products), &existProducts)
		if err != nil {
			return errors.New("basket is empty")
		} else {
			if &existProducts != nil {
				found := false
				for index, product := range existProducts {
					if product.ID == pr.Product.ID {
						found = true
						if existProducts[index].Count > 1 {
							existProducts[index].Count--
						} else {
							existProducts = append(existProducts[:index], existProducts[index+1:]...)
						}
					}
				}
				if !found {
					return errors.New("product not found")
				}
			}
		}
	}

	j, err := json.Marshal(existProducts)
	if err != nil {
		return errors.New("error marshal product")
	}

	const queryUpdate = `update baskets set products = $1 where user_id = $2`
	_, err = r.pool.Exec(ctx, queryUpdate, j, pr.User.ID)

	return err
}
