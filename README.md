Pong
====
This is Pong written in GoLang using OpenGL and a custom engine

Dependencies
------------
- `Taskfile` - Executing 'build' scripts ([link](https://taskfile.dev/#/))
- `go-bindata` - Compiles assets into go code ([link](https://github.com/shuLhan/go-bindata))

Installation
------------
When pulling this repository please be sure to execute the `task init` command. Please be careful when committing
changes made to the submodule, as incorrect usage can lead to breaking it.

Available Tasks
---------------
- `task init` - Initializes the repository by pulling all required submodules.
- `task compileAssets` - In order to bundle assets with our binary we have to compile them into our binary
- `task build` - Builds program and creates an executable in `build/client[.exe]`
- `task run` - Runs the game
