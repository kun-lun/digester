package main

import (
    "fmt"
    "log"

    "github.com/kun-lun/digester/pkg/apis"
)

// For manual test
func main() {
    if err := apis.Run("questionnaire.yml"); err != nil {
        log.Fatal(err)
    }
    if bp, err := apis.ImportBlueprintYaml("questionnaire.yml"); err != nil {
        log.Fatal(err)
    } else {
        fmt.Printf("%#v\n", bp)
    }
}
