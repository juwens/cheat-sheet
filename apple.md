* https://developer.apple.com/download/all/

```
[xml]$infoPlist = Get-Content ./src/App/App.iOS/Info.plist
Select-Xml -Xml $infoPlist -XPath "/plist/dict/key[text()='CFBundleDisplayName']/following-sibling::string[1]" | Select-Object -ExpandProperty Node

(Select-Xml -Xml $infoPlist -XPath "/plist/dict/key[text()='CFBundleDisplayName']/following-sibling::string[1]").node.InnerXML = "App.dev"
(Select-Xml -Xml $infoPlist -XPath "/plist/dict/key[text()='CFBundleIdentifier']/following-sibling::string[1]").node.InnerXML = "de.app.identifier.dev"
(Select-Xml -Xml $infoPlist -XPath "/plist/dict/key[text()='CFBundleName']/following-sibling::string[1]").node.InnerXML = "de.app.dev"
```
