package main

import (
	"errors"
	"fmt"
	"time"
)

type Depositor interface {
	TransferDeposit() (bool, error)
	CancelTransaction() (bool, error)
}

type PayPal struct {
	TransationCode string
	Deposit        float64
}

func (v *PayPal) TransferDeposit() (bool, error) {
	fmt.Printf("Transfering $%.2f to %s\n", v.Deposit, v.TransationCode)
	time.Sleep(2 * time.Second)
	fmt.Println("Transaction completed")
	return true, nil
}

func (v *PayPal) CancelTransaction() (bool, error) {
	return true, nil
}

type MercadoPago struct {
	DestinationAccount string
	Deposit            float64
}

func (v *MercadoPago) TransferDeposit() (bool, error) {
	fmt.Printf("Transfering $%.2f to %s\n", v.Deposit, v.DestinationAccount)
	time.Sleep(2 * time.Second)
	fmt.Println("Transaction completed")
	return true, nil
}

func (v *MercadoPago) CancelTransaction() (bool, error) {
	return true, nil
}

type Oxxo struct {
	AccountName   string
	AccountNumber string
	Deposit       float64
}

func (v *Oxxo) TransferDeposit() (bool, error) {
	fmt.Printf("Transfering $%.2f to %s\n", v.Deposit, v.AccountNumber)
	time.Sleep(2 * time.Second)
	fmt.Println("No hay sistema joven")
	return v.CancelTransaction()
}

func (v *Oxxo) CancelTransaction() (bool, error) {
	return false, errors.New("El joven no quizo pagar comisión de $10 pesitos")
}

// DoTransaction does the transaction given an interface Depositor
func DoTransaction(deposit Depositor) error {
	_, err := deposit.TransferDeposit()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	paypal := PayPal{
		TransationCode: "4242424242424242",
		Deposit:        15.66,
	}
	mercadoPago := MercadoPago{
		DestinationAccount: "4243456367472324",
		Deposit:            67.64,
	}
	oxxo := Oxxo{
		AccountName:   "Orlando Romo",
		AccountNumber: "3456434543454",
		Deposit:       34.12,
	}

	err := DoTransaction(&paypal)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = DoTransaction(&mercadoPago)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = DoTransaction(&oxxo)
	if err != nil {
		fmt.Println(err.Error())
	}
}
