# Tutorials

https://docs.microsoft.com/en-us/learn/browse/?term=xamarin&terms=xamarin

# Analyse Binding Errors

```
#if DEBUG
            Xamarin.Forms.Internals.Log.Listeners.Add(new Xamarin.Forms.Internals.DelegateLogListener((arg1, arg2) =>
            {
                System.Diagnostics.Debug.WriteLine(arg2);
            }));
#endif
```
