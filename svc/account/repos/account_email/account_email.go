package account_email

type AccountEmail struct {
	Email     string `dynamo:"email"`
	AccountID string `dynamo:"account_id"`
}
