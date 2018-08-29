package handler

import (
	"os"
	"testing"

	"github.com/icrowley/fake"
	"github.com/rs/xid"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/services/accounts/pkg/crypt"
)

var (
	testAccountID         = xid.New().String()
	testAuthTypeEmail     = accounts.AuthType_EMAIL
	testFullName          = fake.FullName()
	testEmail             = fake.EmailAddress()
	testPassword          = fake.SimplePassword()
	testHashedPassword, _ = crypt.NewMock().HashPassword(testPassword)
)

func TestMain(m *testing.M) {
	retCode := m.Run()
	os.Exit(retCode)
}
