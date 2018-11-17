// Package q is the gedcomq parser and engine.
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
// Examples
//
// Count all individuals in a document:
//
//   .Individuals | Length
//
package q
