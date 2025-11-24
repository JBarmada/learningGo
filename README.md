# learningGo
I want to learn Go :D

# stuff i learned so far
- modules
- functions
- variables
- slices

## modules
import “fmt” <- functions for formatting text

import "math/rand" <- standard random gens

import "log" <- logging made easy

import "errors" <- err module

can hard code a module location if the module functions should be found locally, using the edit -replace command.

## functions
private functions start with a lowercase letter

public functions start with a capital letter

built-ins are lowercase, and are actually keywords the compiler knows

## declaring variables
“var x string

x = value”

is exactly like:
“x := value”

## slices
slice : dynamic array  “var  []type”

Has the following useful functions: 
Name of functions makes it pretty obvious

append(), copy(), make([]T, l, c), appendByte([]byte, data byte)

