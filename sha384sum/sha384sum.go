/*
    go sha384sum

    Copyright (c) 2014-2015 Dingjun Fang

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License version 3 as
	published by the Free Software Foundation.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

/*

Sha384sum util implement by go.

Usage: sha384sum [OPTION]... [FILE]...

Print or check SHA384 (384-bit) checksums.

With no FILE, or when FILE is -, read standard input.

  -b, --binary  read in binary mode(default)
  -c, --check   read SHA384 sums from the FILEs and check them
  -t, --text    read in text mode
  Note: there is no difference between text and binary mode option.

The following three options are useful only when verifying checksums:
      --quiet    don't print OK for each successfully verified file
      --status   don't output anything, status code shows success
  -w, --warn     warn about improperly formated checksum lines

      --help     show help and exit
      --version  show version and exit

The sums are computed as described in FIPS-180-2.  When checking, the input
should be a former output of this program.  The default mode is to print
a line with checksum, a character indicating type ('*' for binary, ' ' for
text), and name for each FILE.
*/
package sha384sum

import (
	"fmt"
	cc "github.com/fangdingjun/go-coreutils/md5sum/checksum_common"
	flag "github.com/ogier/pflag"
	"os"
)

const (
	Help = `Usage: sha384sum [OPTION]... [FILE]...
Print or check SHA384 (384-bit) checksums.
With no FILE, or when FILE is -, read standard input.

  -b, --binary  read in binary mode(default)
  -c, --check   read SHA384 sums from the FILEs and check them
  -t, --text    read in text mode
  Note: there is no difference between text and binary mode option.

The following three options are useful only when verifying checksums:
      --quiet    don't print OK for each successfully verified file
      --status   don't output anything, status code shows success
  -w, --warn     warn about improperly formated checksum lines

      --help     show help and exit
      --version  show version and exit

The sums are computed as described in FIPS-180-2.  When checking, the input
should be a former output of this program.  The default mode is to print
a line with checksum, a character indicating type ('*' for binary, ' ' for
text), and name for each FILE.
`
	Version = `sha384sum (Go coreutils) 0.1
Copyright (C) 2015 Dingjun Fang
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
`
)

var (
	check_sum    = flag.BoolP("check", "c", false, "")
	no_output    = flag.BoolP("quiet", "q", false, "")
	no_output_s  = flag.BoolP("status", "", false, "")
	show_warn    = flag.BoolP("warn", "w", true, "")
	show_version = flag.BoolP("version", "v", false, "")
	text_mode    = flag.BoolP("text", "t", false, "")
	binary_mode  = flag.BoolP("binary", "b", false, "")
)

func Main() int {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s", Help)
		return 1
	}

	flag.Parse()

	/* trust --status and --quiet as the same */
	if *no_output_s == true {
		*no_output = true
	}

	has_error := false

	file_lists := flag.Args()
	if len(file_lists) == 0 {
		file_lists = append(file_lists, "-")
	}

	switch {
	case *show_version:
		fmt.Fprintf(os.Stdout, "%s", Version)
		return 0
	case *check_sum:
		if r := cc.CompareChecksum(file_lists, "sha384",
			!(*no_output), *show_warn); !r {
			has_error = true
		}
	default:
		if r := cc.GenerateChecksum(file_lists, "sha384"); !r {
			has_error = true
		}
	}

	if has_error {
		return 1
	}

	return 0
}
