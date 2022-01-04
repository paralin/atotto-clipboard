package clipboard

import "errors"

// TODO: bindings to js clipboard APIs
// https://github.com/gowebapi/webapi/blob/master/clipboard/clipboard.go

var missingCommands = errors.New("TODO clipboard support for js")

func init() {
	Unsupported = true
}

func readAll() (string, error) {
	// if Unsupported {
	return "", missingCommands
}

func writeAll(text string) error {
	// if Unsupported {
	return missingCommands
}
