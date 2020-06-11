# minimal libraray
```
<Project Sdk="Microsoft.NET.Sdk">
  <PropertyGroup>
    <TargetFramework>netstandard2.0</TargetFramework>
  </PropertyGroup>
</Project>
```

# minimal wpf
```
<Project Sdk="Microsoft.NET.Sdk.WindowsDesktop">
  <PropertyGroup>
    <UseWPF>true</UseWPF>
    
    <OutputType>WinExe</OutputType>
    <TargetFramework>netcoreapp3.0</TargetFramework>
  </PropertyGroup>
</Project>
```

# disable Generate AssemblyInfo
```
<PropertyGroup>
   <GenerateAssemblyInfo>false</GenerateAssemblyInfo>
</PropertyGroup>
```
