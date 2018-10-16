# goversionflag

goversionflag is an go package aiming to help identifying versions of deployed go binarys.

## Installation
```
go get -u github.com/stockholmuniversity/goversionflag
```

## Usage
```go
package main

import (
        "fmt"
        "github.com/stockholmuniversity/goversionflag"
)

func main() {
        goversionflag.PrintVersionAndExit()
        fmt.Println("End of program, will not be shown when run with '-version'")
}
```

For more information see documentation and comments in printVersion.go

