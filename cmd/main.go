package main

import "github.com/damondouglas/fq/pkg/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
