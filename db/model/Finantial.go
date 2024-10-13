package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Amount      uint
	PaymentType PaymentType
	Status      Status
	Type        TransactionType
	WalletID    *uint
	Wallet      *Wallet
	BasketId    *uint
}

type Wallet struct {
	gorm.Model
	Credit       float64
	Transactions []Transaction
}

type TransactionType int

const (
	PURCHASE TransactionType = iota
	TRANSFER
	DEPOSIT
)

type PaymentType int

const (
	WALLET PaymentType = iota
	CASH
	ONLINE
)

type Status int

const (
	SUCCESS Status = iota
	FAILED
	PENDING
)
