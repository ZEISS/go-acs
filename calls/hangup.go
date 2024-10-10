package calls

import (
	"context"
	"fmt"

	"github.com/zeiss/go-acs/client"
)

// CallHangUp is the method for recognizing call.
func (s *Service) CallHangUp(ctx context.Context, id string, key string, opts ...client.Opt) error {
	return s.client.Delete(ctx, key, fmt.Sprintf("/calling/callConnections/%s", id), "api-version=2024-06-15-preview", nil, opts...)
}
