package responses

import "github.com/spaghetti-/gonduit/entities"

// RepositoryQueryResponse is the result of repository.query operations.
type RepositoryQueryResponse []*entities.Repository
