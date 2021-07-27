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

# Simple Clicked/Tapped Animation
```
private async void Tapped(object sender, EventArgs e)
{
    if (Command != null && Command.CanExecute(null))
    {
        await this.ScaleTo(0.85, length: 50, easing: Easing.SinOut);
        await this.ScaleTo(0.97, length: 50, easing: Easing.SinInOut);
        await Command.ExecuteAsync(null);
        Scale = 1;
    }
}
```

# Release History

| | |
|--|--|
|5.0|01 2021|
|4.0|04 2018|
|3.0|03 2018|
|2.0|11 2015|
|1.0|11 2014|
