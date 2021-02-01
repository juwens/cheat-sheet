# Unity "Raw Format" 

is what in Photoshop "Save as > Photoshop Raw" produces.
No further specification or alternatives 

# What didn't work

- XNView: couldn't display raw files properly
- FastRawViewer: refused to open the files
- imagemagick (which uses ucraw-batch) fails
  - `identify heightmap_Europe_9000x6705.raw`
  - `identify-im6.q16: delegate failed `'ufraw-batch' --silent --create-id=also --out-type=png --out-depth=16 '--output=%u.png' '%i'' @ error/delegate.c/InvokeDelegate/1928.`
  - `identify-im6.q16: unable to open image `/tmp/magick-22411IpelM2sHwzF9.ppm': No such file or directory @ error/blob.c/OpenBlob/2874.`
- ucraw-batch (which uses dcraw) fails
- dcraw fails
  - `dcraw -i -v heightmap_Europe_9000x6705.raw`
  - `Cannot decode file heightmap_Europe_9000x6705.raw`
