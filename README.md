# Unicorn... Fig? What the heck?

**Unicorn Fig** is a really terrible play on the words "universal" and "configuration".

## Now you have my attention

I doubt that! But let's pretend I do for a moment, shall we?

UnicornFig has the ambitious goal of becoming a really simple "Emacs Lisp for everyone."  What that essentially means is that it hopes to not just be *yet another Lisp implementation*, but rather a Lisp interpreter that outputs configuration files in familiar formats like [JSON](https://en.wikipedia.org/wiki/JSON) and [YAML](https://en.wikipedia.org/wiki/YAML) as well as [Golang](https://golang.org/) code containing structs for data.

The project consists, really, of two parts. **Unicorn** is the name given to the interpreter & code generation tool, and **Fig** is the name given to the Lisp dialect.

## How do I get started?

The first thing to do is of course to clone the repository and build the interpreter.
Of course, to do this, you will need to have [git](https://www.git-scm.com/) and [Go](https://golang.org/dl/) installed on your computer.

```bash
git clone https://github.com/redwire/UnicornFig.git
cd UnicornFig
sh build.sh
```

Now you can begin learning the language by checking out the introductory guide in [docs/guide.md](https://github.com/redwire/UnicornFig/blob/master/docs/guide.md) and/or by reading the example programs showing off the language's features in `examples/`.  All of the code there can be executed by unicorn.

You can run a Fig program by running the command

```bash
./unicorn -json output.json -yaml config.yaml -go config.go <file>.fig
```

Of course, both the `-json`, `-yaml`, and `-go` arguments are optional.  If none are provided, Unicorn will execute the
program file provided and not write to any files.

You can see a practical example of how you might use Unicorn and Fig in the
[demo](https://github.com/redwire/UnicornFig/tree/master/demo) contained in the repository.
Further, you can see a full example of Uniorn's code generation in use in the
[demo/code-gen](https://github.com/zsck/UnicornFig/tree/master/demo/code-gen) directory.

It is advised that users read the `demo/code-gen/README.md` file to understand how the feature works and what its
shortcomings are.

## Vim syntax highlighting

For vim users, I've created an ftplugin, ftdetect, and syntax file for vim that provide reasonable syntax highlighting
for Fig code.  Everything is located in the `vim/` directory.  There is also a `vim/install.sh` script that you can run
to automatically copy the necessary files to the right place.

## Feedback? Questions? Suggestions?

I'd love to hear about them!

The best place to go is the [Github issue tracker](https://github.com/redwire/UnicornFig/issues) for the project.
