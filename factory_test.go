package main

import (
	"fmt"
	"testing"
)

type Product interface {
	use()
}

type Factory interface {
	createProduct(owner string) Product
	registerProduct()
}
type IdCard struct {
}

func (c *IdCard) use() {
	fmt.Println("Thr Card use now")
}

type IDCardFactory struct {
	Owners []string
}

func (f *IDCardFactory) createProduct() Product {
	return &IdCard{}
}
func (f *IDCardFactory) registerProduct() {
	fmt.Println("The IdCard register success")
}
func TestFactoryMethod(t *testing.T) {
	owners := make([]string, 0)
	idcardFactory := IDCardFactory{owners}
	idCard := idcardFactory.createProduct()
	idCard.use()
	idcardFactory.registerProduct()
}
