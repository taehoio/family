package account_email

type AccountEmail struct {
	Email     string `dynamo:"email,hash"`
	AccountID string `dynamo:"account_id"`
}
