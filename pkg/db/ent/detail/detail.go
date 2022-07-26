// Code generated by entc, DO NOT EDIT.

package detail

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the detail type in the database.
	Label = "detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldPaymentID holds the string denoting the payment_id field in the database.
	FieldPaymentID = "payment_id"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldPaymentCoinTypeID holds the string denoting the payment_coin_type_id field in the database.
	FieldPaymentCoinTypeID = "payment_coin_type_id"
	// FieldPaymentCoinUsdCurrency holds the string denoting the payment_coin_usd_currency field in the database.
	FieldPaymentCoinUsdCurrency = "payment_coin_usd_currency"
	// FieldUnits holds the string denoting the units field in the database.
	FieldUnits = "units"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldUsdAmount holds the string denoting the usd_amount field in the database.
	FieldUsdAmount = "usd_amount"
	// Table holds the table name of the detail in the database.
	Table = "details"
)

// Columns holds all SQL columns for detail fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldUserID,
	FieldGoodID,
	FieldOrderID,
	FieldPaymentID,
	FieldCoinTypeID,
	FieldPaymentCoinTypeID,
	FieldPaymentCoinUsdCurrency,
	FieldUnits,
	FieldAmount,
	FieldUsdAmount,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/archivement-manager/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultGoodID holds the default value on creation for the "good_id" field.
	DefaultGoodID func() uuid.UUID
	// DefaultOrderID holds the default value on creation for the "order_id" field.
	DefaultOrderID func() uuid.UUID
	// DefaultPaymentID holds the default value on creation for the "payment_id" field.
	DefaultPaymentID func() uuid.UUID
	// DefaultCoinTypeID holds the default value on creation for the "coin_type_id" field.
	DefaultCoinTypeID func() uuid.UUID
	// DefaultPaymentCoinTypeID holds the default value on creation for the "payment_coin_type_id" field.
	DefaultPaymentCoinTypeID func() uuid.UUID
	// DefaultUnits holds the default value on creation for the "units" field.
	DefaultUnits uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)