```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["key"] = 10 // Now this works
    fmt.Println(m["key"])
}
```