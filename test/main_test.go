package test

import (
	"github.com/766800551/mopler"
	"github.com/766800551/mopler/impl/authorization"
	"testing"
)

var sdk *mopler.SdkContext

func TestMain(m *testing.M) {
	q := authorization.NewQuickTokenImpl(mopler.Token)
	sdk = mopler.New(q)
	m.Run()
}
