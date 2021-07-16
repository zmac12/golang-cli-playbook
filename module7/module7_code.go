package module7

import (
	"fmt"
	"runtime"
)

func content() {
	fmt.Println(runtime.GOOS)
}

//go:generate goimports module7_code.go
