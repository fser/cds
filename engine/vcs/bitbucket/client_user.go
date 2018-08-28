package bitbucket

import (
	"context"
	"fmt"
	"net/url"

	"github.com/ovh/cds/sdk"
)

func (b *bitbucketClient) findByEmail(ctx context.Context, email string) (*User, error) {
	var users = UsersResponse{}
	var path = "/admin/users"
	if err := b.do(ctx, "GET", "core", path, url.Values{"filter": []string{email}}, nil, &users, nil); err != nil {
		return nil, sdk.WrapError(err, "Error during consumption")
	}
	if len(users.Values) >= 1 {
		return &users.Values[0], nil
	}
	return nil, fmt.Errorf("User not found")
}
