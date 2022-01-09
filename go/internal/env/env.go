package env

import "os"

var (
	Var = os.Getenv("VAR")
)
