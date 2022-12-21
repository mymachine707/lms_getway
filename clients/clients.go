package clients

import (
	"lms/lms_getway/config"
	"lms/lms_getway/protogen/book_service"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Author book_service.AuthorServiceClient

	Book book_service.BookServiceClient

	Category book_service.CategoryServiceClient

	Location book_service.LocationServiceClient

	conns []*grpc.ClientConn
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	connCategory, err := grpc.Dial(cfg.CategoryServiceGrpcHost+cfg.CategoryServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	category := book_service.NewCategoryServiceClient(connCategory)

	connLocation, err := grpc.Dial(cfg.LocationServiceGrpcHost+cfg.LocationServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	location := book_service.NewLocationServiceClient(connLocation)

	connAuthor, err := grpc.Dial(cfg.AuthorServiceGrpcHost+cfg.AuthorServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	author := book_service.NewAuthorServiceClient(connAuthor)

	connBook, err := grpc.Dial(cfg.BookServiceGrpcHost+cfg.BookServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	book := book_service.NewBookServiceClient(connBook)

	conns := make([]*grpc.ClientConn, 0)
	return &GrpcClients{
		Category: category,
		Location: location,
		Author:   author,
		Book:     book,
		conns:    append(conns, connCategory),
	}, nil
}

// Close for disconnection connection another serevice
func (c *GrpcClients) Close() {
	for _, v := range c.conns {
		v.Close()
	}
}
