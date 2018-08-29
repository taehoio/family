package base

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	DefaultRegion = os.Getenv("AWS_DEFAULT_REGION")
)

type Aws interface {
	Session() *session.Session
}

type DefaultAws struct {
	Aws

	session *session.Session
}

func NewAws(cfgs ...*aws.Config) (Aws, error) {
	if len(cfgs) == 0 {
		cfgs = append(cfgs, &aws.Config{
			Region: aws.String(DefaultRegion),
		})
	}

	sess, err := session.NewSession(cfgs...)
	if err != nil {
		return nil, err
	}

	return &DefaultAws{
		session: sess,
	}, nil
}

func (a *DefaultAws) Session() *session.Session {
	return a.session
}
