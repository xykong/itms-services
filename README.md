# itms-services

A command-line tool for iOS over-the-air (OTA) app distribution via Apple's
`itms-services://` protocol.

Given a `.ipa` file and a public HTTPS base URL, `itms-services` generates a
`manifest.plist` and an `index.html` install page (with QR code), and
optionally starts an HTTP file server — all in one command.

## Installation

### Homebrew (macOS / Linux)

```bash
# Coming soon — track https://github.com/xykong/itms-services/releases
```

### Go install

```bash
go install github.com/xykong/itms-services@latest
```

### Pre-built binaries

Download the latest binary for your platform from the
[Releases](https://github.com/xykong/itms-services/releases) page.

## Quick start

### 1. Generate files only (`build`)

```bash
itms-services build MyApp.ipa \
  --host-url        https://example.com/install \
  --display-image   https://example.com/assets/icon.png \
  --output          ./output
```

Produces `./output/manifest.plist` and `./output/index.html`.  
Point your web server at `./output/` and send users the `index.html` URL.

### 2. Build and serve in one step (`serve`)

```bash
itms-services serve MyApp.ipa \
  --host-url      https://example.com \
  --display-image https://example.com/icon.png \
  --port          8080
```

This copies `MyApp.ipa` into `./output/`, generates the manifest and page,
then starts an HTTP server on `:8080`.

> **Note:** Apple requires `itms-services://` to fetch the manifest over
> **HTTPS**. For local testing use [ngrok](https://ngrok.com/) or
> [Caddy](https://caddyserver.com/) as a TLS reverse proxy in front of the
> built-in server.

## How it works

```
itms-services build MyApp.ipa …
        │
        ├─ reads Info.plist from MyApp.ipa (bundle ID, version, display name)
        ├─ renders manifest.plist  ──► served at <host-url>/manifest.plist
        └─ renders index.html      ──► served at <host-url>/index.html
                  │
                  └─ install button href:
                     itms-services://?action=download-manifest&url=<host-url>/manifest.plist
```

## Flags reference

### `build` / `serve` (shared flags)

| Flag | Default | Description |
|---|---|---|
| `--host-url` | *(required)* | Public base URL, e.g. `https://example.com` |
| `--display-image` | *(required)* | 57×57 icon URL shown on the install page |
| `--output`, `-o` | `./output` | Directory for generated files |
| `--title`, `-t` | `Test app` | App name shown in the install dialog |
| `--bundle-identifier` | | `CFBundleIdentifier` (auto-read from ipa) |
| `--bundle-version` | `1.0` | `CFBundleVersion` (auto-read from ipa) |
| `--build-number` | `0` | Build number badge on the install page |
| `--manifest-name` | `manifest.plist` | Output filename for the manifest |
| `--manifest-software-package` | `<host-url>/<ipa-name>` | Direct URL to the `.ipa` file |
| `--manifest-full-size-image` | *(display-image)* | 512×512 full-size image URL |
| `--manifest-subtitle` | | Subtitle in the install alert |
| `--index-name` | `index.html` | Output filename for the install page |
| `--index-platform` | `iOS` | Platform badge |
| `--index-branch` | `trunk` | Branch badge |
| `--index-jumbotron-description` | *(see help)* | Description text on the page |
| `--index-install-button-text` | `Install` | Install button label |

### `serve`-only flags

| Flag | Default | Description |
|---|---|---|
| `--port`, `-p` | `8080` | TCP port to listen on |

## License

Apache 2.0 — see [LICENSE](LICENSE).
