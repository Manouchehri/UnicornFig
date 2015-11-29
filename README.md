# Unicorn... Fig? What the heck?

Quite right, my good man!

**Unicorn Fig** is a really terrible pun on the words "universal" and "configuration"!

## Now you have my attention

Hah! I doubt that! But let's pretend I do for a moment, shall we?

UnicornFig is a side project of mine with the rather ambitious goal of becoming a really simple "Emacs Lisp for everyone."  What that essentially means is it hopes to not just be *yet another Lisp implementation*, but rather a Lisp interpreter that outputs configuration files in familiar formats like [JSON](https://en.wikipedia.org/wiki/JSON) and [YAML](https://en.wikipedia.org/wiki/YAML) as well as [Golang](https://golang.org/) code containing structs and validators for data.

At this time, the code generation functionality does not exist, however there is a very functional Lisp interpreter in place!

The project consists, really, of two parts. **Unicorn** is the name given to the interpreter & code generation tool, and **Fig** is the name given to the Lisp dialect.

## How do I get started?

The first thing to do is of course to clone the repository and build the interpreter.

```bash
git clone https://github.com/redwire/UnicornFig.git
cd UnicornFig
sh build.sh
```

Now you can begin learning the language by checking out the introductory guide in `docs/guide.md` and/or by reading the example programs showing off the language's features in `examples/`.  All of the code there can be executed by unicorn.

You can run a Fig program by running the command

```bash
./unicorn <file>.fig
```
