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

# deterministic builds
https://gist.github.com/aelij/b20271f4bd0ab1298e49068b388b54ae
```
  <PropertyGroup>
    <Deterministic>true</Deterministic>
  </PropertyGroup>
```

# IternalsVisibleTo without AssemblyInfo
```
  <ItemGroup>
    <AssemblyAttribute Include="System.Runtime.CompilerServices.InternalsVisibleTo">
      <_Parameter1>MyProject, PublicKey=0000000</_Parameter1>
    </AssemblyAttribute>
  </ItemGroup>
 ```
