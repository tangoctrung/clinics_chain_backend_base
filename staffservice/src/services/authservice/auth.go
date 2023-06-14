package authservice

import (
	"context"

	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/staffservice/src/pb/authpb"
)

var Set = wire.NewSet(wire.Bind(new(Service), new(ServiceGRPC)), wire.Struct(new(ServiceGRPC), "*"), ConnectClient)

type Service interface {
	CreateNewUser(ctx context.Context, name string, age int) (*authpb.User, error)
}
