<p align="center">
  <img src="docs/assets/logo.svg" alt="Declutter Logo" width="120" height="96">
</p>

<h1 align="center">Declutter</h1>

<p align="center">
  <strong>Organize your files into Year/Month folders based on their timestamps</strong>
</p>

<p align="center">
  <a href="https://github.com/dale-tomson/declutter/releases">
    <img src="https://img.shields.io/github/v/release/dale-tomson/declutter?style=flat-square" alt="Release">
  </a>
  <a href="https://github.com/dale-tomson/declutter/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/dale-tomson/declutter?style=flat-square" alt="License">
  </a>
  <a href="https://dale-tomson.github.io/declutter">
    <img src="https://img.shields.io/badge/website-live-42b883?style=flat-square" alt="Website">
  </a>
</p>

<p align="center">
  <a href="https://dale-tomson.github.io/declutter">Website</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#building-from-source">Build</a>
</p>

---

## Features

- 📁 **Auto Organization** — Creates year folders (e.g., `2024`) and month subfolders (e.g., `01-January`)
- 🕐 **Timestamp-based** — Uses file modification dates to determine the correct folder
- 🖥️ **Cross-platform** — Works on Windows, macOS, and Linux
- ⚡ **Fast** — Built with Go for blazing fast file operations
- 🎨 **Modern UI** — Clean interface built with Fyne

## How It Works

1. **Select** a folder containing files you want to organize
2. **Click** "Organize Files" to start
3. **Done!** Files are moved to Year/Month folders

### Before & After

```
Downloads/                          Downloads/
├── vacation.jpg (Mar 2024)         ├── 2023/
├── birthday.jpg (Mar 2024)         │   └── 12-December/
├── christmas.jpg (Dec 2023)   →    │       └── christmas.jpg
└── new-year.jpg (Jan 2024)         └── 2024/
                                        ├── 01-January/
                                        │   └── new-year.jpg
                                        └── 03-March/
                                            ├── vacation.jpg
                                            └── birthday.jpg
```

## Installation

### Download

Get the latest release for your platform from the [Releases page](https://github.com/dale-tomson/declutter/releases) or the [website](https://dale-tomson.github.io/declutter).

### Building from Source

#### Prerequisites

- Go 1.21+
- Platform dependencies:

**Linux:**
```bash
sudo apt-get install libgl1-mesa-dev xorg-dev libxrandr-dev pkg-config
```

**macOS:** Xcode command line tools

**Windows:** MinGW or similar C compiler

#### Build

```bash
git clone https://github.com/dale-tomson/declutter.git
cd declutter
go mod tidy
go build -o declutter .
```

#### Run

```bash
./declutter
```

## Usage

1. Launch Declutter
2. Click **Select Folder** and choose a folder with files to organize
3. Review the file count in the activity log
4. Click **Organize Files** and confirm
5. Watch the progress as files are moved to their Year/Month folders

## Testing

```bash
go test ./...
```

## Cross-Compilation

```bash
go install fyne.io/fyne/v2/cmd/fyne@latest

fyne package -name "Declutter"              # Current platform
fyne package -os windows -name "Declutter"  # Windows
fyne package -os darwin -name "Declutter"   # macOS
fyne package -os linux -name "Declutter"    # Linux
```

## License

MIT License — see [LICENSE](LICENSE) for details.

## Contributing

Contributions welcome! Feel free to open issues or submit pull requests.
