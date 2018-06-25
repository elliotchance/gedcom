`github.com/elliotchance/gedcom` is a GEDCOM decoder that aims to support all
GEDCOM files regardless of the version or propritary extensions by parsing the
GEDCOM file into a generic nested structure.

Example
=======

```go
ged := "0 HEAD\n1 CHAR UTF-8"

decoder := gedcom.NewDecoder(strings.NewReader(ged))
document, err := decoder.Decode()
if err != nil {
    panic(err)
}
```
