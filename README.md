Pong
====
This is Pong written in GoLang using OpenGL and a custom engine

Dependencies
------------
- `Taskfile` - Executing 'build' scripts ([link](https://taskfile.dev/#/))
- `go-bindata` - Compiles assets into go code ([link](https://github.com/shuLhan/go-bindata))

Installation
------------
When pulling this repository please be sure to either include the `--recurse-submodules`
argument or execute the `git submodule init` command. To update the submodule make sure
to use the `git submodule update` command, it will pull the latest version of the submodule
that is available.

`Taskfile` is used to automate annoying tasks. Whenever you change your assets
you will need to recompile them, to do this you can execute `task compileAssets`
this will automatically compile the `assets/` directory into the `generated/assets/`
directory.