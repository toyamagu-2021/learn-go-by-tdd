package main

import (
	"os"
	"time"

	"github.com/toyamagu-2021/learn-go-by-tdd/math/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
