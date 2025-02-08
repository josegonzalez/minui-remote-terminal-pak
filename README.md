# minui-remote-terminal.pak

A MinUI app that starts a browser-based remote terminal.

## Requirements

This pak is designed and tested on the following MinUI Platforms and devices:

- `tg5040`: Trimui Brick (formerly `tg3040`), Trimui Smart Pro
- `rg35xxplus`: RG-35XX Plus, RG-34XX, RG-35XX H, RG-35XX SP

Use the correct platform for your device.

## Installation

1. Mount your MinUI SD card.
2. Download the latest release from Github. It will be named `Remote.Terminal.pak.zip`.
3. Copy the zip file to `/Tools/$PLATFORM/Remote Terminal.pak.zip`. Please ensure the new zip file name is `Remote Terminal.pak.zip`, without a dot (`.`) between the words `Remote` and `Terminal`.
4. Extract the zip in place, then delete the zip file.
5. Confirm that there is a `/Tools/$PLATFORM/Remote Terminal.pak/launch.sh` file on your SD card.
6. Unmount your SD Card and insert it into your MinUI device.

## Usage

> [!IMPORTANT]
> If the zip file was not extracted correctly, the pak may show up under `Tools > Remote`. Rename the folder to `Remote Terminal.pak` to fix this.

Browse to `Tools > Remote Terminal` and press `A` to enter the Pak.

### Debug Logging

To enable debug logging, create a file named `debug` in the pak folder. Logs will be written to the `$SDCARD_PATH/.userdata/$PLATFORM/logs/` folder.

### Configuration

#### port

The terminal runs on port `8080`. To utilize a different port, create a file named `port` in the pak folder with the port number you wish to run on.

#### shell

The terminal runs `/bin/sh` as the shell by default, but will detect the correct shell for the platform. To utilize a different shell, create a file named `shell` in the pak folder with the full path to the shell you wish to execute.
