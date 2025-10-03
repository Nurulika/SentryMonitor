// cmd/sentrymonitor/main.go
package main

import (
    "flag"
    "log"
    "os"

    "sentrymonitor/internal/sentrymonitor"
)

func main() {
    verbose := flag.Bool("verbose", false, "Enable verbose logging")
    input := flag.String("input", "", "Input file path")
    output := flag.String("output", "", "Output file path")
    flag.Parse()

    app := sentrymonitor.NewApp(*verbose)
    if err := app.Run(*input, *output); err != nil {
        log.Printf("Error: %v", err)
        os.Exit(1)
    }
}
