package repository

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

func (r *repository) AddProduct(ctx context.Context, pr *pb.BasketProductRequest) (err error) {
	const queryCheck = `select id, products from baskets where user_id = $1`
	row := r.pool.QueryRow(ctx, queryCheck, pr.User.ID)
	var basket models.DBBasket
	err = row.Scan(&basket.ID, &basket.Products)
	var existProducts []models.DBBasketProduct
	if err != nil {
		log.Println("basket not found")
		const queryAdd = `insert into baskets (user_id, products) values ($1, $2)`
		_, err = r.pool.Exec(ctx, queryAdd, pr.User.ID, "[]")
		if err != nil {
			return errors.New("error add product to basket: " + err.Error())
		}

	}
	if &basket.Products != nil {
		err = json.Unmarshal([]byte(basket.Products), &existProducts)
		if err != nil {
			log.Println("Empty products")
			existProducts = append(existProducts, models.DBBasketProduct{ID: pr.Product.ID, Count: 1})
		} else {
			if &existProducts != nil {
				log.Println("Exist products", existProducts)
				found := false
				for index, product := range existProducts {
					if product.ID == pr.Product.ID {
						existProducts[index].Count++
						found = true
					}
				}
				if !found {
					existProducts = append(existProducts, models.DBBasketProduct{ID: pr.Product.ID, Count: 1})
				}
			}
		}
	}
	log.Println("AddProduct:", existProducts)
	j, err := json.Marshal(existProducts)
	if err != nil {
		return errors.New("error marshal product")
	}
	const queryUpdate = `update baskets set products = $1 where user_id = $2`
	_, err = r.pool.Exec(ctx, queryUpdate, j, pr.User.ID)
	return err
}
