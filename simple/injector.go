//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(
		fooSet,
		barSet,
		NewFooBarService,
	)
	return nil
}

// salah
// func InitializedHelloService() *HelloService {
// 	wire.Build(
// 		NewSayHelloImpl,
// 		NewHelloService,
// 	)
// 	return nil
// }

// bener
var HelloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(
		HelloSet,
		NewHelloService,
	)
	return nil
}

var fooBarSet = wire.NewSet(
	NewFoo,
	NewBar,
)

func InitializedFooBar() *FooBar {
	wire.Build(
		fooBarSet,
		wire.Struct(new(FooBar), "Foo", "Bar"),
	)
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(
		wire.InterfaceValue(new(io.Reader), os.Stdin),
	)
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(
		NewConnection, NewFile,
	)
	return nil, nil
}
