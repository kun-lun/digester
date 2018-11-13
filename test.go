package main

import (
	"fmt"
	"log"

	"github.com/kun-lun/common/storage"
	"github.com/kun-lun/digester/pkg/apis"
)

// For manual test
func main() {
	state := storage.State{
		EnvID: "digest-test",
	}
	if err := apis.Run(state, "questionnaire.yml"); err != nil {
		log.Fatal(err)
	}
	if bp, err := apis.ImportBlueprintYaml("questionnaire.yml"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%#v\n", bp)
	}
}
