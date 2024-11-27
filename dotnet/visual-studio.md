# Share Required Components aka ".vsconfig"

- https://devblogs.microsoft.com/setup/configure-visual-studio-across-your-organization-with-vsconfig/

# Testexplorer clean cache
```del solutionfolder\.vs\solution name\v16\TestStore\<number>\*.testlog```
[Test Explorer (VS) shows '&lt;Unknown project&gt;'](https://stackoverflow.com/questions/59391984/test-explorer-vs-shows-unknown-project)

# Default Settings
- Environment > ProjectsAndSolution > ShowTaskListAfterBuild = false - "always show error list if build finishes with errors"

# Must-Have Extensions

* https://marketplace.visualstudio.com/items?itemName=NikolayBalakin.Outputenhancer

# Build Tools

```
# download installer
Invoke-WebRequest https://aka.ms/vs/17/release/vs_buildtools.exe -OutFile vs_buildtools.exe
Invoke-WebRequest https://aka.ms/vs/17/release/vs_enterprise.exe -OutFile vs_enterprise.exe

# install from cli
# https://learn.microsoft.com/en-us/visualstudio/install/use-command-line-parameters-to-install-visual-studio?view=vs-2022
winget install --id Microsoft.VisualStudio.2022.Community --override "--passive --config C:\my.vsconfig"

# add component to an allready installed VS
Start-Process "C:\Program Files (x86)\Microsoft Visual Studio\Installer\vs_installer.exe" -ArgumentList "modify", "--channelId", "VisualStudio.17.Release", "--productId", "Microsoft.VisualStudio.Product.BuildTools", "--add", "Microsoft.VisualStudio.ComponentGroup.VC.Tools.142.x86.x64", "--passive" -Wait
```
