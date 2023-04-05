# Go Complete Developer Guide

This project serves as an overview of all Go language features

## 1 Hello World

A simple example of printing to the console in Golang

## 2 Cards

An example that revolves around the functionality of a deck of cards, for example: shuffling and dealing.

Explores:

- type declaration
- variable declaration
- function declaration (with parameters and as receiver)
- An example of a For loop
- Standard library OS to write and read files
- How to test the above using a test file

## 3 Structs

Exploring the struct datatype in Go in greater detail than the above example.

Structs are similar to objects in other programming languages, where there are keys with values. The values can be simple datatypes (eg. string) or complex types (structs).

Contains receiver functions as well.

## 4 Assignment 1: For Loops

Using conditions within a for loop to interpret a list/array/slice as you go through it.

## 5 Maps

Both Maps and structs are similar to objects in other programming languages.

The primary difference between the two is that structs are capable of having different types within the same map. Whereas a map is declared from the start with a certain data type for the key and the value. Both of which are unchangeable.

## 6 Interfaces

Interfaces allow multiple structs to be placed into categories. This allows Golang to reuse code which may be applicable to separate types.

## 7 HTTP

Uses the net/http standard library to perform a GET request. Then prints the result to the console.

## 8 Assignment 2: Area

Using structs and interfaces in combination.

## 9 Assignment 3: Command line arguments

Uses OS standard library to grab arguments from command line for use within code.

## 10 Website Checker: Go Routines and Channels

Go routines are Golang's method of running concurrent processes.

Go channels (a messaging system used between go routines) must also be used to ensure that all processes are aware of each other to ensure they all complete before the main process completes.

Channels use a unique format:

- Example below of assigning a message to a channel

```
channel<-
```

- Example below of waiting and grabbing a channel's message

```
<-channel
```

## Gorm_service

An example of the clean architecture pattern where three layers are used.
-Controller
-Service
-Repository
