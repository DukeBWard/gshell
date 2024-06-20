# gshell
My attempt at making a shell with Golang

## Notes

* Hardest part so far is implementing `cd` and `pwd`.  You cannot just use built in Go stdlib functions if you want `gshell` to function like an actual shell.  This is because of how processes work, and the main thread that spawns this go program will always have its root as where it was invoked.  You must tackle this by keeping track of the directory you are actively modifying within the code.

* Tackling the `SIGINT` (CTRL + C) issue was a little interesting and I'm still playing around with it.  I would like to be able to interrupt the pritning buffer, but I have since moved on from the small issue.  It generally works as intended.

* TODO: Piping

* TODO: Scrollable history
