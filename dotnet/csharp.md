# Timer
```
System.Windows.Forms.Timer: timer.Enabled = false;
System.Threading.Timer: timer.Change(Timeout.Infinite, Timeout.Infinite);
System.Timers.Timer: timer.Enabled = false; or timer.Stop();
```

# Gzip
```
using (var inFileStream = File.OpenRead(inFile))
using (var compressedFileStream = File.OpenWrite(outFile))
using (var compressedStream = new GZipStream(compressedFileStream, CompressionMode.Compress))
{
    inFileStream.CopyTo(compressedStream);
}
```
