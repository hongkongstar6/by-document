package config

type PathKey string

const (
	CreatePrepayOrder PathKey = "CREATE_PREPAY_ORDER"
	QueryOrder        PathKey = "QUERY_ORDER"
	Refund            PathKey = "REFUND"
	QueryRefund       PathKey = "QUERY_REFUND"
	QueryBalance      PathKey = "QUERY_BALANCE"
	Phone             PathKey = "PHONE"
	QueryPayment      PathKey = "QUERY_PAYMENT"
)

var PathEnum = map[PathKey]string{
	CreatePrepayOrder: "sdk/open/transactions/createPrepaymentOrder",
	QueryOrder:        "sdk/open/transactions/pay/info",
	Refund:            "sdk/open/transactions/refund",
	QueryRefund:       "sdk/open/transactions/refund/info",
	QueryBalance:      "sdk/open/payment/balance",
	Phone:             "sdk/open/payment/phone",
	QueryPayment:      "sdk/open/payment/orders",
}
