package main

import "fmt"

type Reservation interface {
	GetName() string
}
type Invoice interface {}

type AbstractFacoty interface {
	CreateReservation() Reservation
	CreateInvoice() Invoice
}


type HoteFactory struct {

}

func (h HoteFactory) CreateReservation() Reservation {
	return new(HotelReservation)
}

func (h HoteFactory) CreateInvoice() Invoice {
	return new(HotelInvoice)
}

type FlightFactory struct {

}

func (f FlightFactory) CreateReservation() Reservation {
	return new(FlightReservation)
}

func (f FlightFactory) CreateInvoice() Invoice {
	return new(FlightInvoice)
}

type HotelReservation struct {}

func (h HotelReservation) GetName() string {
	return "My Hotel"
}

type HotelInvoice struct {}


type FlightReservation struct {}

func (f FlightReservation) GetName() string {
	panic("implement me")
}

type FlightInvoice struct {}

func GetFactory(vertical string) AbstractFacoty{
	var factory AbstractFacoty
	switch vertical {
	case "flight":
		factory=FlightFactory{}
	case "hotel":
		factory=HoteFactory{}

	}
	return factory
}


func main() {

	hotelFactory:=GetFactory("hotel")
	reservation :=hotelFactory.CreateReservation()
	fmt.Println(reservation.GetName())



}
