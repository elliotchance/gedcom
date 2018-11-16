github.com/elliotchance/gedcom
==============================

[![Build Status](https://travis-ci.org/elliotchance/gedcom.svg?branch=master)](https://travis-ci.org/elliotchance/gedcom)
[![codecov](https://codecov.io/gh/elliotchance/gedcom/branch/master/graph/badge.svg)](https://codecov.io/gh/elliotchance/gedcom)
![GitHub release](https://img.shields.io/github/release/elliotchance/gedcom.svg)
[![Join the chat at https://gitter.im/gedcom-go/Lobby](https://badges.gitter.im/gedcom-go/Lobby.svg)](https://gitter.im/gedcom-go/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

`github.com/elliotchance/gedcom` is an advanced Go-style library and set of
command-line tools for dealing with
[GEDCOM files](https://en.wikipedia.org/wiki/GEDCOM).

You can download the latest binaries for macOS, Windows and Linux on the
[Releases page](https://github.com/elliotchance/gedcom/releases). This will not
require you to install Go or any other dependencies.

| Package              | Description |
| -------------------- | ----------- |
| [![GoDoc](https://godoc.org/github.com/elliotchance/gedcom?status.svg)](https://godoc.org/github.com/elliotchance/gedcom) <br/> `gedcom` | Package gedcom contains functionality for encoding, decoding, traversing, manipulating and comparing of GEDCOM data. |
| [![GoDoc](https://godoc.org/github.com/elliotchance/gedcom/gedcom2html?status.svg)](https://godoc.org/github.com/elliotchance/gedcom/gedcom2html) <br/> `gedcom/gedcom2html` | Gedcom2html renders a GEDCOM file into HTML pages that can be shared and published easily. |
| [![GoDoc](https://godoc.org/github.com/elliotchance/gedcom/gedcom2json?status.svg)](https://godoc.org/github.com/elliotchance/gedcom/gedcom2json) <br/> `gedcom/gedcom2json` | Gedcom2json converts a GEDCOM file to JSON so that it can easily processed and consumed by other applications. |
| [![GoDoc](https://godoc.org/github.com/elliotchance/gedcom/gedcom2text?status.svg)](https://godoc.org/github.com/elliotchance/gedcom/gedcom2text) <br/>`gedcom/gedcom2text` | Gedcom2text converts a GEDCOM file to a human-readable text format. |
| [![GoDoc](https://godoc.org/github.com/elliotchance/gedcom/gedcomdiff?status.svg)](https://godoc.org/github.com/elliotchance/gedcom/gedcomdiff) <br/> `gedcom/gedcomdiff` | Gedcomdiff is a tool for comparing GEDCOM files and producing a HTML report. |
| [![GoDoc](https://godoc.org/github.com/elliotchance/gedcom/q?status.svg)](https://godoc.org/github.com/elliotchance/gedcom/q) <br/> `gedcom/q` | Package q is the gedcomq parser and engine. |
| [![GoDoc](https://godoc.org/github.com/elliotchance/gedcom/gedcomq?status.svg)](https://godoc.org/github.com/elliotchance/gedcom/gedcomq) <br/> `gedcom/gedcomq` | Gedcomq is a command line tool and query language for GEDCOM files heavily inspired by [jq](https://stedolan.github.io/jq/), in name and syntax. |
| [![GoDoc](https://godoc.org/github.com/elliotchance/gedcom/gedcomtune?status.svg)](https://godoc.org/github.com/elliotchance/gedcom/gedcomtune) <br/> `gedcom/gedcomtune` | Gedcomtune is used to calculate the ideal weights and similarities for the main `gedcom` package. |
| [![GoDoc](https://godoc.org/github.com/elliotchance/gedcom/html?status.svg)](https://godoc.org/github.com/elliotchance/gedcom/html) <br/> `gedcom/html` | Package html is shared HTML rendering components that are shared by the other packages. |
| [![GoDoc](https://godoc.org/github.com/elliotchance/gedcom/util?status.svg)](https://godoc.org/github.com/elliotchance/gedcom/util) <br/> `gedcom/util` | Package util contains shared functions used by several packages. |
