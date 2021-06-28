Glossary:

- TPM: Trusted Platform Module. A Hardware in a modern Windows Computer, which can be used/utilized for security:
  - directly by the OS (Windows and Linux)
  - by other Hardware, as Intel TXT/SMX 
- dTPM: discrete TPM 1.0/1.2 or 2.0; a module you (Plug) or the OEM (Plug or solder) needs to add to your main board, usually proprietary modules via a 13, 15, 17, 19 Pin connector (pre 2015)
- fTPM: Firmware TPM; is always TPM 2.0 compatible. A module embedded into the CPU or Chipset. (from 2015 to this day) No need for a dTPM on the mainboard anymore. But dTPM can be used too.
- Intel TXT/SMX: a Intel CPU Extension which **utilizes** a separate dTPM or fTPM; TXT/SMX does **not** contain an fTPM or dTPM
- Intel PTT: intels Hardware implementation of fTPM; embedded/integrated in the Chipset since LGA 1151 (anno 2015)
- AMD PSP: Platform Security Processor, AMDs umbrella Term for any of the Intel equivalents to TXT/SMX, PTT, fTPM, Intel ME
  - fTPM is  build into the CPU/SOC/APU since 2018
  - side note: AMDs fTPM is acutally a [ARM Cortex-A5 TrustZone][1]
- TPM 2.0: (usually) includes support for 1.0 and 1.2
- TPM 1.0/1.2: old TPM spec. (Pre 2013)

There are three options:

| General | Intel| AMD |
|--|--|--|
| mostly no hw-support for a dTPM<br> not able to run Windows 11|pre 2013|
| A lot of the Mainboards have a proprietary socket for a dTPM<br><br>A dTPM (or alternatively fTPM) is required! <br> The TPM itself is not included, but **may** be present. Plugged/soldered on the MB by OEM or the user <br>If you have a dTPM you should be able to enable it in BIOS/UEFI and run Windows 11. | since 2013 <br><br> separate dTPM required <br><br> None of the **LGA1150** (aka. 4th Gen Core CPU aka Haswell) Chipsets (H81, C222, B85, C224, Q85, Q87, C226,H87, Z87) and prior have PTT support, hence no embedded/integrated fTPM|since 2016/2017 <br><br> - separate dTPM required <br>- no embedded fTPM <br><br>this concerns Zen (1fst gen)/Ryzen 1000 Mainboards|
| official windows 11 support. <br>a fTPM module embedded/integrated into the CPU or Chipset | since 2015<br><br> fTPM (which Intel calls PTT) included in every Chipset (except C236) <br><br>All **LGA1200** (8/9/10/11th gen core CPU) with 400 and 500 chipset have PTT support <br> <br> All (except one) of the **LGA1151** (6/7th gen Core cpu) Chipsets (100, 200, 300, **except C236**) have PTT support | since 2018 <br><br> fTPM included in every SOC/CPU<br><br>Socket AM4, "Zen+" (Ryzen 2000), "Zen 2", "Zen 3" and newer contain an embedded fTPM |


  [1]: https://en.wikipedia.org/wiki/ARM_architecture#Security_extensions
