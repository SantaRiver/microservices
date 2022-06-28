package app

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/ptypes/timestamp"
	"log"
	"ozon/internal/basket/internal/client"
	"ozon/internal/models"
	pb "ozon/internal/pkg/api"
)

func (s *Server) GetBasket(ctx context.Context, user *pb.User) (br *pb.BasketResponse, err error) {
	dbBasket, err := s.repo.GetBasket(ctx, user)
	if err != nil {
		return &pb.BasketResponse{}, err
	}
	var unmarshalProducts []models.DBBasketProduct
	err = json.Unmarshal([]byte(dbBasket.Products), &unmarshalProducts)
	if err != nil {
		return &pb.BasketResponse{}, errors.New("unmarshal error")
	}
	var products []*pb.Product
	var count []*pb.Count
	var totalPrice float32
	log.Println("unmarshalProducts", unmarshalProducts)
	for _, product := range unmarshalProducts {
		pbProduct := pb.Product{ID: product.ID}
		productResponse, err := client.GetProduct(ctx, &pbProduct)
		log.Println("productResponse", productResponse)
		if err != nil {
			return &pb.BasketResponse{}, err
		}
		count = append(count, &pb.Count{ID: product.ID, Count: product.Count})
		totalPrice += productResponse.Product.Price * float32(product.Count)
		products = append(products, productResponse.Product)
	}
	pbUser := pb.User{ID: dbBasket.UserID}
	userResponse, err := client.GetUser(ctx, &pbUser)
	if err != nil {
		return &pb.BasketResponse{}, err
	}
	basket := &pb.Basket{
		ID:         dbBasket.ID,
		User:       userResponse.User,
		Products:   products,
		Counter:    count,
		TotalPrice: totalPrice,
		CreatedAt:  &timestamp.Timestamp{Seconds: dbBasket.CreatedAt.Time.Unix()},
		UpdatedAt:  &timestamp.Timestamp{Seconds: dbBasket.UpdatedAt.Time.Unix()},
	}
	basketResponse := &pb.BasketResponse{
		Basket: basket,
	}
	return basketResponse, nil
	/**/
}
