package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/ggchangan/go-scaffold/internal/apiserver"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	apiserver.NewApp("apiserver").Run()
}
