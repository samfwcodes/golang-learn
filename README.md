# Go Basics

This document covers the fundamental concepts of Go in simple language. It is intended as quick revision notes rather than complete documentation.
COMPLEX TOPICS LIKE INTERFACES, CHANNELS AND GOROUTINES ARE IN THERE INDIVIDUAL REPO

---

# Why Go?

Go is designed to be:

- 🚀 Fast
- 🧹 Simple and easy to read
- ⚡ Excellent for concurrency
- 📦 Great for backend development
- 🤝 Backed by a strong community

---

# Variables

Variables are used to store values.

## Declaration

### 1. Explicit Type

```go
var age int
var isAdmin bool
```

Default values:

```go
int  -> 0
bool -> false
string -> ""
```

---

### 2. Type Inference

Go automatically detects the type.

```go
var age = 20
```

---

### 3. Short Declaration (Most Common)


```go
age := 20
name := "Sam"
```

Can only be used **inside functions**.

---

## Important Notes

- Variables are mutable.
- Variables cannot be redeclared in the same scope.
- Variables can be shadowed inside a nested scope.

```go
name := "Sam"

{
    name := "Harsh"
}
```

---

## Naming Conventions

Good:

```go
user
userName
totalPrice
```

Bad:

```go
a
abcxyz
thisIsMyVeryVeryLongVariableName
```

Short names are fine when the variable only lives for a few lines.

---

# Type Conversion

Convert one type into another.

```go
var price float64 = 30.5

count := int(price)
```

Use `strconv` when converting between numbers and strings.

```go
strconv.Itoa(65)
strconv.Atoi("65")
```

---

# Primitive Types

## Boolean

```go
true
false
```

Default value:

```go
false
```

---

## Integer

```go
int
int8
int16
int32
int64
```

Unsigned integers:

```go
uint
uint8
uint16
uint32
uint64
```

`uint` cannot store negative values.

---

## Floating Point

```go
float32
float64
```

Examples:

```go
3.14
12e5
3.2e8
```

---

## Complex Numbers

```go
var c complex64 = 1 + 2i
```

Useful functions:

```go
complex()
real()
imag()
```

---

## String

Strings are:

- UTF-8 encoded
- Immutable
- Concatenated using `+`

```go
name := "Sam"
```

Convert to bytes:

```go
[]byte(name)
```

---

## Rune

Represents one Unicode character.

```go
var r rune = 'A'
```

---

# Constants

Constants cannot be modified.

```go
const Pi = 3.14
```

---

## iota

Useful for creating enums.

```go
const (
    Red = iota
    Green
    Blue
)
```

Output:

```
Red   = 0
Green = 1
Blue  = 2
```

---

# Arrays

Arrays have a fixed size.

```go
arr := [3]int{1,2,3}
```

or

```go
arr := [...]int{1,2,3}
```

Properties:

- Fixed length
- Value type (copied on assignment)

---

# Slices

Slices are built on top of arrays.

```go
nums := []int{1,2,3}
```

Create with `make()`:

```go
nums := make([]int, 3, 10)
```

Properties:

- Dynamic size
- Reference type
- Can use `append()`

```go
nums = append(nums,4)
```

Slice syntax:

```go
nums[1:3]
```

---

# Maps

Maps store key-value pairs.

```go
grades := map[string]int{
    "Sam":95,
}
```

Create using:

```go
make(map[string]int)
```

---

## Operations

Insert

```go
grades["John"] = 99
```

Delete

```go
delete(grades,"John")
```

Check existence

```go
value, ok := grades["Sam"]
```

---

# Structs

Structs group multiple fields together.

```go
type User struct {
    Name string
    Age int
}
```

---

## Embedding

```go
type Animal struct {
    Name string
}

type Bird struct {
    Animal
    CanFly bool
}
```

---

## Export Rules

Capitalized names are exported.

```go
type User struct{}
```

Lowercase names remain package-private.

```go
type user struct{}
```

---

# Loops

Go only has the `for` keyword.

## Traditional

```go
for i := 0; i < 5; i++ {
}
```

---

## While Style

```go
i := 0

for i < 5 {
    i++
}
```

---

## Infinite Loop

```go
for {
}
```

---

## Range Loop

```go
nums := []int{1,2,3}

for index, value := range nums {
    fmt.Println(index,value)
}
```

---

## Loop Keywords

Exit loop

```go
break
```

Skip current iteration

```go
continue
```

---

# Bitwise Operators

Works directly on binary values.

```go
&
|
^
&^
```

Example:

```go
10 & 3
10 | 3
10 ^ 3
```

---

## Bit Shift

Left Shift

```go
a << 2
```

Right Shift

```go
a >> 2
```

---

# Defer

`defer` delays execution until the surrounding function returns.

```go
defer fmt.Println("Done")
```

Execution order:

- Last In
- First Out (LIFO)

Example:

```go
defer fmt.Println(1)
defer fmt.Println(2)
defer fmt.Println(3)
```

Output

```
3
2
1
```

---

# Panic

`panic()` immediately stops normal execution.

```go
panic("Database connection failed")
```

Use only for unrecoverable errors.

---

# Recover

`recover()` catches a panic and prevents the program from crashing.

Usually used with `defer`.

```go
defer func() {
    if err := recover(); err != nil {
        fmt.Println(err)
    }
}()
```

---

# Functions

Functions are reusable blocks of code that perform a specific task.

---

## Declaring a Function

```go
func greet() {
    fmt.Println("Hello, World!")
}
```

Calling the function:

```go
greet()
```

---

## Function Parameters

Functions can accept one or more parameters.

```go
func greet(name string) {
    fmt.Println("Hello", name)
}
```

Calling the function:

```go
greet("Sam")
```

Multiple parameters of the same type:

```go
func add(a, b int) {
    fmt.Println(a + b)
}
```

---

## Return Values

Functions can return values.

```go
func add(a, b int) int {
    return a + b
}
```

Usage:

```go
result := add(10, 20)
```

---

## Multiple Return Values

Go functions can return multiple values.

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }

    return a / b, nil
}
```

Usage:

```go
result, err := divide(10, 2)
```

This pattern is commonly used for error handling.

---

## Named Return Values

Return values can have names.

```go
func rectangle(length, width int) (area int) {
    area = length * width
    return
}
```

Named returns are supported but are generally avoided unless they improve readability.

---

## Variadic Functions

A variadic function accepts any number of arguments.

```go
func sum(nums ...int) int {
    total := 0

    for _, num := range nums {
        total += num
    }

    return total
}
```

Usage:

```go
sum(1, 2, 3)
sum(5, 10, 15, 20)
```

---

## Anonymous Functions

Functions without a name.

```go
func() {
    fmt.Println("Hello")
}()
```

Anonymous functions are useful for short tasks and goroutines.

---

## Function as a Value

Functions can be assigned to variables.

```go
greet := func(name string) {
    fmt.Println("Hello", name)
}

greet("Sam")
```

---

## Passing Functions

Functions can be passed as arguments.

```go
func calculate(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}
```

---

## Key Notes

- Functions can have zero or many parameters.
- Functions can return one or multiple values.
- Functions are first-class citizens (they can be assigned, passed, and returned).
- Error handling in Go usually uses multiple return values.

---

# Methods

A method is a function that belongs to a specific type.

Methods allow you to define behavior for structs or custom types.

---

## Defining a Method

```go
type User struct {
    Name string
}

func (u User) Greet() {
    fmt.Println("Hello", u.Name)
}
```

Usage:

```go
user := User{Name: "Sam"}

user.Greet()
```

---

## Method Receiver

The value inside parentheses is called the **receiver**.

```go
func (u User) Greet() {}
```

Here:

- `u` → receiver variable
- `User` → receiver type

---

## Value Receiver

A value receiver works on a copy of the struct.

```go
type Counter struct {
    Count int
}

func (c Counter) Increment() {
    c.Count++
}
```

```go
counter := Counter{}

counter.Increment()

fmt.Println(counter.Count)
```

Output:

```
0
```

The original value is unchanged because a copy was modified.

---

## Pointer Receiver

A pointer receiver modifies the original struct.

```go
type Counter struct {
    Count int
}

func (c *Counter) Increment() {
    c.Count++
}
```

```go
counter := Counter{}

counter.Increment()

fmt.Println(counter.Count)
```

Output:

```
1
```

---

## When to Use Pointer Receivers

Use pointer receivers when:

- The method modifies the struct.
- The struct is large and copying it is inefficient.
- You want all methods to have consistent behavior.

---

## Methods on Custom Types

Methods can also be attached to custom types.

```go
type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
    return float64(c)*9/5 + 32
}
```

---

## Methods vs Functions

Function:

```go
func greet(name string)
```

Method:

```go
func (u User) Greet()
```

A method always has a **receiver**.

---

## Key Notes

- Methods are functions attached to a type.
- A receiver can be a value or a pointer.
- Pointer receivers modify the original value.
- Value receivers work on copies.
- Methods improve code organization and readability.

# Summary

This README covered:

- Variables
- Primitive Types
- Constants
- Arrays
- Slices
- Maps
- Structs
- Loops
- Bitwise Operators
- Defer
- Panic
- Recover
- Functions
- Methods
