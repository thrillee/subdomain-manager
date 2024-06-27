package internals

import "os"

type Properties struct {
	JSON_FILE_PATH string
}

func CreateProperties() *Properties {
	props := Properties{
		JSON_FILE_PATH: os.Getenv("JSON_FILE_PATH"),
	}

	return &props
}
