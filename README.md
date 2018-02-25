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

with `-r <rev>` option, try to checkout that revision and run command at the
repository root:

    git exec --revision|-r <rev> <command> [<args> ...]

Run command with staged files:

    git exec --with-staged|-s <command> [<args> ...]
