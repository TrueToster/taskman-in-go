Taskman
---------
**taskman** is a simple CLI task manager written in Go.
It stores tasks in JSON files (one task per file) and runs quickly and locally.

_About problems and bugs please write in Issues._
### Install
1. Install latest release
2. And then it depends on your OS
   - On Linux:
     1. Move the file ``taskman`` to /usr/local/bin/
     2. Restart your shell (`source ~/.bashrc` or `source ~/.zshrc`)
   - On Windows:
     1. Create dir: `mkdir $env:USERPROFILE\bin
`
     2. Move the file `taskman.exe` to this dir
     3. Add folder to PATH `setx PATH "$env:USERPROFILE\bin;$env:PATH"
`
     4. Reboot pc
### Usage
**add new task:**
`taskman` `addTask` `[title]` `[temp]` `[priority]` `[details]`
_Note: The priority argument must contain a value between 0 and 5._

**mark a task as completed:**
`taskman` `setTask` `[id]` `[done/undone]` 

**display all task information:**
`taskman` `view` `[id]`

**deleting a task:**
`taskman` `delTask` `[id]`

**show task list:**
`taskman` `list` `all`
_Note:Can be used without 'all' to show pnly unfinished tasks._

**show task list without color**
`taskman` `conky` `all`
_Note:Can be used without 'all' to show pnly unfinished tasks._