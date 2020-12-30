package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type User struct {
	Active bool `json:"active"`
	Exp    int  `json:"exp"`
	UserID int  `json:"user_id"`
}

const UserHeader = "x-user"

func GetUserFromHeader(ctx context.Context) (*User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("mo headers in request")
	}
	authHeaders, ok := md[UserHeader]
	if !ok {
		return nil, fmt.Errorf("missing %s header", UserHeader)
	}

	if len(authHeaders) != 1 {
		return nil, fmt.Errorf("more than 1 %s header in request", UserHeader)
	}

	user := User{}
	err := json.Unmarshal([]byte(authHeaders[0]), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
