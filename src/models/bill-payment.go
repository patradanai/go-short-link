package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BillPayment struct {
	Id                  primitive.ObjectID `bson:"_id,omitempty"`
	PaymentId           primitive.ObjectID `bson:"payment_id"`
	CorporateId         primitive.ObjectID `bson:"corporate_id"`
	BillPaymentItems    []BillPaymentItem  `bson:"bill_payment_items"`
	BillAmount          float32            `bson:"bill_amount,omitempty"`
	BillFee             float32            `bson:"bill_fee,omitempty"`
	BillTotal           float32            `bson:"bill_total,omitempty"`
	BillTotalAmount     float32            `bson:"bill_total_amount,omitempty"`
	BillStatus          string             `bson:"bill_status,omitempty"`
	BillStatusReason    string             `bson:"bill_status_reason,omitempty"`
	BillType            string             `bson:"bill_type,omitempty"`
	PaymentStatus       string             `bson:"payment_status,omitempty"`
	PaymentStatusReason string             `bson:"payment_status_reason,omitempty"`
	PaymentDateTime     time.Time          `bson:"payment_date_time,omitempty"`
	DueDate             time.Time          `bson:"expired_at,omitempty"`
	CreatedAt           time.Time          `bson:"created_at,omitempty"`
	UpdatedAt           time.Time          `bson:"updated_at,omitempty"`
}
