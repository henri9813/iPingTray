<h1>
  <img src="/docs/icon.png" alt="Project Logo" width="50" style="vertical-align:middle;">
  IPingTray
</h1>

This program permit to have your current ping in `ms` inside your menu bar on your computer !
( cf [Screenshots](#screenshots))

This tool is useful when traveling with an unstable connection to know if you can try to work or not 🙂.

## Installation

You can download the last release of the project in the
[project releases](https://github.com/henri9813/iPingTray/releases).

## Usage

Running the app will automatically add you a tray in your menu bar.

The reference to determine your ping is the Google DNS server (`8.8.8.8`),
this option cannot be configured at this time)

### Thresholds

| Indicator |  Latency Range  |
|:---------:|:---------------:|
|     🟢     | Latency < 40ms  |
|     🟠     | Latency < 75ms  |
|     🔴     | Latency >= 75ms |

## Start app automatically on boot

###  macOS

You can follow the article from Apple which explain how to start the app automatically at boot:

<https://support.apple.com/guide/mac-help/open-items-automatically-when-you-log-in-mh15189>

### Windows

The installer automatically add the program to startup programs.

## Screenshots

![macOS screenshot](/docs/screenshot-macos.png)

![windows screenshot](/docs/screenshot-windows.png)

## Development

### Generate icons

If you wish to update the icons, you need to run the `generate-icons.sh` which gonna generate
golang []bytes array representing the icon.

To generate icons, you also need to install [`2goarray`](https://github.com/cratonica/2goarray)

```bash
go get github.com/cratonica/2goarray
go install github.com/cratonica/2goarray
```

> ⚠️ We use .ico file because it's the only supported format on windows, according to
> [@ZGGSONG](https://github.com/getlantern/systray/issues/154#issuecomment-1207607136)


## Contribution

Feel free to contribute !

## Licence

Creative Commons Attribution-NonCommercial 4.0 International License