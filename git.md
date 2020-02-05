# commands
```
git config --global --list
git config --system --list
git config --local --list
git config --global --edit
git config --global alias.co checkout
git diff -w # shorthand for --ignore-all-space; no option for setting this as default in the config
```


# gitconfig aliases
```
[alias]
        co = checkout
        st = status
        fa = fetch --all
        mm = !git fetch origin master && git merge origin/master
```



# fix crlf (lineending) differences repo <-> workspace
```
PS: Get-ChildItem -Filter '*.cs' -Recurse | ForEach-Object { Remove-Item $_.FullName }
git reset --hard
```
