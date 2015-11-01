# Universal Configuration Grammar

**Disclaimer**

>I am not a compiler engineer or designer.  I am a somewhat generalist software
>developer with a degree in computer science and some understanding of theory
>of computation, grammars, etc.  This is a project I am working on for my own
>enjoyment and experience.  If you are reading any of this and think I am doing
>something wrong or in a way that could be improved, I will gladly accept
>suggestions in Github issues or, even better, pull requests.

## Syntax

UConfig files use a distinctly
[Scheme](https://en.wikipedia.org/wiki/Scheme_%28programming_language%29)-like
syntax.  In the simplest terms, this means that everything is an S-Expression.
UConfig should have as few distinct syntactic forms as possible.  Some examples
of the syntax we need to have:

1. Literals

    "strings"
    42
    3.141592

2. S-Expressions: 

    (function argument1 argument2)

3. Quoted Lists

    `(1 2 3 4 5)
    `("hello" "world" "!")

4. Comments

    ; This won't be interpreted

## Special Forms

Since we are defining a language that can be executed, we need a few special
forms that will have special behavior associated with them that won't be
available in the language itself.

### Definitions

    (define (name1 value1) (name2 value2) ...)

Our language needs a way to assign values to names to whatever scope
that our `define` form is invoked within.

### Conditions

    (if condition then else)

Conditionals should exist for branching logic.

### Functions

    (function (argument1 argument2 ...) body)

The function form has to handle the first S-expression as a list of
parameter names and the second as a function body.

### Quotes

    `(1 2 3 4)

Quotes introduce a special structure. In particular, a list.

