Text Transformation Tool

This Go program (main.go) reads a text file, applies various transformations based on specified markers within the text, and then writes the modified text to a new file. The transformations include changing the case of words, converting hexadecimal and binary representations to integers, handling articles ("a" to "an"), and managing punctuation.
Usage

To use this tool, follow these steps:

    Compile and Run
        Ensure you have Go installed on your system.
        Compile the program: go build main.go
        Run the executable passing two arguments:

        bash

    ./main sample.txt result.txt

    Replace sample.txt with the path to your input text file and result.txt with the desired output file name.

Markers for Text Transformation

    The input file can contain specific markers enclosed in parentheses () to indicate transformations:
        (up): Convert the preceding word to uppercase.
        (low): Convert the preceding word to lowercase.
        (cap): Capitalize the first letter of the preceding word.
        (hex): Convert the preceding word from hexadecimal to decimal.
        (bin): Convert the preceding word from binary to decimal.
        (up,n): Convert the preceding n words to uppercase.
        (low,n): Convert the preceding n words to lowercase.
        (cap,n): Capitalize the first letter of the preceding n words.

Example Input File Content

    An example input file (sample.txt) might contain:

        Hello world(up), this is a test(cap) 4e4(hex) 101010(bin).

        Here:
            (up) converts "world" to "WORLD".
            (cap) capitalizes "test".
            (hex) converts "4e4" to "1252".
            (bin) converts "101010" to "42".

    Output
        The transformed text is written to the specified output file (result.txt in the example).

Note

    Ensure the input file exists and is readable.
    The program handles basic transformations and punctuation adjustments as described in the source code (main.go).
    Errors during file operations or transformations are displayed on the console.

Author

This tool was created by [Thadeus Ogondola].

For any issues or improvements, please contact [obimbira60@gmail.com].