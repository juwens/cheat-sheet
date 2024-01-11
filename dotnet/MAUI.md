# Troubleshooting known issues

- https://docs.microsoft.com/en-us/dotnet/maui/troubleshooting

# workloads insights

- https://maui.blob.core.windows.net/metadata/rollbacks/<sdk-version>.json
- https://maui.blob.core.windows.net/metadata/rollbacks/7.0.101.json
- https://maui.blob.core.windows.net/metadata/rollbacks/main.json

# Handlers

- handler progress: https://github.com/dotnet/maui/projects/4
- handlers https://github.com/dotnet/maui/tree/main/src/Core/src/Handlers

# MAUI Sample

https://github.com/dotnet/net6-mobile-samples

# Select iOS Simulator

```
dotnet build -t:Run -f net6.0-ios /p:_DeviceName=:v2:udid=E3B7CC6A-4E29-4148-AACA-2A00B6CB3F52
```

# Version Matrix

| VS     | MAUI           | .Net SDK | Xamarin-Android | Xamarin-iOS | Date | Release Notes                                               |
|--------|----------------|----------|-----------------|-------------|--|-------------------------------------------------------------|
| 17.4.0 | 7.0.49 (GA)    |          | 33.0.4          | 16.0.1478   | 2022-11-08 | https://devblogs.microsoft.com/visualstudio/visual-studio-2022-17-4/<br> https://devblogs.microsoft.com/dotnet/dotnet-maui-dotnet-7/ |
| 17.4.1 |                |          |                 |             | 2022-11-15 | https://learn.microsoft.com/en-us/visualstudio/releases/2022/release-notes#1742--visual-studio-2022-version-1741 |
| 17.4.2 |                |          |                 |             | 2022-11-29 | https://learn.microsoft.com/en-us/visualstudio/releases/2022/release-notes#1742--visual-studio-2022-version-1742 |
| 17.4.3 | 7.0.52 (SR1.1) | 7.0.101  |                 |             | 2022-12-13 | https://learn.microsoft.com/en-us/visualstudio/releases/2022/release-notes#1743--visual-studio-2022-version-1743 |
| 17.4.4 |                | 7.0.102  |                 |             | 2023-01-10 | https://learn.microsoft.com/en-us/visualstudio/releases/2022/release-notes#1744--visual-studio-2022-version-1744 |
| 17.5.0 | 7.0.58         |          |                 |             |  | |

* https://developercommunity.visualstudio.com/VisualStudio?q=%5BVisual+Studio+2022+version+17.4%5D&sort=votes&page=1

# SDK Download

* https://aka.ms/dotnet/7.0.1xx/daily/dotnet-sdk-win-x64.zip
* https://dotnet.microsoft.com/en-us/download/dotnet/7.0
