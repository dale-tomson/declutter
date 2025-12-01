# Declutter

A cross-platform standalone application built with [Fyne](https://fyne.io/) (Go) that organizes files into Year/Month folders based on their timestamps.

## Features

- 📁 **Automatic folder creation**: Creates year folders (e.g., `2024`, `2025`) and month subfolders (e.g., `01-January`, `02-February`)
- 🕐 **Timestamp-based organization**: Uses file modification dates to determine the correct folder
- ⚠️ **No duplicate folders**: Ensures folders are only created once
- 🔄 **Skip existing files**: Files that already exist at the destination are skipped
- 📝 **Activity log**: Real-time logging of all operations
- 🖥️ **Cross-platform**: Works on Windows, macOS, and Linux

## How It Works

1. Select a folder containing files you want to organize
2. Click "Organize Files" to start the process
3. The app will:
   - Scan all files in the selected folder
   - Create year folders based on file modification dates
   - Create month subfolders inside year folders
   - Move each file to its corresponding month folder

### Folder Structure Example

Before:
```
my-photos/
├── vacation.jpg (modified: 2024-03-15)
├── birthday.jpg (modified: 2024-03-22)
├── christmas.jpg (modified: 2023-12-25)
└── new-year.jpg (modified: 2024-01-01)
```

After:
```
my-photos/
├── 2023/
│   └── 12-December/
│       └── christmas.jpg
└── 2024/
    ├── 01-January/
    │   └── new-year.jpg
    └── 03-March/
        ├── vacation.jpg
        └── birthday.jpg
```

## Building from Source

### Prerequisites

- Go 1.21 or later
- Fyne dependencies (platform-specific)

#### Linux
```bash
sudo apt-get install libgl1-mesa-dev xorg-dev libxrandr-dev pkg-config
```

#### macOS
Xcode command line tools are required.

#### Windows
A C compiler (like MinGW) is required.

### Build

```bash
go mod tidy
go build -o declutter .
```

### Run

```bash
./declutter
```

### Test

```bash
go test ./...
```

## Cross-Compilation

Using Fyne's packaging tools:

```bash
# Install fyne CLI
go install fyne.io/fyne/v2/cmd/fyne@latest

# Package for current platform
fyne package -name "Declutter"

# Cross-compile (requires additional setup)
fyne package -os windows
fyne package -os darwin
fyne package -os linux
```

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
