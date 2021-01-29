# boilerplate

```
Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'
trap {break}
```

# profile
```
if (!(Test-Path -Path $PROFILE)) {New-Item -ItemType File -Path $PROFILE -Force}
notepad $PROFILE
```
* https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_profiles?view=powershell-7#the-profile-files

# list disks
```
get-disk | Select Number, FriendlyName, BusType, PartitionStyle, Manufacturer, UniqueId, UniqueIdFormat, Signature, LogicalSectorSize, PhysicalSectorSize | Format-Table
```
