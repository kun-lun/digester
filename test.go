package main

import (
    "fmt"
    "log"

    "github.com/kun-lun/digester/pkg/questionnaire"
)

// For manual test
func main() {
    bp := questionnaire.Run()
    fmt.Printf("%#v\n", bp)
    if err := bp.ExposeYaml("questionnaire.yml"); err != nil {
        log.Fatal(err)
    }
    /*
    if err := bp.ImportYaml("questionnaire.yml"); err != nil {
        log.Fatal(err)
    } else {
        fmt.Printf("%#v\n", bp)
    }
    */
}
