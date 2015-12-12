# An introduction to Fig

Hello and welcome, dear reader! It is my hope that you have found yourself reading this document after having browsed through the project's README.

Here, you will get a very quick overview of the Fig language, which you will be able to write powerful executable configuration-generating programs with.

**Contents**

* [Syntax](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#syntax)
  * [Variable names](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#variable-names)
  * [Literals](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#literals)
  * [Comments](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#comments)
  * [Calling functions](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#calling-functions)
  * [Special forms](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#special-forms)
* [Standard Library](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#standard-library)
  * [Table of functions](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#standard-library)
  * [Mathematical operations](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#math)
  * [String functions](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#strings)
  * [Boolean operations](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#booleans)
  * [Lists and list operations](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#lists)
  * [Maps and map operations](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#maps)
  * [IO - Input and output](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#io)
* [Functional Programming](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#functional-programming)
  * [Recursion](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#recursion)
  * [Closures](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md#closures)

## Syntax

Since Fig is 100% a [Lisp](https://en.wikipedia.org/wiki/Lisp_%28programming_language%29) dialect, its syntax consists of [S-Expressions](https://en.wikipedia.org/wiki/S-expression) as well as a few special forms, described below.

### Variable names

Fig is very unrestrictive in the kinds of variable names that are allowed.  You can name them anything including the characters a to z, A to Z, 0 to 9, and any of `!, @, #, $, %, ^, &, *, -, _, +, =, :, <, >, ., ?, /` and the comma `,`.

As such, each of the following are valid and clear names.

* isZero?
* +
* count-to-n
* !=

### Literals

Fig supports three fundamental types of literals.

1. Numbers
  1. Integers such as 0, 1, 103, 9001, -123, and so on
  2. Floats such as 3.14, -2.5, and 0.999
2. Strings
  1. Double-quoted such as "string1" can contain single-quotes like "I said 'bye!' to them."
  2. Single-quoted such as 'string2' can contain double-quotes like 'and then we "laughed".'
3. Boolean keywords true and false

### Comments

Everything following a semi-colon `;` is considered part of a comment and terminates at the end of the line.

### Calling functions

A function is called by simply writing an S-Expression where the first name is the name of the function, and all successive names (or literals) are passed to the function as arguments.

For example, given a function `multiply` that accepts two numbers to multiply, you would write the following to invoke the function.

    (multiply x y)

So that `x` and `y` become arguments to the `multiply` function.

### Special forms

Fig defines three special forms which each have a somewhat special syntax and are each handled in a particular way.

You can see each of the special forms described here in use in the [examples/specialforms.fig](https://github.com/redwire/UnicornFig/blob/master/examples/specialforms.fig) example script.

1. `define` assigns a value to a name
2. `if` begins a conditional branch
3. `function` creates a function that can be called later

#### Define

The syntax of `define` is as follows:

```
(define (name1 expression1) (name2 expression2) ...)
```

As indicated by the ellipses `...`, `define` accepts any number of S-Expressions of the form `(name expression)` where the name is assigned the value resulting from evaluating the provided expression.

#### If

The syntax of `if` is as follows:

```
(if condition then else)
```

It works by evaluating the provided `condition` expression and, if it resolves to `true`, evaluates and returns the `then` expression. Otherwise it evaluates and returns the `else` expression.

#### Function

The syntax of `function` is as follows:

```
(function (argument1 argument2 ...) body)
```

The first S-Expression provided is treated as a list of argument names, and the `body` is an expression that will be evaluated and returned when the function is invoked.

## Standard Library

Math | Strings    | Booleans | Lists   | Maps      | IO
-----|------------|----------|---------|-----------|------
`*`  | `concat`   | `=`      | `list`  | `mapping` | `print`
`/`  | `substr`   | `not`    | `first` | `assoc`   | `env`
`+`  | `index`    | `and`    | `tail`  | `get`     | `ignored`
`-`  | `length`   | `or`     | `append`| `keys`    |
`%`  | `upcase`   |          | `size`  |
`>`  | `downcase` |
`<`  | `split`    |
`>=` | `at`
`<=` |
`zero?` |

### Math

#### * (a, b, ... integer)

Multiplies two or more numbers together.  If any of the arguments are floats, the result will be a float.

#### / (a, b, ... integer)

Divides two or more numbers.  If any of the arguments are floats, the result will be a float.

Note that calling `(/ 12 4 2)` is calculated like `(12 / 4) / 2`.

#### + (a, b, ... integer)

Adds two or more numbers together. If any of the arguments are floats, the result will be a float.

#### - (a, b, ... integer)

Subtracts two or more numbers. If any of the arguments are floats, the result will be a float.

Note that calling `(- 3 2 1)` is calculated like `(3 - 2) - 1`.

#### % (a, b integer)

Calculates the modulo of exactly two integers. Floats are not accepted.

#### > (a, b number)

Tests if one number is greater than another. Works with floats and integers.

#### < (a, b number)

Tests if one number is less than another. Works with floats and integers.

#### >= (a, b number)

Tests if one number is greater than or equal to another. Works with floats and integers.

#### <= (a, b number)

Tests if one number is less than or equal to another. Works with floats and integers.

#### zero? (n number)

Tests if its argument is 0. Works on both integers and floats.

### Strings

#### concat (s1, s2, ... string)

Concatenates two or more strings together.

#### substr (s string, start, end integer)

Produces a substring. The first argument is a string. The second is an integer representing the index to start at (0-based and inclusive). The third argument is an integer representing the index to stop at (non-inclusive).

#### index (haystack, needle string)

Determines the first index at which a substring can be found within another string. The first argument is the string to search in. The second argument is the substring to search for. If the substring is not found, returns `-1`.

#### length (s string)

Calculates the length of a single string as an integer.

#### upcase (s string)

Converts a single string to uppercase.

#### downcase (s string)

Converts a single string to lowercase.

#### split (s, separator string)

Splits the first string argument by a second string separator.

#### at (s string, index integer)

Get the character of a first string argument at some integer index.

### Booleans

#### = (a, b, ... number/boolean/string)

Test if two or more values are equal. Works with bools, numbers, and strings. All arguments must be the same type.

#### not (b boolean)

Negates a single boolean value.

#### and (a, b, ... boolean)

Determines whether all of the arguments are true. Only works with booleans.

#### or (a, b, ... boolean)

Determines whether any of the arguments are true. Only works with booleans,

### Lists

#### list (v1, v2, ... any)

Create a list composed of all of the values provided as arguments in order.

#### first (l list)

Retrieves the first element of a single list argument.

#### tail (l list)

Retrieves a sub-list containing all but the first element of a single list argument.

#### append (l list, v1, v2, ... any)

Creates a new list with all the argument values following the first list argument appended to that list.

#### size (l list)

Computes the number of elements in a single list argument as an integer.

### Maps

#### mapping (k1 string, v1 any, k2 string, v2 any, ...)

Creates a map/dictionary associating values to string keys. Each even-indexed argument (starting from 0) must be a string that wil be a key mapping to the following value.

Example

```js
(mapping "hello" 3.14 "world" 2) => {"hello": 3.14, "world": 2}
```

#### assoc (m map, k1 string, v1 any, k2 string, v2 any, ...)

Adds new key-value pairs to a map. The first argument must be a map and then all successive arguments must be string keys followed by the corresponding value.

Example

```js
(assoc (mapping "hello" 3.14) "world" 2 "!" "woah") => {"hello": 3.14, "world": 2, "!": "woah"}
```

#### get (m map, key string)

Retrieves the value associated with a given key from a map. The first argument is a map. The second argument is a string key.

Example

```js
(get (mapping "a" 1 "b" 2) "a") => 1
```

#### keys (m map)

Get a list of the keys in a map as a list of strings.

### IO

#### print (v1, v2, ... any)

Prints any number of values to the console.

#### env (name string)

Returns the value of the environment variable corresponding to the provided name as a string.
For example, the following program

```js
; test.fig
(print (env "WHATWASSET"))
```

when run like

```bash
WHATWASSET=thishaha ./unicorn test.fig
```

outputs `thishaha`.

#### ignored (value any)

By calling `ignored` on a value, Unicorn will set a flag on the variable internally to prevent that value from appearing
in output files.  Note that if new values are created by passing ignored values into functions to produce a new value,
the resulting value **will not** be ignored. That is, you must explicitly ignore each value created.

For example, the following program

```js
(define
    (baseUrl (ignored "localhost:3000"))
    (index (concat "http://" baseUrl "/index.html")))
```

would produce a JSON file containing only

```js
{
    "index": "http://localhost:3000/index.html"
}
```

when run.

## Functional Programming

Fig is a purely functional programming language, much like [Haskell](https://en.wikipedia.org/wiki/Haskell_%28programming_language%29).  
What this means is that functions cannot have side effects. They are only able to perform computations on their input and produce a single output (which may be a list or map).

This classification is slightly relaxed, however, considering that functions can bind (assign) values to names more than once and printing to the console is possible without the use of an IO monad.

You can see some cool examples of functional programming in [examples/functional.fig](https://github.com/redwire/UnicornFig/blob/master/examples/functional.fig)

### Recursion

Fig does not support iteration like what is offered by for-loops and while-loops from C-like languages. Instead, programmers must write recursive functions. Doing so, however, is quite simple.

For example, consider the `factorial` function below that computes a factorial
by recursively multiplying `n * (n - 1) * (n - 1 - 1) * ... * 1`

```js
(define
  (factorial (function (n)
    (if (= n 1)
        1
        (* n (factorial (- n 1)))))))
```

Another way to write this would be to use an accumulator argument like so.

```js
(define
  (recursive-factorial (function (n accumulator)
    (if (= n 1)
        accumulator
        (recursive-factorial (- n 1) (* accumulator n)))))
  (factorial (function (n)
    (recurisve-factorial n 1))))
```

You can see some more examples of recursion in Fig in [examples/recursion.fig](https://github.com/redwire/UnicornFig/blob/master/examples/recursion.fig)

### Closures

[Closures](https://en.wikipedia.org/wiki/Closure_%28computer_programming%29) are functions that retain the state of the environment they are created in.  Often one will write a function that accepts some arguments and returns a function that accepts more arguments and operates on both sets of arguments.

Consider the following example of a the `multiplier` function which accepts an argument, and then returns a new function that, when called with another argument, multiplies the two together. This pattern allows us to easily create funcions that multiply an argument by a constant value.

```js
(define
  (multiplier (function (n)
    (function (m)
      (* n m)))))

(define (double (multiplier 2))
        (triple (multiplier 3)))
```

As you might expect, the `double` and `triple` functions can now be used to multiply whatever arguments are provided to them by `2` and `3` respectively.

You can see this example in action in [examples/closures.fig](https://github.com/redwire/UnicornFig/blob/master/examples/closures.fig)
