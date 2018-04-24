package responses

import "github.com/spaghetti-/gonduit/entities"

// PHIDQueryResponse is the result of phid.query operations.
type PHIDQueryResponse map[string]*entities.PHIDResult
