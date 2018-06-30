github.com/elliotchance/gedcom
==============================

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
  -string-name
    	Convert NAME tags to a string (instead of the object parts).
  -tag-keys
    	Use tags (pretty or raw) as object keys rather than arrays.
```
