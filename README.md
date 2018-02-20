git-exec
========

Execute command in your git repository!


Install
-------

    go get -u github.com/10sr/git-exec


Usage
-----


Run command your git repository root:

    git exec <command> [<args>...]

Run command with specified git revision:

    git exec --revision|-r <rev> <command> [<args> ...]

Run command with staged files:

    git exec --with-staged|-s <command> [<args> ...]
