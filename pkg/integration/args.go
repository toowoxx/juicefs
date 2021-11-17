package integration

import "os"

// Args is used as a replacement for os.Args so that we can
// directly embed the entire app in a plugin
var Args []string

func SetArgs(args []string) {
	Args = append(os.Args[:1], args...)
}
