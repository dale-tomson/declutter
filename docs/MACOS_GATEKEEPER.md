# macOS "Not Opened" / Malware Warning

## Is This Safe?

**Yes, Declutter is completely safe.** This warning appears because the application is not notarized by Apple, which requires a paid Apple Developer account ($99/year).

### Why Does This Happen?

macOS includes a security feature called **Gatekeeper** that blocks applications from unidentified developers. Since Declutter is a free, open-source project, we currently do not have a paid Apple Developer certificates to sign the application.

This is a common "hurdle" for open-source software on macOS. You can verify the source code is open and safe at [github.com/dale-tomson/declutter](https://github.com/dale-tomson/declutter).

## Solution

### The "Right-Click" Workaround

You only need to do this **once** when you first open the application:

1.  **Locate the App**: Open Finder and find the `Declutter` app (usually in your `Applications` folder or `Downloads`).
2.  **Right-Click**: Right-click (or Control-click) on the `Declutter` icon.
3.  **Select Open**: Choose **Open** from the context menu.
4.  **Confirm**: A dialog box will appear asking if you're sure you want to open it. Click **Open**.

After doing this once, macOS will remember your choice and you can open Declutter normally by double-clicking it in the future.

> [!NOTE]
> If you try to double-click the app *without* doing the right-click step first, you will simply see a message saying the app "cannot be opened" with only a "Move to Bin" or "Cancel" option. You *must* use the right-click menu to see the "Open" option.

## Verification

You can verify the authenticity of Declutter:

### Check the Source Code
- All code is open source: [github.com/dale-tomson/declutter](https://github.com/dale-tomson/declutter)

### Verify the Download
- Only download from official sources:
  - [GitHub Releases](https://github.com/dale-tomson/declutter/releases)
  - [Official Website](https://dale-tomson.github.io/declutter)
