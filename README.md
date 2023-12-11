# To

This is a simple bookmark tools to help me remember dir and access them with
shortcut.

Usage:

```sh
to save foo     # save current dir as foo

to delete foo   # delete foo bookmark

to list         # list all saved dirs
to list -c      # list all saved dirs under current dir
to list -f foo  # list all saved dirs with foo prefix

to find foo     # find the bookmarked dir keyword match to foo

j foo           # cd to foo matched bookmarked dir
```

Need to use shell's function to actually cd to the dir.

## Matching Algorithm

1. find if an exact match. eg. if "foo", "foobar" is saved, `to find foo` will
   match "foo"
2. find the shortest match with given word as prefix.  eg. if "foo", "foobar"
   is saved, `to find f` will match "foo"

## Generate Completion

```sh
to completion fish > ~/.config/fish/completions/to.fish

# or other shell
```

## Installation

For fish:

```sh
scripts/fish/install.sh
```
