// Package gedcomq is a command line tool and query language for GEDCOM files.
// It is heavily inspired by jq, in name and syntax.
//
// The basic syntax of the tool is:
//
//   gedcomq -gedcom file.ged '.Individuals | .Name'
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
// Length -- Returns an integer with the number of items in the slice. This
// value will be 0 or more. If the input is not a slice then 1 will always be
// returned.
//
// Here is an example of counting all individuals in a document:
//
//   .Individuals | Length
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
package main
