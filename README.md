# ugo-git

uGo-Git is a small implementation of a Git-like version control version (VCS) to studies Git and Golang.


## Commands

```bash
$ cd /tmp/new
$ ugit init
Initialized empty ugit repository in /tmp/new/.ugit
$ echo some file > bla
$ ugit hash-object bla
0e08b5e8c10abc3e455b75286ba4a1fbd56e18a5
$ ugit cat-file 0e08b5e8c10abc3e455b75286ba4a1fbd56e18a5
some file

ugit write-tree
13bfa525aff7d5578af5ffe952b146c2df3db1e9	./.gitignore
543acf7ae3e1bd4e51ecdd6eff9ce8ec6d83ff9e	./LICENSE
298b5dcc9fb071c7150a6208586f24d2f4dbc05d	./README.md
acbf441966c502bb5d54efe885877e0e970e85b9	./bla
63a86e68b95f4d29b20e54f7170f0a5fd6d9a4fc	./cli.go
fcf714a2cf0e3bbcccac4988ca6e43adfea42f50	./data/base.go
b0afaf3a3c1e6565f4eee33df43acfce99b05f9a	./data/data.go
4dd32ad6633d8630eb05fd7e708a704d705f6fc3	./data/objects.go
4e7b7a4a2ab1dad7cb0d177c26ff558e494cee7c	./go.mod
23e75ce6f602eb1cf3df080f9dcfaa8f419b5d28	./go.sum

```
