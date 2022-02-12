# `<Project Sdk="...">` 

- official:
  - `Microsoft.NET.Sdk` - for lib and console apps wich don't need a special SDK, also WPF and WinForms for .Net >= 5.0
  - `Microsoft.NET.Sdk.Web` - asp.net
  - `MSBuild.Sdk.Extras` - https://github.com/novotnyllc/MSBuildSdkExtras
- https://docs.microsoft.com/en-us/dotnet/core/project-sdk/overview

* order of import:
  * ".props"
  * your stuff
  * ".targets"
* Recursive Imports
  * Directory.Build.props
  * Directory.Build.targets
* Debugging
  * `dotnet build /pp` preprocess
  * `msbuild /pp:foo.targets MyProj.csproj`
    * works only with projects at the moment, not with solution: https://github.com/dotnet/msbuild/issues/2131
  * `msbuild solution.sln /bl` binlog
    * https://github.com/KirillOsenkov/MSBuildStructuredLog
    * VS: https://marketplace.visualstudio.com/items?itemName=VisualStudioProductTeam.ProjectSystemTools
  
# Custom Code inside csproj
https://stackoverflow.com/questions/12772428/how-do-i-specify-targetpath-in-startarguments-a-csvbfsproj-user/12776579#12776579
```
<TargetDir Condition="'$(OutDir)' != ''">$([MSBuild]::Escape($([System.IO.Path]::GetFullPath($([System.IO.Path]::Combine($(MSBuildProjectDirectory), $(OutDir)))))))</TargetDir>

<Target Name="PostBuild" AfterTargets="PostBuildEvent">
  <Exec Command="echo ProjectDir='$(ProjectDir)'&#xD;&#xA;echo SolutionDir='$(SolutionDir)'" />
</Target>
```
  

# Talks

 * [MS Build: Things You Should Know About Project Files - .NET Oxford - January 2020](https://www.youtube.com/watch?v=5HEbsyU5E1g)



