# Nuget Archive Structure

```
/ref/{tfm}/my_refrerence.dll
/lib/{tfm}/refrerence_and_implementation.dll
/runtimes/{rid}/lib/{tfm}/my.dll
/runtimes/{rid}/native/my.dll
/contentFiles/{codeLanguage}/{TxM}/{any?}
```

- tfm | Target Framework Moniker: 
  - https://docs.microsoft.com/en-us/tfmd/otnet/standard/frameworks
  - https://docs.microsoft.com/en-us/nuget/reference/target-frameworks
- {rid} or {platform}-{architecture}
  - Runtime Identifier [runtime.json](https://github.com/dotnet/runtime/blob/main/src/libraries/Microsoft.NETCore.Platforms/src/runtime.json)
- /contentFiles/ [nuspec#including-content-files](https://docs.microsoft.com/en-us/nuget/reference/nuspec#including-content-files)
  - codeLanguages may be cs, vb, fs, any, or the lowercase equivalent of a given $(ProjectLanguage)
  - TxM is any legal target framework moniker that NuGet supports (see Target frameworks).
  - Any folder structure may be appended to the end of this syntax.

- https://docs.microsoft.com/en-us/nuget/create-packages/creating-a-package#from-a-convention-based-working-directory
- https://github.com/NuGet/Home/wiki/NuGet-Errors-and-Warnings

# System.Collections.Immutable problem

problem:

```
"Error	CS7038	Failed to emit module 'FooApp': Changing the version of an assembly reference is not allowed during debugging: 'System.Collections.Immutable, Version=1.2.2.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a' changed version to '1.2.5.0'"       
```

Nuget and dll AssemblyVersion differ

| Nuget | AssemblyVersion |
|=======|=================|
| 1.4.0 | 1.2.2.0         |
| 1.5.0 | 1.2.3.0         |
| 1.6.0 | 1.2.4.0         |
| 1.7.1 | 1.2.5.0         |
| 5.0.0 | 5.0.0.0         |
| 6.0.0 | 6.0.0.0         |
