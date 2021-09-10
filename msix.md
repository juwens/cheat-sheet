# msix

Combines setup.exe, msi, ClickOnce, APPX (Store/UWP Apps)

- https://microsoft.github.io/MSIX-Labs/
- https://docs.microsoft.com/en-us/windows/msix/resources
- https://medium.com/@amansrivastava_11501/msix-the-future-of-universal-app-format-and-application-management-in-depth-part-1-f5e28c78819c
- https://www.advancedinstaller.com/msix-introduction.html

# special case: framework dependent msix

https://techcommunity.microsoft.com/t5/windows-dev-appconsult/packaging-a-net-core-3-0-application-with-msix/bc-p/2076112/highlight/true#M496

# Tools

- https://msixhero.net/get/

# misc

- Glossar
  - Package Family Name: "MSIXHero_zxq1da1qqbeze"
  - AppUserModelId: "MSIXHero_zxq1da1qqbeze!App"
  - Full package name: "MSIXHero_2.2.8.0_neutral__zxq1da1qqbeze"

- launch packaged app from Win32
  - register app specific uri-schema. `my-fancy-app:`
  - execution alias
    - folder `%userprofile%\AppData\Local\Microsoft\WindowsApps\`
  - `shell:appsFolder\Microsoft.People_8wekyb3d8bbwe!App`
  - `"C:\Program Files (x86)\Windows Kits\10\App Certification Kit\microsoft.windows.softwarelogo.appxlauncher.exe" MyPackageName_ph1m9x8skttmg!AppId`
