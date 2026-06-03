package boilerplate

import "github.com/kkrypt0nn/tangra/v2"

func Example42() int {
	l := tangra.NewLogger()
	l.Info("Returning 42 soon")
	return 42
}
