# golangintro

A comprehensive one day Golang workshop for
[devopenspace](https://devopenspace.de/) 2018.

## Overview

This workshop contains slides, code, examples and exercises. Target audience
are people familiar with *some* programming language.

The following sections are available:

* [Golang Origins](Origins.md)
* [Installation](Installation.md)
* [Editors](Editors.md)
* [Go Workspace](Workspace.md)
* [Hello World](Hello.md)
* [Basic data types](Types.md)
* [Basics tour](Basics.md) (17)
* [Flow control tour](Flowcontrol.md) (14)
* [More types tour](Motetypes.md) (27)
* [Methods and Interfaces](Methods.md) (26)
* [Concurrency](Concurrency.md) (11)
* [Working with Files](Files.md)

----

* Working with files
* The io package
* Working with XML and JSON
* The sync and context packages
* Testing code
* Writing web servers
* Notable projects, tools, helpers
* Writing documentation (doc.go, godoc)
* Further resources (Docs, Proverbs)

## Exercises

### Small web server

1. Write a small HTTP server that returns the current time and the number of
   times it has been called.
2. Extend the server to return JSON.

### A website inspection tool

1. Write a small tool for making requests to a number of websites. For each
   website, report the number of bytes in the response body and time it took to
   fetch the resource. The list of URLs can be hardcoded.
2. Take a file with a list of files as parameters (package flag, bufio).
3. Make concurrent requests.
4. Allow to set a timeout.

## What makes Go special?

* It imposes a workspace structure.
* Every type has a zero value. A proverb says: Make the zero value useful.
* Error handling without exceptions.
* Multiple return values.
* Encourages tools for code (linter, struct generator).
* A nil value in an interface can be handled gracefully.