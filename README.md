
# ccwc Tool

Welcome to the `ccwc` tool! This project is a Go implementation of the classic Unix `wc` (word count) command-line tool, designed to work on Windows, macOS, and Linux. Like its sibling `wc`, `ccwc` is used to count bytes, characters, words, and lines in files or from standard input.

## Features

- **Count Bytes (-c)**: Use the `c` flag to get the total number of bytes in a file.
- **Count Lines (-l)**: The `l` flag provides the number of lines in the file.
- **Count Words (-w)**: Use the `w` flag to count the number of words in the file.
- **Count Characters (-m)**: The `m` flag counts characters in the file, supporting multibyte characters.
- **Default Mode**: When no flags are provided, the tool will display the byte, word, and line counts for the file.
- **Support for Standard Input**: The tool can also read from standard input if no file is provided, making it useful in pipelines.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/rafaelmgr12/cwcc-tool.git
    cd ccwc-tool
    ```

2. Build the tool using Go:

    ```bash
    go build -o ccwc
    ```

3. You can now run the `ccwc` tool on your system.

## Usage

`ccwc` provides the following options:

- Count the number of bytes in a file:

    ```bash
    ./ccwc -c filename.txt
    ```

- Count the number of lines in a file:

    ```bash
    ./ccwc -l filename.txt
    ```

- Count the number of words in a file:

    ```bash
    ./ccwc -w filename.txt
    ```

- Count the number of characters in a file:

    ```bash
    ./ccwc -m filename.txt
    ```

- Default: Get the line, word, and byte count without any flags:

    ```bash
    ./ccwc filename.txt
    ```

- Read from standard input:

    ```bash
    cat filename.txt | ./ccwc -l
    ```

## Example

Here’s how you can use `ccwc` with a sample text file:

1. **File `test.txt`:**

    ```text
    Go is an open-source programming language that makes it easy to build simple, reliable, and efficient software.
    ```

2. **Running `ccwc` on `test.txt`:**

    ```bash
    ./ccwc -w test.txt
    ```

    **Output:**

    ```bash
    12 test.txt
    ```

    This shows that `test.txt` contains 12 words.

## Using `ccwc` as a CLI Tool (Setting it in your PATH)

To make `ccwc` available as a command-line tool that you can run from any directory (similar to `wc`), follow these steps to add it to your system's `PATH`:

### Linux/macOS

1. After building `ccwc`, move the binary to a directory that's in your system's `PATH` (for example, `/usr/local/bin`):

    ```bash
    sudo mv ./ccwc /usr/local/bin/ccwc
    ```

2. Verify that the tool is now accessible from anywhere:

    ```bash
    ccwc -h
    ```

   If you see the help output for `ccwc`, you've successfully set it up!

### Windows

1. After building `ccwc`, move the binary to a directory that is included in your system's `PATH`. For example, you can place it in `C:\Program Files\`:

    ```powershell
    Move-Item .\ccwc.exe 'C:\Program Files\ccwc.exe'
    ```

2. Add the directory to the system `PATH`:

    1. Open the **Start Menu**, search for "Environment Variables," and select "Edit the system environment variables."
    2. In the "System Properties" window, click "Environment Variables."
    3. Under "System Variables," find the `Path` variable, select it, and click "Edit."
    4. Click "New" and add the path to where you moved `ccwc.exe` (e.g., `C:\Program Files\`).
    5. Click "OK" to close all windows.

3. Verify that `ccwc` is now accessible from the command prompt:

    ```bash
    ccwc -h
    ```

   If you see the help output for `ccwc`, you've successfully set it up!

## Contributing

Contributions are welcome! If you would like to contribute, feel free to fork the repository, make your changes, and open a pull request. Please ensure that your changes are well-documented and that you write tests for any new functionality.

1. Fork this repository
2. Create your feature branch:

    ```bash
    git checkout -b feature/YourFeatureName
    ```

3. Commit your changes:

    ```bash
    git commit -m "Add feature description"
    ```

4. Push to the branch:

    ```bash
    git push origin feature/YourFeatureName
    ```

5. Open a pull request

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
