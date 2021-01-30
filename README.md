# Wallpape
Cross-platform (Linux, Windows, and macOS) app for searching, downloading and setting the desktop background.

![output](https://user-images.githubusercontent.com/49096838/106103657-7bb84180-616b-11eb-89eb-a61e0d17265e.gif)

## Usage
To build use the Makefile.
```
make build
```

Or use executables in [out folder](https://github.com/forChin/wallpape/tree/main/out).

Just start by running `wallpape` binary.
You can also specify query for your wallpaper by using `-q` flag.
```
./wallpape -q 'white dog'
```

## Photos
Photo is randomly selected and downloaded to the cache directory, which, according to your platform, will be:
- **Linux**: ~/.cache
- **Windows**: the first non-empty value from `%TMP%`, `%TEMP%`, `%USERPROFILE%`, or the Windows directory
- **macOS**: ~/Library/Caches

All photos provided by [Pexels](https://www.pexels.com).
