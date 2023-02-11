package uptime

import (
	"fmt"
	"os"

	"github.com/EricLagergren/go-gnulib/utmp"

	flag "github.com/ogier/pflag"
)

func Main() int {
	flag.Usage = func() {
		// This is a little weird because I want to insert the correct
		// UTMP/WTMP file names into the Help output, but usually my
		// Help constants are raw string literals, so I had to
		// break it up into a couple chunks and move around some formatting.
		fmt.Fprintf(os.Stderr, "%s %s.  %s %s",
			Help1, utmp.UtmpFile, utmp.WtmpFile, Help2)
		return 1
	}
	flag.Parse()

	if *version {
		fmt.Printf("%s\n", Version)
		return 0
	}

	switch flag.NArg() {
	case 0:
		uptime(utmp.UtmpFile, utmp.CheckPIDs)
	case 1:
		uptime(flag.Arg(0), 0)
	default:
		fatal.Fatalf("extra operand %s\n", flag.Arg(1))
	}
}
