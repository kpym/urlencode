# urlencode
Cli tool to encode stdin to URL.
This tool is a thin wrapper around the standard go escape functions `PathEscape` and `QueryEscape` from `net/url`.

## Examples

```bash
> ./urlencode.exe -h
urlencode (version: v0.4.0)

This program is a thin wrapper around the standard go url escape functions.
Available flgs:

  -keep-spaces
        keep spaces as they are
  -path-escape
        use PathEscape in place of QueryEscape
  -trim
        trim (from both sides) spaces and new lines
```

```bash
echo " Ceci est un message étrage " | urlencode
+Ceci+est+un+message+%C3%A9trage+%0A
```

```bash
echo " Ceci est un message étrage " | urlencode -keep-spaces
 Ceci est un message %C3%A9trage %0A
```

```bash
echo " Ceci est un message étrage " | urlencode -path-escape
%20Ceci%20est%20un%20message%20%C3%A9trage%20%0A
```

```bash
echo " Ceci est un message étrage " | urlencode -trim
Ceci+est+un+message+%C3%A9trage
```

Note : `echo` adds a new line at the end that is visible as `%0A` if no trim option is used.
