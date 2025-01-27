# trimui-brick-remote-terminal.pak

A TrimUI Brick app that starts a browser-based remote terminal.

## Requirements

- Docker (for building)

## Building

```shell
make release
```

## Installation

1. Mount your TrimUI Brick SD card.
2. Download the latest release from Github. It will be named `Remote.Terminal.pak.zip`.
3. Copy the zip file to `/Tools/tg5040/Remote Terminal.pak.zip`. Please ensure the new zip file name is `Remote Terminal.pak.zip`, without a dot (`.`) between the words `Remote` and `Terminal`.
4. Extract the zip in place, then delete the zip file.
5. Confirm that there is a `/Tools/tg5040/Remote Terminal.pak/launch.sh` file on your SD card.
6. Unmount your SD Card and insert it into your TrimUI Brick.

> [!NOTE]
> The device directory changed from `/Tools/tg3040` to `/Tools/tg5040` in `MinUI-20250126-0` - released 2025-01-26. If you are using an older version of MinUI, use `/Tools/tg3040` instead.

## Usage

> [!IMPORTANT]
> If the zip file was not extracted correctly, the pak may show up under `Tools > Remote`. Rename the folder to `Remote Terminal.pak` to fix this.

Browse to `Tools > Remote Terminal` and press `A` to turn on the remote terminal. It is accessible on port 8080 of the device IP.

### daemon-mode

By default, `remote-term` runs as a foreground process, terminating on app exit. To run `remote-term` in daemon mode, create a file named `daemon-mode` in the pak folder. This will turn the app into a toggle for `remote-term`.

### port

The terminal runs on port 8080. To utilize a different port, create a file named `port` in the pak folder with the port number you wish to run on.

### shell

The terminal runs `/bin/sh` as the shell by default. To utilize a different shell, create a file named `shell` in the pak folder with the full path to the shell you wish to execute.
