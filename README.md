# botrecept

# if BUILD error
runtime/cgo
# runtime/cgo
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%

The error message "C compiler "gcc" not found" indicates that the Go compiler is unable to find the C compiler on your system. This is a common issue when working with the runtime/cgo package in Go, which allows you to call C code from your Go program.

To resolve this issue, you need to ensure that you have a C compiler installed on your system. The most common C compiler is GCC (GNU Compiler Collection), which is available on most Unix-like systems (including Linux and macOS) and can also be installed on Windows.

Here are the steps to resolve the issue on different operating systems:

Windows:
Install the MinGW (Minimalist GNU for Windows) toolchain, which includes the GCC compiler. You can download MinGW from the official website: https://sourceforge.net/projects/mingw/
After installing MinGW, make sure to add the MinGW bin directory to your system's PATH environment variable, so that the Go compiler can find the gcc executable.
macOS:
Install the Xcode Command Line Tools, which include the GCC compiler. You can do this by running the following command in the Terminal:
Code

Copy Code
xcode-select --install
Linux:
Install the GCC compiler package using your system's package manager. For example, on Ubuntu/Debian, you can run:
Code

Copy Code
sudo apt-get install build-essential
On other Linux distributions, the package name may be different, but the general process is the same: install the package that provides the GCC compiler.