# Desktop Bridge

- https://docs.microsoft.com/en-us/windows/apps/desktop/modernize/desktop-to-uwp-enhance
- Nugets:
  - https://www.nuget.org/packages/Microsoft.Windows.SDK.Contracts
  - https://www.nuget.org/packages/System.Runtime.InteropServices.WindowsRuntime/
  - https://github.com/qmatteoq/DesktopBridgeHelpers
- Share Mutex between UWP and Win32: https://stackoverflow.com/questions/46186350/share-named-mutex-between-uwp-and-desktop-app

# Known Locations

```
var userDataFolder = Windows.Storage.ApplicationData.Current.LocalFolder
var installedLocation = Windows.ApplicationModel.Package.Current.InstalledLocation
```

# Build Pipeline

- MakePri.exe 
  - compiles Resources to PRI (Package resource indexing) Files: 
  - https://docs.microsoft.com/en-us/windows/uwp/app-resources/pri-apis-custom-build-systems
  - https://docs.microsoft.com/en-us/windows/uwp/app-resources/build-resources-into-app-package
- MakeAppx.exe
  - pack appx
