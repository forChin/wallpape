# Wallpape
Cross-platform (Linux, Windows, and macOS) app for searching and setting the desktop background.
![output](https://user-images.githubusercontent.com/49096838/106103657-7bb84180-616b-11eb-89eb-a61e0d17265e.gif)

## Usage
To build use the Makefile.
```
make build
```

Just start by running `wallpape` binary.
You can also specify key-words for your wallpaper by using `-q` flag.
```
./wallpape -q 'white dog'
```
Without specifying `-q` flag, default key-word will be "wallpaper".

## Photos
All photos downloads to cache directory, which, according to your platform, will be:
- **Linux**: \*HomeDir\*/.cache
- **Windows**: the first non-empty value from %TMP%, %TEMP%, %USERPROFILE%, or the Windows directory
- **MacOS**: \*HomeDir\*/Library/Caches

All photos provided by [Pexels](https://www.pexels.com).