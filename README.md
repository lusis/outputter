# outputter
Outputter is a small go library for formatting data you want to present to stdout
The feel is very much modeled after [tablewriter](github.com/olekukonko/tablewriter)

## usage

In the `example` directory, there are examples for `flag`, `kingpin` and `cobra` usages

## color support

If you like to use the awesome [color](https://github.com/fatih/color) package, you can but it's up to each output formatter to decide if it supports colorized output or not. The best way to handle this is like so:

```go
color.NoColor = !outputFormatter.ColorSupport()
outputFormatter.SetHeaders([]string{
    "header1",
    "header2",
    "header3",
})
rowErr := outputFormatter.AddRow([]string{
    color.YellowString("value1"),
    color.YellowString("value2"),
    color.YellowString("value3"),
})
```

It's important to trigger the flag BEFORE you start setting color codes.

There's an example in `examples/color`

## pretty output

If an output format supports a pretty print mechanism of some kind, you can call `outputFormatter.SetPretty()` to trigger it.

Currently only the `json` output supports that.

## note about json output

Depending on your shell configuration and prompts, a single line json may get "erased" when outputting. You see this frequently with complicated zsh prompts. You can be sure your output is working by piping through `jq`.

## extending
