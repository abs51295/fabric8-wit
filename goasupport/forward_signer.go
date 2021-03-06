package goasupport

import (
	"context"
	"net/http"

	goaclient "github.com/goadesign/goa/client"
	goajwt "github.com/goadesign/goa/middleware/security/jwt"
)

// ForwardSigner reuse Token from caller and forward to target Request
type forwardSigner struct {
	token string
}

// Sign set the Auth header
func (f forwardSigner) Sign(request *http.Request) error {
	request.Header.Set("Authorization", "Bearer "+f.token)
	return nil
}

// NewForwardSigner return a new signer based on current context
func NewForwardSigner(ctx context.Context) goaclient.Signer {
	token := goajwt.ContextJWT(ctx)
	if token == nil {
		return nil
	}
	return &forwardSigner{token: token.Raw}
}
