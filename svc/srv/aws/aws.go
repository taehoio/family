package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	DefaultAwsRegion = "us-west-2"
)

type IFace interface {
	Session() *session.Session
}

type Aws struct {
	IFace

	session *session.Session
}

func New(cfgs ...*aws.Config) (IFace, error) {
	if len(cfgs) == 0 {
		cfgs = append(cfgs, &aws.Config{
			Region: aws.String(DefaultAwsRegion),
		})
	}

	sess, err := session.NewSession(cfgs...)
	if err != nil {
		return nil, err
	}

	return &Aws{
		session: sess,
	}, nil
}

func (a *Aws) Session() *session.Session {
	return a.session
}
