package v1

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hugomatus/order-system/productinfo/api/v1"
	//v1 "github.com/hugomatus/order-system/productinfo/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)


const (
	port = ":50051"
)

// server is used to implement ecommerce/product_info.
type server struct {
	productMap map[string]*v1.Product
}

// AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context,
	in *v1.Product) (*v1.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*v1.Product)
	}
	s.productMap[in.Id] = in
	log.Printf("Product %v : %v - Added.", in.Id, in.Name)
	return &v1.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

// GetProduct implements ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *v1.ProductID) (*v1.Product, error) {
	product, exists := s.productMap[in.Value]
	if exists && product != nil {
		log.Printf("Product %v : %v - Retrieved.", product.Id, product.Name)
		return product, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}

func Run() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	v1.RegisterProductInfoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}