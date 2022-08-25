# commands
```
git fetch -pn
git remote prune origin -n # remove branches not on remote

# config
git config --global --list
git config --system --list
git config --local --list
git config --global --edit
git config --global alias.co checkout

# diff
git diff -w # shorthand for --ignore-all-space; no option for setting this as default in the config
git diff -b # shorthand for --ignore-space-change
```


# gitconfig aliases
```
git config --global alias.co checkout
git config --global alias.st status
git config --global alias.sw switch
git config --global alias.fa 'fetch --all --prune'
git config --global alias.moma '!git fetch origin main && git merge origin/main'
```

```
[alias]
        co = checkout
        st = status
        sw = switch
        fa = fetch --all --prune
        moma = !git fetch origin main && git merge origin/main
```



# fix crlf (lineending) differences repo <-> workspace
```
Get-ChildItem -Filter '*.cs' -Recurse | ForEach-Object { Remove-Item $_.FullName }
git status --porcelain|% {$_.substring(3)}|% { Remove-Item $_ }
git reset --hard
```
