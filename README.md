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
* [Basics tour](Basics.md)
* [Flow control tour](Flowcontrol.md)
* [More types tour](Motetypes.md)

----

* Language overview and design documents
* The first program, explained
* Basic data types
* Slices
* Maps
* Compound data types
* Functions
* Methods
* Interfaces
* Goroutines and Channels
* Select
* The sync and context packages
* Writing web servers
* Notable projects, tools, helpers
* Further resources

## Exercises

### Small web server

1. Write a small HTTP server that returns the current time and the number of
   times it has been called.

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
