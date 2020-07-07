package main

import (
	"archive/tar"
	"archive/zip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"flag"
	"fmt"
	"github.com/mholt/archiver/v3"
	"github.com/nwaples/rardecode"
	"hash"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	VERSION   = `v0.0.0`
	BUILD     = `dev`
	BUILDDATE = `0000-00-00T00:00:00+00:00`
)

const (
	AUTHOR             = `Pekka Järvinen`
	YEAR               = 2020
	HOMEPAGE           = `https://github.com/raspi/tarkistaja`
	CHECKSUM_SEPARATOR = " "
)

func main() {

	hasherArg := flag.String(`t`, `sha256`, `Checksum type (sha1, sha256, sha512, md5)`)
	outputFileArg := flag.String(`o`, ``, `Output checksums to file <filename> instead of STDOUT`)
	prefixWithArchiveArg := flag.Bool(`a`, false, `Add archive's file name as a directory name (as additional information)`)

	flag.Usage = func() {
		f := filepath.Base(os.Args[0])

		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "tarkistaja v%v - (%v)\n", VERSION, BUILDDATE)
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "(c) %v %v- [ %v ]\n", AUTHOR, YEAR, HOMEPAGE)
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "List file checksums inside of compressed archive.\n")
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "\n")

		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "  Usage:\n")
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "    %v [parameters] <compressed file>\n", f)

		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "\nParameters:\n")
		flag.PrintDefaults()
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "\n")

		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Examples:\n")
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "  List checksums:\n")
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "    %v important_files.zip\n", f)
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "  List checksums to file:\n")
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "    %v -o checksums.sha256 important_files.zip\n", f)
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "  List checksums using md5:\n")
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "    %v -t md5 important_files.zip\n", f)
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "\n")
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	filename := flag.Arg(0)
	filenameBase := filepath.Base(filename)

	f, err := archiver.ByExtension(filename)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, `%v`, err)
		os.Exit(1)
	}

	w, ok := f.(archiver.Walker)
	if !ok {
		_, _ = fmt.Fprint(os.Stderr, `invalid file format`)
		os.Exit(1)
	}

	// Select hashing method
	var hasher hash.Hash

	switch *hasherArg {
	case `md5`:
		hasher = md5.New()
	case `sha1`:
		hasher = sha1.New()
	case `sha256`:
		hasher = sha256.New()
	case `sha512`:
		hasher = sha512.New()
	default:
		_, _ = fmt.Fprintf(os.Stderr, `invalid checksum format %#v`, *hasherArg)
		os.Exit(1)
	}

	var outputter io.Writer

	outputIsStdOut := true
	if *outputFileArg == `` {
		outputter = os.Stdout
	} else {
		outputIsStdOut = false
		tmpf, err := os.Create(*outputFileArg)
		outputter = tmpf
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, `%v`, err)
			os.Exit(1)
		}
	}

	err = w.Walk(filename, func(f archiver.File) error {

		if f.IsDir() {
			// Skip directories
			return nil
		}

		hasher.Reset()

		name := f.Name()

		switch h := f.Header.(type) {
		case zip.FileHeader:
			name = h.Name
		case *tar.Header:
			name = h.Name
		case *rardecode.FileHeader:
			name = h.Name
		}

		if !outputIsStdOut {
			log.Printf(`hashing %#v`, name)
		}

		buffer := make([]byte, 1024)

		for {
			rb, err := f.Read(buffer)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}

				return err
			}

			_, err = hasher.Write(buffer[0:rb])
			if err != nil {
				return err
			}
		}

		_, _ = fmt.Fprintf(outputter, "%x%s", hasher.Sum(nil), CHECKSUM_SEPARATOR)

		if *prefixWithArchiveArg {
			// Add filename as it were a directory
			_, _ = fmt.Fprintf(outputter, "%s%c", filenameBase, os.PathSeparator)
		}

		_, _ = fmt.Fprintf(outputter, "%s\n", name)

		return nil
	})

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, `%v`, err)
		os.Exit(2)
	}

	if !outputIsStdOut {
		log.Print(`Done.`)
	}
}
