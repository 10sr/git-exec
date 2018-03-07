git-exec
========

Execute command in your git repository!


Install
-------

    go get -u github.com/10sr/git-exec


Usage
-----

Without any argument, just run given comand at the root directory of your git
repository (fail if you are not inside of any git repository):

    git exec <command> [<args>...]

With `-r <rev>` option, try to checkout that revision and run command.

    git exec --revision|-r <rev> <command> [<args> ...]

In this case, files are checked out into a working directory first (defaults to
`$HOME/.git-exec/$name.$id`, where `$name` is the base
name of the directory and `$id` is calculated from the directory full path)
and then the command will be invoked.

Also, you can run command with staged files:

    git exec --with-staged|-s <command> [<args> ...]


License
-------

This software is unlicensed. See [LICENSE](LICENSE) for detail.
