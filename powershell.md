# boilerplate

```
Set-StrictMode -Version Latest
```

# profile
```
if (!(Test-Path -Path $PROFILE)) {New-Item -ItemType File -Path $PROFILE -Force}
notepad $PROFILE
```
* https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_profiles?view=powershell-7#the-profile-files
