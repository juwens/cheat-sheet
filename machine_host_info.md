```
dotnet --info
echo $PSVersionTable

$vswhere = 'C:\Program Files (x86)\Microsoft Visual Studio\Installer\vswhere.exe'
& $vswhere
[xml]$doc = & ".\tools\vswhere.exe" @("-utf8", "-include", "packages", "-format", "xml", "-property", "packages")
$doc.GetElementsByTagName("package") | Sort-Object id | Format-Table
```
