// Package util provides utility functions.
package util

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// SourceInfo returns the filename and line number of the caller.
func SourceInfo() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok { // not easily testable
		return "<unknown>"
	}
	return fmt.Sprintf("%v:%v", filepath.Base(file), line)
}

// Wrap returns a slice of the given values.
func Wrap(x ...any) []any {
	return x
}
