// cmd/nftmarketplacegeneratorsdk/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplacegeneratorsdk/internal/nftmarketplacegeneratorsdk"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplacegeneratorsdk.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
