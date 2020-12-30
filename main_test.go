package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Daniel-W-Innes/shopify-challenge-2021/controller"
	"github.com/Daniel-W-Innes/shopify-challenge-2021/proto"
	"github.com/Daniel-W-Innes/shopify-challenge-2021/util"
	"github.com/vitali-fedulov/images"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"
	"image/png"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"strconv"
	"testing"
)

const bufSize = 1024 * 1024
const userID = 123
const imageName = "test"
const testingImage = "test_res/grace_hopper.png"

var lis *bufconn.Listener

func init() {
	err := os.Setenv("BASE_PATH", controller.BasePath)
	if err != nil {
		log.Fatalf("failed to set base path: %v", err)
	}
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	proto.RegisterShopifyAddServer(s, &controller.Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestImageAdd(t *testing.T) {
	data, err := ioutil.ReadFile(testingImage)
	if err != nil {
		t.Error("failed to encode image")
	}
	user := util.User{UserID: userID}
	userBuf, err := json.Marshal(user)
	if err != nil {
		t.Error("failed to marshal user")
	}
	md := metadata.New(map[string]string{util.UserHeader: string(userBuf)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Error("failed to open connection")
	}
	defer conn.Close()

	c := proto.NewShopifyAddClient(conn)
	var header metadata.MD
	response, err := c.ImageAdd(ctx, &proto.Image{ImageData: data, Info: &proto.ImageInfo{Name: imageName}}, grpc.Header(&header))
	if err != nil {
		t.Errorf("error raised by client: %s", err)
	}
	if response.Error != "" {
		t.Errorf("bad response from server %s", response.Error)
	}

	imgFile, err := os.Open(path.Join(controller.BasePath, strconv.Itoa(userID), fmt.Sprintf("%s.png", imageName)))
	if err != nil {
		t.Error("error opening saved file")
	}
	newImageObj, err := png.Decode(imgFile)
	if err != nil {
		t.Error("error decoding saved image")
	}
	imageObj, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		t.Error("error decoding image")
	}
	hashA, imgSizeA := images.Hash(imageObj)
	hashB, imgSizeB := images.Hash(newImageObj)
	if !images.Similar(hashA, hashB, imgSizeA, imgSizeB) {
		t.Error("images are not the same")
	}

}
