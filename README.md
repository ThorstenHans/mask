# mask - Hide your sensitive data from STDOUT

With `mask` you can hide sensitive information from being printed to `stdout`. I've created `mask` because I do quite a lot of public conference speaking and found it cumbersome to memorize all CLIs that may leak sensitive information during a live demo.

You can pipe (`|`) any content to `mask`. It will replace all occurrence (case-insensitive). 

## Installation

Install `mask` using the latest [release for your platform](https://github.com/ThorstenHans/mask/releases). Once downloaded and extracted, ensure the location is in your `PATH`.

### Configuration

#### Manage masks

You can add any amount of strings that should be masked before they'll be written to `stdout`.

```bash
# Add new masks
mask add bar

# Remove existing masks
mask remove bar

# See all masks
mask list
```

#### Customize the mask-char

By default, `mask` will mask every char of an occurrence with an `*`. This ensures potential table layouts from the original output will not be harmed. You can manage the "mask-char" with the following commands:

```bash
# show the current mask-char
mask char get

# use different mask-char
mask char set "#"
```

## Mask output

```bash
# assumes 'bar' is in your list of masks
echo "foo bar baz" | mask
foo *** baz

# mask output from more complex commands
az account list -o table | mask

```

## Help

For further information consult the help `mask --help` or `mask [command] --help`. If you encounter any problems, feel free to file an issue.
