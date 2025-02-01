# botrecept

# if BUILD error
runtime/cgo
# runtime/cgo
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%

The error message "C compiler "gcc" not found" indicates that the Go compiler is unable to find the C compiler on your system. This is a common issue when working with the runtime/cgo package in Go, which allows you to call C code from your Go program.

To resolve this issue, you need to ensure that you have a C compiler installed on your system. The most common C compiler is GCC (GNU Compiler Collection), which is available on most Unix-like systems (including Linux and macOS) and can also be installed on Windows.

Here are the steps to resolve the issue on different operating systems:

Install a 64-bit version of MinGW:
Go to the MinGW download page: https://sourceforge.net/projects/mingw-w64/
Download the 64-bit version of MinGW, such as "x86_64-8.1.0-release-posix-seh-rt_v6-rev0.7z"
Extract the downloaded archive to a directory of your choice (e.g., C:\MinGW64)
Add the MinGW64 bin directory to your system's PATH environment variable:
Open the Windows Start menu and search for "Environment Variables"
Click on "Edit the system environment variables"
In the System Properties window, click on the "Environment Variables" button
In the Environment Variables window, find the "Path" variable in the "System Variables" section
Click on "Edit" and then "New"
Add the path to the bin directory of your MinGW64 installation (e.g., C:\MinGW64\bin)
Click "OK" to save the changes and close the windows

macOS:
Install the Xcode Command Line Tools, which include the GCC compiler. You can do this by running the following command in the Terminal:
xcode-select --install

Linux:
Install the GCC compiler package using your system's package manager. For example, on Ubuntu/Debian, you can run:
sudo apt-get install build-essential

On other Linux distributions, the package name may be different, but the general process is the same: install the package that provides the GCC compiler.