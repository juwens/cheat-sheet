| alias | url | syntax |
|--|--|--
| "Microsoft Symbol Servers" | http://msdl.microsoft.com/download/symbols | `SRV*http://msdl.microsoft.com/download/symbols` |
| "NuGet.org Symbol Server" | https://symbols.nuget.org/download/symbols | `SRV*https://symbols.nuget.org/download/symbols` |
| "http://symweb" | http://symweb.corp.microsoft.com (not available public) | `SRV*http://symweb.corp.microsoft.com` | 


- "Microsoft Symbol Servers" is located in `Microsoft.VisualStudio.Debugger.UserControls.SymbolPathControl.SymbolPathControlResources.resources/MicrosoftSymbolServersCaption`

  Links:

- https://github.com/terrajobst/experimentation/blob/master/Documentation/GettingStarted.md
- https://stackoverflow.com/questions/556383/how-to-use-windows-symbol-packages-with-visual-studio-2008
- https://github.com/deeprobin/performance/blob/b09098dbc05774e92213eb3d1862e47e5b382262/src/benchmarks/gc/src/analysis/clr.py#L274

```
C:\Program Files\Microsoft Visual Studio\2022\Enterprise\Common7\IDE\Profiles\General.vssettings:

...
		<PropertyValue name="NugetSymbolServerName">https://symbols.nuget.org/download/symbols</PropertyValue>
...
		<PropertyValue name="PublicSymbolServerName"/>
		<PropertyValue name="PublicSymbolServerName2">https://msdl.microsoft.com/download/symbols</PropertyValue>
...
```

```
$ grep -R "/download/symbols" /c/Program\ Files/Microsoft\ Visual\ Studio/2022
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/BlendExtensions/Profiles/Blend.vssettings:         <PropertyValue name="NugetSymbolServerName">https://symbols.nuget.org/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/BlendExtensions/Profiles/Blend.vssettings:         <PropertyValue name="PublicSymbolServerName2">https://msdl.microsoft.com/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/CSharp.vssettings:                <PropertyValue name="NugetSymbolServerName">https://symbols.nuget.org/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/CSharp.vssettings:                <PropertyValue name="PublicSymbolServerName2">https://msdl.microsoft.com/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/General.vssettings:               <PropertyValue name="NugetSymbolServerName">https://symbols.nuget.org/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/General.vssettings:               <PropertyValue name="PublicSymbolServerName2">https://msdl.microsoft.com/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/JavaScript.vssettings:            <PropertyValue name="NugetSymbolServerName">https://symbols.nuget.org/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/JavaScript.vssettings:            <PropertyValue name="PublicSymbolServerName2">https://msdl.microsoft.com/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/VB.vssettings:            <PropertyValue name="NugetSymbolServerName">https://symbols.nuget.org/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/VB.vssettings:            <PropertyValue name="PublicSymbolServerName2">https://msdl.microsoft.com/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/VC.vssettings:            <PropertyValue name="NugetSymbolServerName">https://symbols.nuget.org/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/VC.vssettings:            <PropertyValue name="PublicSymbolServerName2">https://msdl.microsoft.com/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/Web.vssettings:           <PropertyValue name="NugetSymbolServerName">https://symbols.nuget.org/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/Web.vssettings:           <PropertyValue name="PublicSymbolServerName2">https://msdl.microsoft.com/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/WebCode.vssettings:               <PropertyValue name="NugetSymbolServerName">https://symbols.nuget.org/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/Profiles/WebCode.vssettings:               <PropertyValue name="PublicSymbolServerName2">https://msdl.microsoft.com/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/VC/Profile/VC.vssettings:          <PropertyValue name="NugetSymbolServerName">https://symbols.nuget.org/download/symbols</PropertyValue>
/c/Program Files/Microsoft Visual Studio/2022/Enterprise/Common7/IDE/VC/Profile/VC.vssettings:          <PropertyValue name="PublicSymbolServerName2">https://msdl.microsoft.com/download/symbols</PropertyValue>
```

VS Logic
![image](https://github.com/juwens/cheat-sheet/assets/11560817/c5bdf831-f762-406d-a591-ad84b219d73c)
