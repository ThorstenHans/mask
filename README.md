# mask - Hide your sensitive data from STDOUT

```bash
# Add new masks
mask add bar

# Mask an output

echo "foo bar baz" | mask
foo *** baz
```

For further information consult the help `mask --help` or `mask [command] --help`
