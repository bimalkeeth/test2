package main

import "fmt"

type Seller interface {
	CancelReservation(charge float64)
}
type Reservationx struct {
	sellerRef Seller
}

func (r *Reservationx) Cancel() {
	r.sellerRef.CancelReservation(10)
}

type PremiumReservation struct {
	Reservationx
}

func (r PremiumReservation) Cancel() {
	r.sellerRef.CancelReservation(0)
}

type InstitutionalSeller struct{}

func (i InstitutionalSeller) CancelReservation(charge float64) {
	fmt.Println("Institutional seller cancel reservation", charge)
}

type SmallScaleSeller struct{}

func (s SmallScaleSeller) CancelReservation(charge float64) {

	fmt.Println("Small scale seller cancel reservation", charge)
}

func main() {
	res := Reservationx{InstitutionalSeller{}}
	res.Cancel()

	premiumRes := PremiumReservation{Reservationx{SmallScaleSeller{}}}
	premiumRes.Cancel()
}
