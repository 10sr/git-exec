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

With `-r <rev>` option, try to checkout that revision and run command at the
repository root:

    git exec --revision|-r <rev> <command> [<args> ...]

This command will fail if any staged or unstaged chagnes are found in current
repository.
In such a case, you can use `-w` option to checkout the revision to another
working directory:

    git exec --revision|-r <rev> -w <command> [<args> ...]

The working directory is `$XDG_CACHE_HOME/.git-exec/$name-$id`, where `$name` is the base
name of the directory and `$id` is calculated from the directory full path.

Run command with staged files:

    git exec --with-staged|-s <command> [<args> ...]
