package gonduit

import (
	"github.com/spaghetti-/gonduit/requests"
	"github.com/spaghetti-/gonduit/responses"
)

// UserQuery performs a call to differential.query.
func (c *Conn) UserQuery(
	req requests.UserQueryRequest,
) (*responses.UserQueryResponse, error) {
	var res responses.UserQueryResponse

	if err := c.Call("user.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
