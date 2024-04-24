// Code generated by proroc-gen-graphql, DO NOT EDIT.
package greeter

import (
	"context"

	"github.com/envelope-org/grpc-graphql-gateway/runtime"
	"github.com/graphql-go/graphql"
	"google.golang.org/grpc"
)

var (
	gql__type_HelloRequest   *graphql.Object // message HelloRequest in greeter.proto
	gql__type_HelloReply     *graphql.Object // message HelloReply in greeter.proto
	gql__type_GoodbyeRequest *graphql.Object // message GoodbyeRequest in greeter.proto
	gql__type_GoodbyeReply   *graphql.Object // message GoodbyeReply in greeter.proto
)

func Gql__type_HelloRequest() *graphql.Object {
	if gql__type_HelloRequest == nil {
		gql__type_HelloRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Greeter_Type_HelloRequest",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.String),
					Description: `Below line means the "name" field is required in GraphQL argument`,
				},
			},
		})
	}
	return gql__type_HelloRequest
}

func Gql__type_HelloReply() *graphql.Object {
	if gql__type_HelloReply == nil {
		gql__type_HelloReply = graphql.NewObject(graphql.ObjectConfig{
			Name: "Greeter_Type_HelloReply",
			Fields: graphql.Fields{
				"message": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_HelloReply
}

func Gql__type_GoodbyeRequest() *graphql.Object {
	if gql__type_GoodbyeRequest == nil {
		gql__type_GoodbyeRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Greeter_Type_GoodbyeRequest",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.String),
					Description: `Below line means the "name" field is required in GraphQL argument`,
				},
			},
		})
	}
	return gql__type_GoodbyeRequest
}

func Gql__type_GoodbyeReply() *graphql.Object {
	if gql__type_GoodbyeReply == nil {
		gql__type_GoodbyeReply = graphql.NewObject(graphql.ObjectConfig{
			Name: "Greeter_Type_GoodbyeReply",
			Fields: graphql.Fields{
				"message": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_GoodbyeReply
}

// graphql__resolver_Greeter is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_Greeter struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_Greeter) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
	// If x.conn is not nil, user injected their own connection
	if x.conn != nil {
		return x.conn, func() {}, nil
	}

	// Otherwise, this handler opens connection with specified host
	conn, err := grpc.DialContext(ctx, x.host, x.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *graphql__resolver_Greeter) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"hello": &graphql.Field{
			Type: Gql__type_HelloReply(),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					Description:  `Below line means the "name" field is required in GraphQL argument`,
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req *HelloRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, err
				}
				client := NewGreeterClient(conn)
				resp, err := client.SayHello(p.Context, req)
				if err != nil {
					return nil, err
				}
				return resp, nil
			},
		},
		"goodbye": &graphql.Field{
			Type: Gql__type_GoodbyeReply(),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					Description:  `Below line means the "name" field is required in GraphQL argument`,
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req *GoodbyeRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, err
				}
				client := NewGreeterClient(conn)
				resp, err := client.SayGoodbye(p.Context, req)
				if err != nil {
					return nil, err
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_Greeter) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterGreeterGraphqlHandler with *grpc.ClientConn manually.
func RegisterGreeterGraphql(mux *runtime.ServeMux) error {
	return RegisterGreeterGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
// service Greeter {
//    option (graphql.service) = {
//        host: "host:port"
//        insecure: true or false
//    };
//
//    ...with RPC definitions
// }
func RegisterGreeterGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(&graphql__resolver_Greeter{
		conn: conn,
		host: "localhost:50051",
		dialOptions: []grpc.DialOption{
			grpc.WithInsecure(),
		},
	})
}
