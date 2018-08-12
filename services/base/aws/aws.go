package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	DefaultRegion = os.Getenv("AWS_DEFAULT_REGION")
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
			Region: aws.String(DefaultRegion),
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
