package responses

import "github.com/spaghetti-/gonduit/entities"

// DifferentialQueryResponse is the response of calling differential.query.
type DifferentialQueryResponse []*entities.DifferentialRevision
