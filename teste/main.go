package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/Rhymond/go-money"
)

func main() {
	var myItems []Item = []Item{
		{
			ID: "123",
			Name: "Product1",
			Price: money.New(1000, "USD"),
		},
		{
			ID: "124",
			Name: "Product2",
			Price: money.New(1200, "USD"),
		},
	}
	var myCart Cart = Cart{
		ID: "1",
		Items: myItems,
		IsLocked: false,
	}

	price, err := myCart.TotalPrice()
	if err!=nil{
		log.Fatalf("%s",err)
	}
	fmt.Println(price.Display())
	err= myCart.Lock()
	if err!=nil{
		log.Fatalf("%s", err)
	}
	myProd:=Product{ID: "123", Description: "teste"}
	err=myProd.SetDescription("teste_pra_funcionar")
	if err!=nil{
		log.Fatalf("%+v",err)
	}
	fmt.Println(myProd)
}

type Item struct {
	ID    string
	Name  string
	Price *money.Money
}

type Cart struct {
	ID string
	IsLocked bool
	Items []Item
}

func (c Cart) TotalPrice() (*money.Money, error){
	var totalPrice *money.Money = money.New(0, "USD")
	var err error 
	for _, item := range c.Items{
		totalPrice, err = totalPrice.Add(item.Price)
		if err !=nil{
			return nil, fmt.Errorf("%s", err)
		}
	}
	return totalPrice, nil
}

func (c *Cart) Lock() error{
	if c.IsLocked {
		return errors.New("already locked") 
	}
	c.IsLocked = true
	return nil
}

type Product struct{
	ID string
	Description string
	Stock uint
}

func (p Product) IsInStock() bool{
	return p.Stock > 0
}

func (p *Product) SetDescription(description string) error{
	if len(description) <=10 || len(description) >= 250{
		return fmt.Errorf("product: %s has wrong size: %d",p.ID,len(description))
	}
	p.Description = description
	return nil
} 



