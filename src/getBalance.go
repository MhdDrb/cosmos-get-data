package src

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func queryState() error {
	myAddress, err := sdk.AccAddressFromBech32("cosmos1u3g4rhmr29qgz5fqtjl98xeua4r7k0t4dhmnre") // the my_validator or recipient address.
	if err != nil {
		return err
	}

	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		//"127.0.0.1:9090", // your gRPC server address.
		"https://rpc.cosmos.network:26657", // your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	// This creates a gRPC client to query the x/bank service.
	bankClient := banktypes.NewQueryClient(grpcConn)
	bankRes, err := bankClient.Balance(
		context.Background(),
		&banktypes.QueryBalanceRequest{Address: myAddress.String(), Denom: "stake"},
	)
	if err != nil {
		return err
	}

	fmt.Println(bankRes.GetBalance()) // Prints the account balance

	return nil
}

func main() {
	if err := queryState(); err != nil {
		panic(err)
	}
}