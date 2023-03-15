# Timer
```
System.Windows.Forms.Timer: timer.Enabled = false;
System.Threading.Timer: timer.Change(Timeout.Infinite, Timeout.Infinite);
System.Timers.Timer: timer.Enabled = false; or timer.Stop();
```

# Gzip
```
using (var inFileStream = File.OpenRead(inFile))
using (var outFileStream = File.OpenWrite(outFile))
using (var compressionStream = new GZipStream(outFileStream, CompressionMode.Compress))
{
    inFileStream.CopyTo(compressionStream);
}
```

# CallerArgumentExpression
```
public static void Foo(
    bool bar,
    [CallerArgumentExpression(nameof(bar))] string barExpression = default)
{
}
```
https://blog.jetbrains.com/dotnet/2021/11/04/caller-argument-expressions-in-csharp-10/
