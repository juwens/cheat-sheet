# Timer
```
System.Windows.Forms.Timer: timer.Enabled = false;
System.Threading.Timer: timer.Change(Timeout.Infinite, Timeout.Infinite);
System.Timers.Timer: timer.Enabled = false; or timer.Stop();
```

# Gzip
```
using (FileStream compressedFileStream = File.Create(fileToCompress.FullName + ".gz"))
{
    using (FileStream originalFileStream = fileToCompress.OpenRead());
    using (GZipStream compressionStream = new GZipStream(compressedFileStream, CompressionMode.Compress));
    
    originalFileStream.CopyTo(compressionStream);
}
```
