```
git config --global --list
git config --system --list
git config --local --list
```

gitconfig
```
alias.co=checkout
alias.st=status
alias.fa=fetch --all

# (m)erge (m)aster
alias.mm=!git fetch origin master && git merge origin/master
```

fix crlf (lineending) differences repo <-> workspace
```
PS: Get-ChildItem -Filter '*.cs' -Recurse | ForEach-Object { Remove-Item $_.FullName }
git reset --hard
```
