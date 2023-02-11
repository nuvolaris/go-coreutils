package true

import (
	"fmt"
	"os"
)

const (
	Help = `Usage: /usr/bin/true [ignored command line arguments]
  or:  /usr/bin/true OPTION
Exit with a status code indicating success.

      --help     display this help and exit
      --version  output version information and exit

NOTE: your shell may have its own version of true, which usually supersedes
the version described here.  Please refer to your shell's documentation
for details about the options it supports.
`
	Version = `true (Go coreutils) 1.0
Copyright (C) 2015 Eric Lagergren
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
`
)

func Main() int {
	if len(os.Args) == 2 {
		if os.Args[1] == "--help" {
			fmt.Printf("%s", Help)
		}

		if os.Args[1] == "--version" {
			fmt.Printf("%s", Version)
		}
	}
}
