# X - CLI Bookmark Manager

Manage your bookmarks seamlessly from the command line with **x**. This tool offers a range of commands to simplify bookmarking tasks, enhancing your workflow effortlessly.

## Usage

```sh
x [command]
```

## Available Commands

- **completion (c)**: Generate autocompletion scripts for your specified shell
- **delete (d)**: Remove a bookmark
- **export (e)**: Export the database to a JSON file
- **help (h)**: Get detailed information about any command
- **list (l)**: Display your saved bookmarks
- **open (o)**: Open a bookmark
- **save (s)**: Save a new bookmark
- **update (u)**: Update an existing bookmark

## Flags

```sh
# # Display help for the main command
x -h, --help

# For more detailed information about a specific command
x [command] --help
```

---

### List Command

Usage:

```sh
x list [flags]

```

Aliases:

```sh
list, l
```

Flags:

```sh
-h, --help # Display help for the list command
--type string # Filter bookmarks by type (default "all")
```

---

### Save Command

Usage:

```sh
x save [flags]
```

Aliases:

```sh
save, s
```

Flags:

```sh
-h, --help # Display help for the save command

```

---

### Delete Command

Usage:

```sh
x delete [flags]
```

Aliases:

```sh
delete, d
```

Flags:

```sh
-h, --help # Display help for the delete command
```

---

### Export Command

Usage:

```sh
x export [flags]
```

Flags:

```sh
-h, --help # Display help for the export command

```

---

### Open Command

Usage:

```sh
x open [flags]
```

Aliases:

```sh
open, o
```

Flags:

```sh
-h, --help: # Display help for the open command

```

---

### Update Command

Usage:

```sh
x update [flags]
```

Flags:

```sh
-h, --help #  Display help for the update command
--id int # ID of the bookmark to update
--url string # URL of the bookmark to update
```

Enhance your command-line bookmarking experience with **x**!