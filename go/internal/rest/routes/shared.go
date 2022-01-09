package routes

import "github.com/cleanarc/go-core/rest"

const (
	V1 = "/api/v1"
)

// Setup for router configurations which are independent of endpoint domains
func Setup(r *rest.Router) {
	r.AddGroup(V1)
}
