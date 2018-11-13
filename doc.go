// Package gedcom contains functionality for encoding, decoding, traversing,
// manipulating and comparing of GEDCOM data.
//
// Installation
//
// You can download the latest binaries for macOS, Windows and Linux on the
// Releases page: https://github.com/elliotchance/gedcom/releases
//
// This will not require you to install Go or any other dependencies.
//
// If you wish to build it from source you must install the dependencies with:
//
//   dep ensure
//
// Decoding a Document
//
// Decoding a GEDCOM stream:
//
//   ged := "0 HEAD\n1 CHAR UTF-8"
//
//   decoder := gedcom.NewDecoder(strings.NewReader(ged))
//   document, err := decoder.Decode()
//   if err != nil {
//     panic(err)
//   }
//
// If you are reading from a file you can use NewDocumentFromGEDCOMFile:
//
//   document, err := gedcom.NewDocumentFromGEDCOMFile("family.ged")
//   if err != nil {
//       panic(err)
//   }
//
// Encoding a Document
//
//   buf := bytes.NewBufferString("")
//
//   encoder := NewEncoder(buf, doc)
//   err := encoder.Encode()
//   if err != nil {
//     panic(err)
//   }
//
// If you need the GEDCOM data as a string you can simply using fmt.Stringer:
//
//   data := document.String()
//
// Traversing a Document
//
// On top of the raw document is a powerful API that takes care of the complex
// traversing of the Document. Here is a simple example:
//
//   for _, individual := range document.Individuals() {
//     fmt.Println(individual.Name().String())
//   }
//
// Some of the nodes in a GEDCOM file have been replaced with more function rich
// types, such as names, dates, families and more.
//
// Comparing & Diffing Nodes
//
// CompareNodes recursively compares two nodes. For example:
//
//   0 INDI @P3@           |  0 INDI @P4@
//   1 NAME John /Smith/   |  1 NAME J. /Smith/
//   1 BIRT                |  1 BIRT
//   2 DATE 3 SEP 1943     |  2 DATE Abt. Sep 1943
//   1 DEAT                |  1 BIRT
//   2 PLAC England        |  2 DATE 3 SEP 1943
//   1 BIRT                |  1 DEAT
//   2 DATE Abt. Oct 1943  |  2 DATE Aft. 2001
//                         |  2 PLAC Surry, England
//
// Produces a *NodeDiff than can be rendered with the String method:
//
//   LR 0 INDI @P3@
//   L  1 NAME John /Smith/
//   LR 1 BIRT
//   L  2 DATE Abt. Oct 1943
//   LR 2 DATE 3 SEP 1943
//    R 2 DATE Abt. Sep 1943
//   LR 1 DEAT
//   L  2 PLAC England
//    R 2 DATE Aft. 2001
//    R 2 PLAC Surry, England
//    R 1 NAME J. /Smith/
//
// Dates
//
// Dates in GEDCOM files can be very complex as they can cater for many
// scenarios:
//
// 1. Incomplete, like "Dec 1943"
//
// 2. Anchored, like "Aft. 3 Sep 2003" or "Before 1923"
//
// 3. Ranges, like "Bet. 4 Apr 1823 and 8 Apr 1823"
//
// This package provides a very rich API for dealing with all kind of dates in a
// meaningful and sensible way. Some notable features include:
//
// 1. All dates, even though that specify an specific day have a minimum and
// maximum value that are their true bounds. This is especially important for
// larger date ranges like the whole month of "Jun 1945".
//
// 2. Upper and lower bounds of dates can be converted to the native Go
// time.Time object.
//
// 3. There is a Years function that provides a convenient way to normalise a
// date range into a number for easier distance and comparison measurements.
//
// 4. Algorithms for calculating the similarity of dates on a configurable
// parabola.
//
// Node Equality
//
// Node.Equals performs a shallow comparison between two nodes. The
// implementation is different depending on the types of nodes being compared.
// You should see the specific documentation for the Node.
//
// Equality is not to be confused with the Is function seen on some of the
// nodes, such as Date.Is. The Is function is used to compare exact raw values
// in nodes.
//
// DeepEqual tests if left and right are recursively equal.
//
// Tree Walking & Filtering
//
// The Filter function recursively removes or manipulates nodes with a
// FilterFunction:
//
//   newNodes := gedcom.Filter(node, func (node gedcom.Node) (gedcom.Node, bool) {
//     if node.Tag().Is(gedcom.TagIndividual) {
//       // false means it will not traverse children, since an
//       // individual can never be inside of another individual.
//       return node, false
//     }
//
//     return nil, false
//   })
//
//   // Remove all tags that are not official.
//   newNodes := gedcom.Filter(node, gedcom.OfficialTagFilter())
//
// Some examples of Filter functions include BlacklistTagFilter,
// OfficialTagFilter, SimpleNameFilter and WhitelistTagFilter.
//
// Individual Names
//
// A NameNode represents all the parts that make up a single name. An individual
// may have more than one name, each one would be represented by a NameNode.
//
// Apart from functions to extract name parts there is also Format which works
// similar to `fmt.Printf` where placeholders represent different components of
// the name:
//
//   %% "%"
//   %f GivenName
//   %l Surname
//   %m SurnamePrefix
//   %p Prefix
//   %s Suffix
//   %t Title
//
// Each of the letters may be in upper case to convert the name part to upper
// case also. Whitespace before, after and between name components will be
// removed:
//
//   name.Format("%l, %f")     // "Smith, Bob"
//   name.Format("%f %L")      // "Bob SMITH"
//   name.Format("%f %m (%l)") // "Bob (Smith)"
//
package gedcom
