package responses

import "github.com/spaghetti-/gonduit/entities"

// ManiphestGetTaskTransactionsResponse is the response of calling maniphest.query.
type ManiphestGetTaskTransactionsResponse map[string][]*entities.ManiphestTaskTranscation
