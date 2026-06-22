package steel

import (
	"context"
)

type WellKnown struct {
	client *Client
}

// Get Public Keys (JWKS)
func (r *WellKnown) List(ctx context.Context, opts ...RequestOption) (*WellknownListResponse, error) {
	path := "/.well-known/jwks.json"
	var out WellknownListResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}
