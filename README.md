# Schmanguage CLI

A CLI tool for translating an input to schmanguage written in Go.

The program takes the text to translate as input and prints it as schmanguage to stdout.

## Usage

Install it by cloning and building the go source code:

```bash
git clone https://github.com/schmanguage/schmanguage-cli
cd schmanguage-cli
go build .
```

That creates a executeable called `schmanguage` which is used as the cli tool.

You can simply provide a literal text as input that will be translated:

```bash
./schmanguage Hello World
> Schmello Schmorld
```

### Flags

An optional flag `type` specifies the input type. Currently available types are `text` and `json`. If omitted it defaults to `text`.

So this gives the same result:

```bash
./schmanguage Hello World
> Schmello Schmorld

./schmanguage --type text Hello World
> Schmello Schmorld
```

When using the `json` type, the tool expexts a path to a json file containing string-string key-value pairs.

```json
// my_file.json
{
	"myText": "This is the BEST Language",
	"foo": "bar",
	"Hello": "World"
}
```

Assuming this file `my_file.json` and giving it to the schmanguage-cli:

```bash
./schmanguage --type json my_file.json
```

whould change the file to this
```json
// my_file.json
{
	"myText": "Schmis is the SCHMEST Schmanguage",
	"foo": "bar",
	"Hello": "Schmorld"
}
```

**Currently only string values are supported in a json file. If your file contains a non-string value, schmanguage-cli prints a warning to stdout and ignores this key-value pair and it remains unchanged.**
