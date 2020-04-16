// Package q is the "gedcom query" parser and engine.
//
// Language Basics
//
// The query is split into expressions. The pipe (|) indicates that the result
// of one expression is the input into the next expression.
//
// The starting expression is the gedcom.Document itself that is passed into the
// first expression (".Individuals" in the example above).
//
// ".Individuals" is called an "accessor", denoted by the "." prefix. An
// accessor will try to find a property or method of that name, returning the
// value of the property or the result of invoking the method. The above example
// would return a slice ([]*IndividualNode).
//
// The next expression, ".Name" receives that slice. Since it is a slice the
// ".Name" accessor is performed on each of the individual slice members,
// creating a new slice with the results. In this case IndividualNode has a
// method called Name that returns a *NameNode. That means that result of the
// processing the slice will be []*NameNode.
//
// After all of the expressions have been evaluated the result is encoded into
// JSON and output.
//
// It's important to note that some structures implement the json.Marshaller
// interface which controls how the structure is represented in JSON. Many
// structures also implement fmt.Stringer (the String method) which can be
// helpful for seeing more simple representations of values.
//
// With the example ".Individuals | .Name" on a document that contains two
// individuals:
//
//   [
//     {
//       "Nodes": [
//         {
//           "Tag": "GIVN",
//           "Value": "Lucy Alcott"
//         },
//         {
//           "Tag": "SURN",
//           "Value": "Chauncey"
//         }
//       ],
//       "Tag": "NAME",
//       "Value": "Lucy Alcott /Chauncey/"
//     },
//     {
//       "Nodes": [
//         {
//           "Tag": "GIVN",
//           "Value": "Sarah"
//         },
//         {
//           "Tag": "SURN",
//           "Value": "Taylor"
//         }
//       ],
//       "Tag": "NAME",
//       "Value": "Sarah /Taylor/"
//     }
//   ]
//
// If this is too verbose for you, here is the same output using
// ".Individuals | .Name | .String":
//
//   [
//     "Lucy Alcott Chauncey",
//     "Sarah Taylor"
//   ]
//
// Functions
//
// Some functions are provided as part of the gedcomq language that exist
// outside of the gedcom package:
//
//   Combine(Slices...)
//
// Combine will combine multiple slices of the same type into a single slice.
//
//   First(number)
//
// First returns up to the number of elements in a slice.
//
// If the input value is not a slice then it is converted into a slice of one
// element before evaluating. This means that the result will always be a slice.
// The only exception to this is if the input is nil, then the result will also
// be nil.
//
// There must be exactly one argument and it must be 0 or greater. If the number
// is greater than the length of the slice all elements are returned.
//
//   Last(number)
//
// Last returns up to the number of elements in a slice.
//
// If the input value is not a slice then it is converted into a slice of one
// element before evaluating. This means that the result will always be a slice.
// The only exception to this is if the input is nil, then the result will also
// be nil.
//
// There must be exactly one argument and it must be 0 or greater. If the number
// is greater than the length of the slice all elements are returned.
//
//   Length
//
// Length returns an integer with the number of items in the slice.
//
// This value will be 0 or more. If the input is not a slice then 1 will always
// be returned.
//
//   MergeDocumentsAndIndividuals(doc1, doc2)
//
// Merges two documents while also merging similar individuals.
//
//   NodesWithTagPaths(Tags...)
//
// NodesWithTagPath returns all of the nodes that have an exact tag path. The
// number of nodes returned can be zero and tag must match the tag path
// completely and exactly.
//
// Find all Death nodes that belong to all individuals:
//
//   .Individuals | NodesWithTagPath("DEAT")
//
// From the individuals find all the Date nodes within only the Birth nodes.
//
//   .Individuals | NodesWithTagPath("BIRT", "DATE")
//
// Combine all of the birth and death dates:
//
//   Births are .Individuals | NodesWithTagPath("BIRT", "DATE") | {type: "birth", date: .String};
//   Deaths are .Individuals | NodesWithTagPath("DEAT", "DATE") | {type: "death", date: .String};
//   Combine(Births, Deaths)
//
// If the node is nil the result will also be nil.
//
//   Only(condition)
//
// The Only function returns a new slice that only contains the entities that
// have returned true from the condition. For example:
//
//   .Individuals | Only(.Age > 100)
//
// The Question Mark
//
// "?" is a special function that can be used to show all of the possible next
// functions and accessors. This is useful when exploring data by creating the
// query interactively.
//
// For example the following query:
//
//   .Individuals | ?
//
// Returns (most items removed for brevity):
//
//   [
//     ".AddNode",
//     ".Age",
//     ".AgeAt",
//     ...
//     ".SurroundingSimilarity",
//     ".Tag",
//     ".Value",
//     "?",
//     "Length"
//   ]
//
// Variables
//
// Variables allow more complex logic to be processed in separate discreet
// steps. It also applies in cases where the logic would normally be duplicated
// if it couldn't be referenced from multiple places.
//
// Variable are defined in on of the two forms:
//
//   Events are .Individuals | .AllEvents
//   Name is .Individual | .Name
//
// The keywords "are" and "is" do exactly the same thing. They are both offered
// to make the semantics of reading the expression easier.
//
// Variables can then be references in separate expressions. For example the
// following:
//
//   .Individuals | .Name | .String
//
// Could also be written as:
//
//   Names are .Individuals | .Name; Names | .String
//
// Or even more verbosely as:
//
//   Indi is .Individuals; Names are Indi | .Name; Names | .String
//
// The semicolon (;) is used to separate variable definitions. The result
// returned will always be the return value of the last statement.
//
// Available variables will be shown as options with the special Question Mark
// function.
//
// Data Types
//
// gedcomq does not define strict data types. Instead it will perform an
// operation as best it can under the conditions provided.
//
// To help simplify things here are general descriptions of how certain data
// types are handled:
//
// - Numbers can be actual whole of floating-point numbers, or they can also be
// represented as a string. For example 1.23 and "1.230" are considered equal
// because they both represent the same numerical value, even though they are in
// different forms.
//
// - Strings are text of any length (including zero characters). If it's value
// represents a number, such as "123" or "4.56" it will change the behaviour of
// the operator used on it because they will be treated as numbers rather than
// text. It's also very important to note that strings are compared internally
// without case-sensitivity and whitespace that exists at the start or end of
// the string will be ignore. For example "John Smith" is considered to be equal
// to "  john SMITH ".
//
// - Slices are an ordered set of items, often also called an "array". The name
// was chosen as "slice" rather than "array" because it is more inline with the
// description of types in Go. A slice may contain zero elements but if it does
// have items they will almost certainly be of the same type. Such as a slice of
// individuals.
//
// - Objects (sometimes referred to as a "map" or "dictionary") consists as a
// zero or more key-value pairs. The values may be of any type, but the keys are
// always strings and always unique in that object. Objects may be generic, or
// they may be a specific type from the gedcom package. If they are a specific
// type, such as an IndividualNode they may also have methods available which
// can be accessed just like properties.
//
// Operators
//
// gedcomq supports several binary operators that can be used for comparison of
// values. All operators will return a boolean (true/false) result:
//
//   =  (equal)
//   != (not equal)
//
// If the left and right both represent numeric values then the values are
// compared numerically. That is to say 1.23 and "1.2300" are equal.
//
// If either the left or right is not a number then the values are compared
// without case and any whitespace at the start or end is ignore. This means
// that "John Smith" is considered to be equal to "  john SMITH ", but not equal
// to "John  Smith".
//
// Not equal works exactly opposite.
//
//   >  (greater than)
//   >= (greater than or equal)
//   <  (less than)
//   >= (less than or equal)
//
// If the left and right both represent numeric values then the values are
// compared numerically. That is to say 1.2301 is greater than "1.23".
//
// If the left or right does not represent a numeric value then the values are
// compared as strings using the same case-insensitive rules as "=".
//
// One string is greater than another string by comparing each of the
// characters. So "Jon" is greater than "John" because "n" is greater than "h".
//
// Creating Objects
//
// Custom objects can be constructed on one more items. For example:
//
//   .Individuals | { name: .Name | .String, born: .Birth | .String }
//
// May output something similar to:
//
//   [
//     {
//       "born": "1863",
//       "name": "Charles W Chauncey"
//     },
//     {
//       "born": "12 Dec 1859",
//       "name": "Lucy Alcott Chauncey"
//     },
//     {
//       "born": "1831",
//       "name": "Sarah Taylor"
//     }
//   ]
//
// It's also worth noting that object can contain zero key-value pairs, such as:
//
//   .Individuals | {}
//
// This would output (using the same individuals in the previous example):
//
//   [
//     {},
//     {},
//     {}
//   ]
//
// Also see the Examples below.
//
// Outputting In Other Formats
//
// There are several formatters (see Formatter interface) that allow the result
// of a query to be output in different ways. Such as pretty json or CSV.
//
// This can be controlled with the "-format" option with gedcomq, or by
// instantiating one of the formatter instances in your own code.
//
// Examples
//
// Count all individuals in a document:
//
//   .Individuals | Length
//
// result:
//
//   3401
//
// Retrieve the basic details of the first 3 individuals:
//
//   .Individuals | First(3) | { name: .Name | .String, born: .Birth | .String, died: .Death | .String}
//
// result:
//
//   [
//     {
//       "born": "6 Dec 1636",
//       "died": "2 Dec 1713",
//       "name": "Gershom Bulkeley"
//     },
//     {
//       "born": "5 Nov 1592",
//       "died": "19 Feb 1672",
//       "name": "Charles Chauncey"
//     },
//     {
//       "born": "1408",
//       "died": "7 May 1479",
//       "name": "John Chauncy Esq."
//     },
//   ]
//
// Retrieve the names of individuals that have a given name (first name) of
// "John".
//
//   .Individuals | .Name | Only(.GivenName = "John") | .String
//
// result:
//
//   [
//     "John Chaunce",
//     "John Chaunce",
//     "John Chance",
//     "John Unett",
//     "John Chance",
//     "John de Chauncy",
//   ]
//
// Find all of the living people with their current age:
//
//   .Individuals | Only(.IsLiving) | { name: .Name | .String, age: .Age | .String}
//
// result:
//
//   [
//     {
//       "age": "82y 6m",
//       "name": "Robert Walter Chance"
//     },
//     {
//       "age": "~ 90y 10m",
//       "name": "Sir Robert Temple Armstrong"
//     },
//   ]
//
// Merge two GEDCOM files (full command):
//
//   gedcomq -gedcom file1.ged -gedcom file2.ged -format gedcom \
//     'MergeDocumentsAndIndividuals(Document1, Document2)' > merged.ged
//
package q
