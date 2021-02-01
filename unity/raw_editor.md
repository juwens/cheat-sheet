# Unity "Raw Format" 

- is what in Photoshop "Save as > Photoshop Raw" produces.
- No further specification or alternatives (as for example png import)

- Photoshop Raw format
  - https://helpx.adobe.com/photoshop/using/file-formats.html#photoshop_raw_format
  - `contain 8-bits-per-channel RGB data in top-to-bottom, left-to-right pixel order. Dimensions must be input manually when such files are re-opened, or a square image is assumed.` (Source: [wikipedia](https://en.wikipedia.org/wiki/Raw_image_format#Raw_bitmap_files))

# Editor/Viewer for "Unity/Photoshop Raw"

# The Solution

```
magic stream -map r -storage-type short image.png image_1081x1081_16bit.praw
magic stream -map r -storage-type char image.png image_1081x1081_8bit.praw
```

- open with GIMP:
  - File > Open ... > select file (don't click open yet)
  - expand "Select File Type"
    - select "Raw Image Data" (Extension 'data')
  - Click Open
  - Wizard "Load Image from Raw Data" will open
    - set correct! image resolution
    - Image Type
      - "Gray 8 bit" for 8 bit 
      - "Gray unsigned 16 bit Little Endian" for 16 bit (unity identifies this as "windows byte order")

# What did work

- IrfanView (opening and viewing raw files worked; but was not able to import files saved with Irfanview)
- GIMP (opening and viewing raw files dit **not** work; but GIMP is able to export as raw files, which can opened by unitiy "import raw")

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
