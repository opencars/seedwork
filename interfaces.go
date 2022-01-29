package seedwork

import "context"

type SessionChecker interface {
	CheckSession(ctx context.Context, sessionToken, cookie string) (*User, error)
}
