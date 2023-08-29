* https://learn.microsoft.com/en-us/windows/win32/shell/links
* https://learn.microsoft.com/en-us/windows/win32/properties/property-system-overview
* https://syedhasan010.medium.com/forensics-analysis-of-an-lnk-file-da68a98b8415
* https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-shllink/16cb4ca1-9339-4d0c-a68d-bf1d6cc0f943
* https://learn.microsoft.com/en-us/troubleshoot/windows-client/admin-development/create-desktop-shortcut-with-wsh
* https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishelllinkw
* https://github.com/microsoft/win32metadata/blob/main/generation/WinSDK/RecompiledIdlHeaders/um/propkey.h
* https://github.com/microsoft/Windows-classic-samples/blob/main/Samples/DesktopToasts/CS/ShellPropertyKeys.cs
* https://github.com/winsiderss/systeminformer/blob/master/phlib/include/appresolverp.h

IShellLink was never extended for UWP/MSIX/WinRT, you need to cast it to a property-store an then set specific properties. 

```
		auto link {winrt::create_instance<IShellLink>(CLSID_ShellLink)};
		winrt::check_hresult(link->SetPath(module_path.c_str()));

		auto store {link.as<IPropertyStore>()};
		prop_variant value;
		winrt::check_hresult(InitPropVariantFromString(app_user_model_id.c_str(), &value));
		winrt::check_hresult(store->SetValue(PKEY_AppUserModel_ID, value));
```
