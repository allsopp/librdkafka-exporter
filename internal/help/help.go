package help

import (
	_ "embed"
	"log"

	"github.com/stretchr/objx"
)

//go:embed help.json
var json string
var help objx.Map

func Get(name string) string {
	return help.Get(name).Str()
}

func init() {
	var err error
	help, err = objx.FromJSON(json)
	if err != nil {
		log.Fatalf("error loading help text: %v", err)
	}
}
