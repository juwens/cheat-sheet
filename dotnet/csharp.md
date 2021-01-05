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
