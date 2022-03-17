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
