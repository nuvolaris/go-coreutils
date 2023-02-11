// Copyright (c) 2014-2016 Eric Lagergren
// Use of this source code is governed by the GPL v3 or later.

package cat

import (
	"bufio"
	"os"
	"syscall"

	k32 "github.com/EricLagergren/go-gnulib/windows"
	flag "github.com/ogier/pflag"
)

func main() {
	var ok int // return status

	outHandle := syscall.Handle(os.Stdout.Fd())
	outType, err := syscall.GetFileType(outHandle)
	if err != nil {
		fatal.Fatalln(err)
	}
	outBsize := 4096

	// catch (./cat) < /etc/group
	var args []string
	if flag.NArg() == 0 {
		args = []string{"-"}
	} else {
		args = flag.Args()
	}

	// the main loop
	var file *os.File
	for _, arg := range args {
		var inStat os.FileInfo

		if arg == "-" {
			file = os.Stdin
		} else {
			file, err = os.Open(arg)
			if err != nil {
				fatal.Fatalln(err)
			}

			inStat, err = file.Stat()
			if err != nil {
				fatal.Fatalln(err)
			}
			if inStat.IsDir() {
				fatal.Printf("%s: Is a directory\n", file.Name())
			}
		}

		inHandle := syscall.Handle(file.Fd())
		inBsize := 4096

		// See http://stackoverflow.com/q/29360969/2967113
		// for why this differs from the Unix versions.
		//
		// Make sure we're not catting a file to itself,
		// provided it's a regular file. Catting a non-reg
		// file to itself is cool, e.g. cat file > file
		if outType == syscall.FILE_TYPE_DISK {

			inPath := make([]byte, syscall.MAX_PATH)
			outPath := make([]byte, syscall.MAX_PATH)

			err = k32.GetFinalPathNameByHandleA(inHandle, inPath, 0)
			if err != nil {
				fatal.Fatalln(err)
			}

			err = k32.GetFinalPathNameByHandleA(outHandle, outPath, 0)
			if err != nil {
				fatal.Fatalln(err)
			}

			if string(inPath) == string(outPath) {
				if n, _ := file.Seek(0, os.SEEK_CUR); n < inStat.Size() {
					fatal.Fatalf("%s: input file is output file\n", file.Name())
				}
			}
		}

		if simple {
			outBuf := bufio.NewWriterSize(os.Stdout, 4096)
			ok ^= simpleCat(file, outBuf)

			// Flush because we don't have a chance to in
			// simpleCat() because we use io.Copy()
			outBuf.Flush()
		} else {
			// If you want to know why, exactly, I chose
			// outBsize -1 + inBsize*4 + 20, read GNU's cat
			// source code. The tl;dr is the 20 is the counter
			// buffer, inBsize*4 is from potentially prepending
			// the control characters (M-^), and outBsize is
			// due to new tests for newlines.
			size := outBsize - 1 + inBsize*4 + 20
			outBuf := bufio.NewWriterSize(os.Stdout, size)
			inBuf := make([]byte, inBsize+1)
			ok ^= cat(file, inBuf, outBuf)
		}

		file.Close()
	}

	os.Exit(ok)
}
