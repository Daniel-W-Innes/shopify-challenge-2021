package controller

import (
	"bytes"
	"fmt"
	"github.com/Daniel-W-Innes/shopify-challenge-2021/proto"
	"github.com/Daniel-W-Innes/shopify-challenge-2021/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"image"
	"image/png"
	"log"
	"os"
	"path"
	"strconv"
)

const BasePath = "img"

type Server struct{}

func (s *Server) ImageAdd(ctx context.Context, imageMsg *proto.Image) (*proto.Response, error) {
	user, err := util.GetUserFromHeader(ctx)
	if err != nil {
		log.Print(err.Error())
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	imageInfo := imageMsg.GetInfo()
	imageObj, _, err := image.Decode(bytes.NewReader(imageMsg.ImageData))
	if err != nil {
		log.Print(err.Error())
		return nil, status.Error(codes.InvalidArgument, "failed to decode image")
	}
	if imageInfo.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "image name can not be empty")
	}
	if user.UserID == 0 {
		return nil, status.Error(codes.InvalidArgument, "bad UserID")
	}
	savePath := path.Join(BasePath, strconv.Itoa(user.UserID), fmt.Sprintf("%s.png", imageInfo.GetName()))
	err = os.MkdirAll(path.Dir(savePath), os.ModePerm)
	if err != nil {
		log.Print(err.Error())
		return nil, status.Error(codes.Internal, "failed to save image")
	}
	log.Printf("path to save image %s", savePath)
	outputFile, err := os.Create(savePath)
	if err != nil {
		log.Print(err.Error())
		return nil, status.Error(codes.Internal, "failed to save image")
	}
	defer outputFile.Close()
	err = png.Encode(outputFile, imageObj)
	if err != nil {
		log.Print(err.Error())
		return nil, status.Error(codes.InvalidArgument, "failed to encode image")
	}
	return &proto.Response{}, nil
}

func (s *Server) ImagesAdd(server proto.ShopifyAdd_ImagesAddServer) error {
	return status.Error(codes.Unimplemented, "not implemented yet")
}

func (s *Server) ImageAddTiles(server proto.ShopifyAdd_ImageAddTilesServer) error {
	return status.Error(codes.Unimplemented, "not implemented yet")
}

func (s *Server) ImageRegisterTiles(ctx context.Context, registration *proto.ImageTileRegistration) (*proto.Response, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (s *Server) ImageRegisterTile(server proto.ShopifyAdd_ImageRegisterTileServer) error {
	return status.Error(codes.Unimplemented, "not implemented yet")
}
