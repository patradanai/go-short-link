package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentTransaction struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	PaymentType     PaymentType        `bson:"payment_type"`
	CorporateId     primitive.ObjectID `bson:"corporate_id"`
	BillPaymentId   primitive.ObjectID `bson:"bill_payment_id"`
	Ref1            string             `bson:"ref1,omitempty"`
	Ref2            string             `bson:"ref2,omitempty"`
	Ref3            string             `bson:"ref2,omitempty"`
	TransactionId   string             `bson:"transaction_id,omitempty"`
	TransactionDate time.Time          `bson:"transaction_date,omitempty"`
	RequestId       string             `bson:"request_id,omitempty"`
	RequestDate     time.Time          `bson:"request_date,omitempty"`
	Amount          string             `bson:"amount,omitempty"`
	Currency        string             `bson:"currency,omitempty"`
	Status          string             `bson:"status,omitempty"`
	StatusReason    string             `bson:"status_reason,omitempty"`
	FromName        string             `bson:"from_name,omitempty"`
	FromBank        string             `bson:"from_bank,omitempty"`
	ApproveCode     string             `bson:"approve_code,omitempty"`
	ConfirmId       string             `bson:"confirm_id,omitempty"`
	RequestLog      string             `bson:"request_log,omitempty"`
	FromIp          string             `bson:"from_ip,omitempty"`
	DueDate         time.Time          `bson:"due_date,omitempty"`
	CreatedAt       time.Time          `bson:"created_at,omitempty"`
	UpdatedAt       time.Time          `bson:"updated_at,omitempty"`
}
