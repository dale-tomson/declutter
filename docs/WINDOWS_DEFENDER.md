# Windows Defender False Positive Warning

## Is This Safe?

**Yes, Declutter is completely safe.** This is a false positive detection that commonly affects applications built with Go, especially GUI applications using the Fyne framework.

### Why Does This Happen?

Windows Defender's heuristic detection flags Go applications because:

1. **Go Binary Structure** - Go compiles to a unique binary format that differs from traditional C/C++ applications
2. **Unsigned Executable** - The application isn't code-signed with a certificate (costs $100-400/year)
3. **Cross-Compilation** - Building Windows executables on Linux can trigger additional flags
4. **Historical Context** - Some malware has been written in Go, causing overly aggressive detection

This is a **well-documented issue** affecting thousands of legitimate Go applications. You can verify the source code is open and safe at [github.com/dale-tomson/declutter](https://github.com/dale-tomson/declutter).

## Solutions

### Option 1: Add Windows Defender Exclusion (Recommended)

This is the fastest way to use Declutter without warnings:

1. **Open Windows Security**
   - Press `Win + I` to open Settings
   - Go to **Privacy & Security** → **Windows Security**
   - Click **Virus & threat protection**

2. **Add an Exclusion**
   - Under "Virus & threat protection settings", click **Manage settings**
   - Scroll down to **Exclusions**
   - Click **Add or remove exclusions**
   - Click **Add an exclusion** → **File**
   - Browse to and select `declutter.exe`

3. **Run Declutter**
   - The application will no longer be blocked

### Option 2: Restore from Quarantine

If Windows Defender already quarantined the file:

1. Open **Windows Security** → **Virus & threat protection**
2. Click **Protection history**
3. Find the entry for `declutter.exe`
4. Click on it and select **Actions** → **Allow on device**
5. The file will be restored

### Option 3: Submit to Microsoft

Help improve Windows Defender by reporting this false positive:

1. Visit [Microsoft Security Intelligence Submission Portal](https://www.microsoft.com/en-us/wdsi/filesubmission)
2. Submit `declutter.exe` for analysis
3. Microsoft will review and potentially update their definitions

> [!NOTE]
> This process can take several weeks, and you may need to resubmit for major version updates.

## Verification

You can verify the authenticity of Declutter:

### Check the Source Code
- All code is open source: [github.com/dale-tomson/declutter](https://github.com/dale-tomson/declutter)
- Review the code yourself or have a developer audit it

### Verify the Download
- Only download from official sources:
  - [GitHub Releases](https://github.com/dale-tomson/declutter/releases)
  - [Official Website](https://dale-tomson.github.io/declutter)
- Check SHA256 checksums (provided in release notes)

### Scan with Multiple Antivirus Tools
- Upload to [VirusTotal](https://www.virustotal.com) to see results from 70+ antivirus engines
- Most will show it as clean; a few may flag it due to Go binary characteristics

## Future Plans

We're working on reducing false positives:

- ✅ **Build Metadata** - Adding Windows resource information (version, company, description)
- 🔄 **Microsoft Submission** - Submitting releases to Microsoft for analysis
- 📋 **Code Signing** - Considering code signing certificates for future releases (requires ongoing investment)

## Still Have Concerns?

If you're uncomfortable adding an exclusion:

1. **Build from Source** - Follow the [build instructions](../README.md#building-from-source) to compile yourself
2. **Use in a VM** - Run Declutter in a virtual machine for isolation
3. **Wait for Updates** - As the application gains reputation, false positives may decrease

## Questions?

Open an issue on [GitHub](https://github.com/dale-tomson/declutter/issues) if you have questions or concerns about security.
