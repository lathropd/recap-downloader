### Installation

- [Windows version](https://github.com/lathropd/recap-downloader/raw/main/bin/recap_downloader.exe) 
- [MacOS universal binary](https://github.com/lathropd/recap-downloader/raw/main/bin/recap_downloader_universal)


#### MacOS issue

Because this code is not developer signed for MacOS, your security permissions may not allow you to run the binaries on MacOS 
without administrator permissions.

One solution to this is to save the [MacOS universal binary](https://github.com/lathropd/recap-downloader/raw/main/bin/recap_downloader_universal) 
into your Downloads folder and run the following commands in the Terminal app (Applications->Utilities->Terminal).

```sh
cd ~/Downloads
xattr -dr com.apple.quarantine ./recap_downloader_universal
chmod +x ./recap_downloader_universal
```

Included for informational purposes only. I *do not* recommend, endorse, warrant, or otherwise express any support for this approach. 
