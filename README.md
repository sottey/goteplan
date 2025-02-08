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
  goteplan [command]

Commands:
  create      Create a new note - Example: goteplan create Notes/Home/NewNote.md
  day         Display note from a specific date specified in the format YYYMMDD 
            (if no date provided, today's note is shown)
  delete      Delete specified note - Example: goteplan delete Notes/Home/MyNote.md
  edit        Edit specified note - Example: goteplan edit Notes/Home/MyNote.md
  list        List all notes - Example: goteplan list
  search      Search for notes with the specified string (NOTE: This IS case sensitive) - Example: goteplan search MikeS
  tasks       Display tasks from a specific date specified in the format YYYMMDD 
            (if no date provided, today's note is used)
  view        View specified note - Example: goteplan view Notes/Home/MyNote.md

Additional Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -b, --basedir string      Root location of the NotePlan data
  -h, --help                help for goteplan
  -r, --render              If present, display will attempt to render markdown. If not, source will be shown.
  -t, --todosymbol string   When using task command, a line starting with this symbol will be considered a task (default "*")

Use "goteplan [command] --help" for more information about a command.


### Config file (defaults to ~/.goteplan.json)
  "basedir": "Root/Location/Of/Noteplan/Data" - Specifies location Noteplan saves data
  "render": false - If true, raw note will be shown. If false, markdown will be (attempted to be) rendered
  "todosymbol": "*" - Symbol used for tasks (Noteplan allows *, -. Numbering not yet supported) 

```
