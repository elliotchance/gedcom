github.com/elliotchance/gedcom
==============================

[![Build Status](https://travis-ci.org/elliotchance/gedcom.svg?branch=master)](https://travis-ci.org/elliotchance/gedcom)
[![GoDoc](https://godoc.org/github.com/elliotchance/gedcom?status.svg)](https://godoc.org/github.com/elliotchance/gedcom)

`gedcom` provides a Go-style encoder and decoder for GEDCOM files.

The goals of this project are:

1. **Support all GEDCOM files by supporting the encoding and not the GEDCOM
standard itself.** Many GEDCOM libraries that try to follow the standard run
into trouble when applications do not follow the same standards or the standard
is interpreted differently. `gedcom` retains all tags and structure in the
original GEDCOM file.

2. **Build some structures/types that provide a nicer API for common
operations.** For example iterating through individuals in a file,
accessing/formatting names, standardising dates, etc.

3. **Provide a program to convert GEDCOM to JSON**. So that any GEDCOM file can
be ingested and processed more easily in other applications. The bundled
`gedcom2json` program does this (and offers several options).

Decoding GEDCOM
===============

```go
ged := "0 HEAD\n1 CHAR UTF-8"

decoder := gedcom.NewDecoder(strings.NewReader(ged))
document, err := decoder.Decode()
if err != nil {
    panic(err)
}
```

Encoding GEDCOM
===============

If you have already decoded a GEDCOM into a `Document` (in the previous example)
then you can simply encode it back to a GEDCOM string with:

```go
document.String()
```

This is really just a shorthand for using the proper encoder:

```go
buf := bytes.NewBufferString("")

encoder := NewEncoder(buf, doc)
err := encoder.Encode()
if err != nil {
	panic(err)
}
```

gedcom2html
===========

`gedcom2html` converts a GEDCOM file to a directory of HTML files.

```txt
Usage of gedcom2html:
  -gedcom string
    	Input GEDCOM file.
  -output-dir string
    	Output directory. It will use the current directory if output-dir is not provided. Output files will only be added or replaced. Existing files will not be deleted. (default ".")
```

gedcom2json
===========

`gedcom2json` is a subpackage and binary that converts a GEDCOM file to a JSON
structure. It offers several options for the output:

```
Usage of gedcom2json:
  -exclude-tags string
    	Comma-separated list of tags to ignore.
  -gedcom string
    	Input GEDCOM file.
  -no-pointers
    	Do not include Pointer values ("ptr" attribute) in the output JSON. This is useful to activate when comparing GEDCOM files that have had pointers generated from different sources.
  -only-official-tags
    	Only include tags from the GEDCOM standard in the output.
  -only-tags string
    	Only include these tags in the output.
  -pretty-json
    	Pretty print JSON.
  -pretty-tags
    	Output tags with their descriptive name instead of their raw tag value. For example, "BIRT" would be output as "Birth".
  -single-name
    	When there are multiple names for an individual this will return the first of the name nodes only.
  -string-name
    	Convert NAME tags to a string (instead of the object parts).
  -tag-keys
    	Use tags (pretty or raw) as object keys rather than arrays.
```

gedcom2text
===========

`gedcom2text` is a subpackage and binary that converts a GEDCOM file to a simple
text output (or split into individual files) that is ideal for easily reading
(by a person) and designed to be as friendly as possible when using diff tools.

```
Usage of gedcom2text:
  -gedcom string
    	Input GEDCOM file.
  -no-change-times
    	Do not change timestamps.
  -no-empty-deaths
    	Do not include Death node if there are no visible details.
  -no-places
    	Do not include places.
  -no-sources
    	Do not include sources.
  -only-official-tags
    	Only output official GEDCOM tags.
  -single-name
    	Only output the primary name.
  -split-dir string
    	Split the individuals into separate files in this directory.
```

Comparing GEDCOM Files
----------------------

Here is an example to compare two large GEDCOM files:

```bash
gedcom2text -gedcom file1.ged -no-sources -only-official-tags -split-dir out1
gedcom2text -gedcom file2.ged -no-sources -only-official-tags -split-dir out2
diff -bur out1/ out2/
```

You can (and probably should) also use
[more pretty diffing tools](https://en.wikipedia.org/wiki/Comparison_of_file_comparison_tools).
