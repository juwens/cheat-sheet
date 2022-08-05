# fix: `You must install or update .NET to run this application`

Problem:
```
Testhost process exited with error: You must install or update .NET to run this application.
App: C:\dev\MyTestProject\bin\Debug\netcoreapp3.1\testhost.exe
Architecture: x64
Framework: 'Microsoft.WindowsDesktop.App', version '3.1.0' (x64)
.NET location: C:\Program Files\dotnet
The following frameworks were found:
  6.0.7 at [C:\Program Files\dotnet\shared\Microsoft.WindowsDesktop.App]
Learn about framework resolution:
https://aka.ms/dotnet/app-launch-failed
To install missing framework, download:
https://aka.ms/dotnet-core-applaunch?framework=Microsoft.WindowsDesktop.App&framework_version=3.1.0&arch=x64&rid=win10-x64
. Please check the diagnostic logs for more information.

Test Run Aborted.
```

Solution: 

```
dotnet test --environment DOTNET_ROLL_FORWARD=LatestMajor
```
