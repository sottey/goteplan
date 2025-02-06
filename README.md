# goteplan
Unofficial Noteplan Command Line Utility

This is an unofficial app to manage NotePlan notes from the command line. It should be able to compile for any operating system supporting the Go language.

* NOTE: The first time you run this, you should specify the argument -b [root of noteplan docs]. This will be saved to the config file. Absolute or relative pathing can be used. *

### Usage:
```
Usage:
  goteplan [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      Create a new note
  edit        Edit specified file
  help        Help about any command
  list        List all notes
  search      Search for notes with the specified string (NOTE: This IS case sensitive)
  view        View specified note

Flags:
  -b, --basedir string   Root location of the NotePlan data
  -h, --help             help for goteplan

Use "goteplan [command] --help" for more information about a command.
```
