package controller

import (
	"github.com/Daniel-W-Innes/shopify-add/proto"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) ImageAdd(ctx context.Context, image *proto.Image) (*proto.Response, error) {
	panic("implement me")
}

func (s *Server) ImagesAdd(server proto.ShopifyAdd_ImagesAddServer) error {
	panic("implement me")
}

func (s *Server) ImageAddTiles(server proto.ShopifyAdd_ImageAddTilesServer) error {
	panic("implement me")
}

func (s *Server) ImageRegisterTiles(server proto.ShopifyAdd_ImageRegisterTilesServer) error {
	panic("implement me")
}
