package gonduit

import (
	"github.com/spaghetti-/gonduit/requests"
	"github.com/spaghetti-/gonduit/responses"
)

// DifferentialQuery performs a call to differential.query.
func (c *Conn) DifferentialQuery(
	req requests.DifferentialQueryRequest,
) (*responses.DifferentialQueryResponse, error) {
	var res responses.DifferentialQueryResponse

	if err := c.Call("differential.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
