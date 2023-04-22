# urlencode
Cli tool to encode stdin to URL.
This tool is a thin wrapper around the standard go escape functions `PathEscape` and `QueryEscape` from `net/url`.

## Examples

```console
$ ./urlencode.exe -h
urlencode (version: v0.7.0)

This program is a thin wrapper around the standard go url escape functions.
Available flags:

  -input string
    	string to escape, if empty (default) read from stdin
  -keep-spaces
    	keep spaces as they are
  -path-escape
    	use PathEscape instead of QueryEscape
  -trim
    	trim (from both sides) whitespaces and newlines
```

```console
$ urlencode -input " Ceci est un message étrange "
+Ceci+est+un+message+%C3%A9trange+
```

```console
$ echo " Ceci est un message étrange " | urlencode
+Ceci+est+un+message+%C3%A9trange+%0A
```

```console
$ echo " Ceci est un message étrange " | urlencode -keep-spaces
 Ceci est un message %C3%A9trange %0A
```

```console
$ echo " Ceci est un message étrange " | urlencode -path-escape
%20Ceci%20est%20un%20message%20%C3%A9trange%20%0A
```

```console
$ echo " Ceci est un message étrange " | urlencode -trim
Ceci+est+un+message+%C3%A9trange
```

Note : `echo` adds a new line at the end that is visible as `%0A` if no trim option is used.

## Some explanations

The url encoding is not the same before the `?` and after it.

- Before the `?` the spaces should be encodes as `%20` (and not `+`). This part is encoded with `PathEscape` go function.
- After the `?` the spaces could be encodes as `+` (`%20` is also possible but is longer ). This part is encoded with `QueryEscape` go function.

For mode info you can check [this answer](https://stackoverflow.com/a/29948396) at SX.

If you want to encode some text in `data:` uri, for example in the case of `data:image/xml+svg, ...` you can use `urlencode -path-escape -keep-spaces -trim`.
