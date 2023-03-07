// Code generated by sdkgen. DO NOT EDIT.

package lockbox

import (
	"context"

	"google.golang.org/grpc"
)

// LockboxSecret provides access to "lockbox" component of Yandex.Cloud
type LockboxSecret struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewLockboxSecret creates instance of LockboxSecret
func NewLockboxSecret(g func(ctx context.Context) (*grpc.ClientConn, error)) *LockboxSecret {
	return &LockboxSecret{g}
}

// Secret gets SecretService client
func (l *LockboxSecret) Secret() *SecretServiceClient {
	return &SecretServiceClient{getConn: l.getConn}
}