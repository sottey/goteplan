# goteplan
Unofficial Noteplan Command Line Utility

This is an unofficial app to manage NotePlan notes from the command line. It should be able to compile for any operating system supporting the Go language.

>NOTE: The first time you run this, you should specify the argument -b [root of noteplan docs]. This will be saved to the config file. Absolute or relative pathing can be used.
EXAMPLE:
``` 
goteplan list -b "~/NotePLan/"
```

### Usage:
```
Usage:
  goteplan [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      Create a new note - Example: goteplan create Notes/Home/NewNote.md
  edit        Edit specified file - Example: goteplan edit Notes/Home/MyNote.md
  help        Help about any command
  list        List all notes - Example: goteplan list
  search      Search for notes with the specified string (NOTE: This IS case sensitive) - Example: goteplan search MikeS
  view        View specified note - Example: goteplan view Notes/Home/MyNote.md

Flags:
  -b, --basedir string   Root location of the NotePlan data
  -h, --help             help for goteplan

Use "goteplan [command] --help" for more information about a command.
```
