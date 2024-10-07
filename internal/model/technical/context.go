package technical

type ContextKey string

const (
	IsAuthenticated ContextKey = "IsAuthenticated"
	MerchantID      ContextKey = "MerchantID"
	Merchant        ContextKey = "Merchant"
	UserInfo        ContextKey = "UserInfo"
)
