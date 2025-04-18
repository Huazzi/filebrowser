// Copyright 2015 Light Code Labs, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package runner

import (
	"errors"
	"runtime"
	"unicode"

	"github.com/flynn/go-shlex"
)

const (
	osWindows = "windows"
	osLinux   = "linux"
)

var runtimeGoos = runtime.GOOS

// SplitCommandAndArgs takes a command string and parses it shell-style into the
// command and its separate arguments.
func SplitCommandAndArgs(command string) (cmd string, args []string, err error) {
	var parts []string

	if runtimeGoos == osWindows {
		parts = parseWindowsCommand(command) // parse it Windows-style
	} else {
		parts, err = parseUnixCommand(command) // parse it Unix-style
		if err != nil {
			err = errors.New("error parsing command: " + err.Error())
			return
		}
	}

	if len(parts) == 0 {
		err = errors.New("no command contained in '" + command + "'")
		return
	}

	cmd = parts[0]
	if len(parts) > 1 {
		args = parts[1:]
	}

	return
}

// parseUnixCommand parses a unix style command line and returns the
// command and its arguments or an error
func parseUnixCommand(cmd string) ([]string, error) {
	return shlex.Split(cmd)
}

// parseWindowsCommand parses windows command lines and
// returns the command and the arguments as an array. It
// should be able to parse commonly used command lines.
// Only basic syntax is supported:
//   - spaces in double quotes are not token delimiters
//   - double quotes are escaped by either backspace or another double quote
//   - except for the above case backspaces are path separators (not special)
//
// Many sources point out that escaping quotes using backslash can be unsafe.
// Use two double quotes when possible. (Source: http://stackoverflow.com/a/31413730/2616179 )
//
// This function has to be used on Windows instead
// of the shlex package because this function treats backslash
// characters properly.
func parseWindowsCommand(cmd string) []string {
	const backslash = '\\'
	const quote = '"'

	var parts []string
	var part string
	var inQuotes bool
	var lastRune rune

	for i, ch := range cmd {
		if i != 0 {
			lastRune = rune(cmd[i-1])
		}

		if ch == backslash {
			// put it in the part - for now we don't know if it's an
			// escaping char or path separator
			part += string(ch)
			continue
		}

		if ch == quote {
			if lastRune == backslash {
				// remove the backslash from the part and add the escaped quote instead
				part = part[:len(part)-1]
				part += string(ch)
				continue
			}

			if lastRune == quote {
				// revert the last change of the inQuotes state
				// it was an escaping quote
				inQuotes = !inQuotes
				part += string(ch)
				continue
			}

			// normal escaping quotes
			inQuotes = !inQuotes
			continue
		}

		if unicode.IsSpace(ch) && !inQuotes && part != "" {
			parts = append(parts, part)
			part = ""
			continue
		}

		part += string(ch)
	}

	if part != "" {
		parts = append(parts, part)
	}

	return parts
}
