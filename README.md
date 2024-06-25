# gshell
My attempt at making a shell with Golang

## Commands
The goal was to behave like a normal shell program.  I have tried my hardest to support any external applications that you already have installed on your PC while also building some of my own.  

* `pwd`: shows your current working directory
* `ls`: lists files in your current directory
* `cd <directory>`: allows the user to change directory
* `history`: shows previous command history

* External commands are ran just like how you would expect.  ex: `cat <filename>` works as intended and `git add <files>` etc.

## Notes

* Hardest part so far is implementing `cd` and `pwd`.  You cannot just use built in Go stdlib functions if you want `gshell` to function like an actual shell.  This is because of how processes work, and the main thread that spawns this go program will always have its root as where it was invoked.  You must tackle this by keeping track of the directory you are actively modifying within the code.

* Tackling the `SIGINT` (CTRL + C) issue was a little interesting and I'm still playing around with it.  I would like to be able to interrupt the pritning buffer, but I have since moved on from the small issue.  It generally works as intended.

* TODO: Piping