package boilerplate

import "github.com/kkrypt0nn/tangra"

func Example42() int {
	l := tangra.NewLogger()
	l.Info("Returning 42 soon")
	return 42
}
