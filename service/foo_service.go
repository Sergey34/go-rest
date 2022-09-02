package service

import "go-rest/model"

type FooService struct {
}

func NewFooService() FooService {
	return FooService{}
}

func (srv FooService) Foo(body []byte) (model.Foo, error) {
	var foo = model.Foo{Foo: "foo", FooFoo: "fooFoo"}
	return foo, nil
}
