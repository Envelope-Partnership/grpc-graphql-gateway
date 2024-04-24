// Code generated by proroc-gen-graphql, DO NOT EDIT.
package starwars

import (
	"context"

	"github.com/envelope-org/grpc-graphql-gateway/runtime"
	"github.com/graphql-go/graphql"
	"google.golang.org/grpc"
)

var (
	gql__enum_Type               *graphql.Enum      // enum Type in starwars/starwars.proto
	gql__enum_Episode            *graphql.Enum      // enum Episode in starwars/starwars.proto
	gql__interface_Character     *graphql.Interface // message Character in starwars/starwars.proto
	gql__type_ListHumansResponse *graphql.Object    // message ListHumansResponse in starwars/starwars.proto
	gql__type_ListDroidsResponse *graphql.Object    // message ListDroidsResponse in starwars/starwars.proto
	gql__type_GetHumanRequest    *graphql.Object    // message GetHumanRequest in starwars/starwars.proto
	gql__type_GetHeroRequest     *graphql.Object    // message GetHeroRequest in starwars/starwars.proto
	gql__type_GetDroidRequest    *graphql.Object    // message GetDroidRequest in starwars/starwars.proto
	gql__type_Character          *graphql.Object    // message Character in starwars/starwars.proto
)

func Gql__enum_Type() *graphql.Enum {
	if gql__enum_Type == nil {
		gql__enum_Type = graphql.NewEnum(graphql.EnumConfig{
			Name: "Starwars_Enum_Type",
			Values: graphql.EnumValueConfigMap{
				"HUMAN": &graphql.EnumValueConfig{
					Value: 0,
				},
				"DROID": &graphql.EnumValueConfig{
					Value: 1,
				},
			},
		})
	}
	return gql__enum_Type
}

func Gql__enum_Episode() *graphql.Enum {
	if gql__enum_Episode == nil {
		gql__enum_Episode = graphql.NewEnum(graphql.EnumConfig{
			Name: "Starwars_Enum_Episode",
			Values: graphql.EnumValueConfigMap{
				"_": &graphql.EnumValueConfig{
					Value: 0,
				},
				"NEWHOPE": &graphql.EnumValueConfig{
					Value: 1,
				},
				"EMPIRE": &graphql.EnumValueConfig{
					Value: 2,
				},
				"JEDI": &graphql.EnumValueConfig{
					Value: 3,
				},
			},
		})
	}
	return gql__enum_Episode
}

func Gql__interface_Character() *graphql.Interface {
	if gql__interface_Character == nil {
		gql__interface_Character = graphql.NewInterface(graphql.InterfaceConfig{
			Name: "Starwars_Interface_Character",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"appears_in": &graphql.Field{
					Type: graphql.NewList(Gql__enum_Episode()),
				},
				"home_planet": &graphql.Field{
					Type: graphql.String,
				},
				"primary_function": &graphql.Field{
					Type: graphql.String,
				},
				"type": &graphql.Field{
					Type: Gql__enum_Type(),
				},
			},
			ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
				return Gql__type_Character()
			},
		})
	}
	return gql__interface_Character
}

func Gql__type_ListHumansResponse() *graphql.Object {
	if gql__type_ListHumansResponse == nil {
		gql__type_ListHumansResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Starwars_Type_ListHumansResponse",
			Fields: graphql.Fields{
				"humans": &graphql.Field{
					Type: graphql.NewList(Gql__type_Character()),
				},
			},
		})
	}
	return gql__type_ListHumansResponse
}

func Gql__type_ListDroidsResponse() *graphql.Object {
	if gql__type_ListDroidsResponse == nil {
		gql__type_ListDroidsResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Starwars_Type_ListDroidsResponse",
			Fields: graphql.Fields{
				"droids": &graphql.Field{
					Type: graphql.NewList(Gql__type_Character()),
				},
			},
		})
	}
	return gql__type_ListDroidsResponse
}

func Gql__type_GetHumanRequest() *graphql.Object {
	if gql__type_GetHumanRequest == nil {
		gql__type_GetHumanRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Starwars_Type_GetHumanRequest",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: `id of the human`,
				},
			},
		})
	}
	return gql__type_GetHumanRequest
}

func Gql__type_GetHeroRequest() *graphql.Object {
	if gql__type_GetHeroRequest == nil {
		gql__type_GetHeroRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Starwars_Type_GetHeroRequest",
			Fields: graphql.Fields{
				"episode": &graphql.Field{
					Type:        Gql__enum_Episode(),
					Description: `If omitted, returns the hero of the whope saga. If provided, returns the hero of that particular episode.`,
				},
			},
		})
	}
	return gql__type_GetHeroRequest
}

func Gql__type_GetDroidRequest() *graphql.Object {
	if gql__type_GetDroidRequest == nil {
		gql__type_GetDroidRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Starwars_Type_GetDroidRequest",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: `id of the droid`,
				},
			},
		})
	}
	return gql__type_GetDroidRequest
}

func Gql__type_Character() *graphql.Object {
	if gql__type_Character == nil {
		gql__type_Character = graphql.NewObject(graphql.ObjectConfig{
			Name: "Starwars_Type_Character",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"friends": &graphql.Field{
					Type: graphql.NewList(Gql__interface_Character()),
				},
				"appears_in": &graphql.Field{
					Type: graphql.NewList(Gql__enum_Episode()),
				},
				"home_planet": &graphql.Field{
					Type: graphql.String,
				},
				"primary_function": &graphql.Field{
					Type: graphql.String,
				},
				"type": &graphql.Field{
					Type: Gql__enum_Type(),
				},
			},
			Interfaces: []*graphql.Interface{
				Gql__interface_Character(),
			},
		})
	}
	return gql__type_Character
}

// graphql__resolver_StartwarsService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_StartwarsService struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_StartwarsService) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
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
func (x *graphql__resolver_StartwarsService) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"hero": &graphql.Field{
			Type: Gql__type_Character(),
			Args: graphql.FieldConfigArgument{
				"episode": &graphql.ArgumentConfig{
					Type:        Gql__enum_Episode(),
					Description: `If omitted, returns the hero of the whope saga. If provided, returns the hero of that particular episode.`,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req *GetHeroRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, err
				}
				client := NewStartwarsServiceClient(conn)
				resp, err := client.GetHero(p.Context, req)
				if err != nil {
					return nil, err
				}
				return resp, nil
			},
		},
		"human": &graphql.Field{
			Type: Gql__type_Character(),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: `id of the human`,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req *GetHumanRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, err
				}
				client := NewStartwarsServiceClient(conn)
				resp, err := client.GetHuman(p.Context, req)
				if err != nil {
					return nil, err
				}
				return resp, nil
			},
		},
		"droid": &graphql.Field{
			Type: Gql__type_Character(),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: `id of the droid`,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req *GetDroidRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, err
				}
				client := NewStartwarsServiceClient(conn)
				resp, err := client.GetDroid(p.Context, req)
				if err != nil {
					return nil, err
				}
				return resp, nil
			},
		},
		"humans": &graphql.Field{
			Type: graphql.NewList(Gql__type_Character()),
			Args: graphql.FieldConfigArgument{},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req *ListEmptyRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, err
				}
				client := NewStartwarsServiceClient(conn)
				resp, err := client.ListHumans(p.Context, req)
				if err != nil {
					return nil, err
				}
				return resp.GetHumans(), nil
			},
		},
		"droids": &graphql.Field{
			Type: graphql.NewList(Gql__type_Character()),
			Args: graphql.FieldConfigArgument{},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req *ListEmptyRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, err
				}
				client := NewStartwarsServiceClient(conn)
				resp, err := client.ListDroids(p.Context, req)
				if err != nil {
					return nil, err
				}
				return resp.GetDroids(), nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_StartwarsService) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterStartwarsServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterStartwarsServiceGraphql(mux *runtime.ServeMux) error {
	return RegisterStartwarsServiceGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
// service StartwarsService {
//    option (graphql.service) = {
//        host: "host:port"
//        insecure: true or false
//    };
//
//    ...with RPC definitions
// }
func RegisterStartwarsServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(&graphql__resolver_StartwarsService{
		conn: conn,
		host: "grpc:50051",
		dialOptions: []grpc.DialOption{
			grpc.WithInsecure(),
		},
	})
}
