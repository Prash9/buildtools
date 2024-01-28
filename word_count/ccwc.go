package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "strings"
    "unicode/utf8"
    "io"
    "io/ioutil"
)

// Command to run
//  ./ccwc -c -l -w test.txt
var fileObj *os.File

func main() {
    var remainingArgs []string
    var filepath string
    var scanner *bufio.Scanner

	// remove program path from list
    // filepath := flag.String("filepath", "", "File Path for the program")
    size_flag := flag.Bool("c", false, "Get byte size of the file")
    noOfLines := flag.Bool("l", false, "Get number of lines in the file")
    noOfWords := flag.Bool("w", false, "Get number of words in the file")
    noOfChars := flag.Bool("m", false, "Get number of charcters in the file")
    flag.Parse()

    if !(*size_flag || *noOfLines || *noOfWords) {
        *size_flag = true
        *noOfLines = true
        *noOfWords = true
    }
    
    // filepath = os.Args[len(os.Args)-1]
    remainingArgs = flag.Args()
    // stat, _ := os.Stdin.Stat()
    
    if len(remainingArgs) == 1 {
        var err error
        filepath = remainingArgs[0]
        fileObj, err = os.Open(filepath)
        if err != nil {
            fmt.Println("Error reading file: ", err)
            os.Exit(1)
        }
        defer fileObj.Close()
        if fileObj != nil {
            scanner = getScanner(fileObj)
        }

    } else {
        fileObj = getFileObjForStdin(os.Stdin)
        scanner = getScanner(fileObj)
    }

    if *size_flag {
        filesize := getFileStat(fileObj)
        fmt.Println("File Size(kb):", filesize)
    }
    
    if *noOfLines {
        scanner := getScanner(fileObj)
        fmt.Println("Number of lines in file:", getNoOfLines(scanner)) 
    }

    if *noOfWords {
        scanner = getScanner(fileObj)
        fmt.Println("Number of words in file:", getNoOfWords(scanner))
    }

    if *noOfChars {
        scanner = getScanner(fileObj)
        fmt.Println("Number of charcters in file:", getNoOfChars(scanner))
    }

}


func getNoOfChars(scanner *bufio.Scanner) int {
    
    charCount := 0
    // scanner := bufio.NewScanner(fileObj)
    for scanner.Scan() {
        line := scanner.Text()
        charCount += utf8.RuneCountInString(line) + 1 // add new line charcter
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning file: ", err)
        os.Exit(1)
    }
    
    return charCount 
}

func getFileStat(fileObj *os.File) int64 {
    fileInfo, err := fileObj.Stat()
    if err != nil {
        fmt.Println("Error reading file: ", err)
        os.Exit(1)
    }
    return fileInfo.Size()
}

func getNoOfLines(scanner *bufio.Scanner) int {
    
    lineCount := 0
    // scanner := bufio.NewScanner(fileObj)
    // fmt.Printf("%T \n", scanner)
    for scanner.Scan() {
        // fmt.Println(lineCount)
        lineCount++
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning file: ", err)
        os.Exit(1)
    }

    return lineCount
}


func getNoOfWords(scanner *bufio.Scanner) int {
    
    wordCount := 0
    // scanner := bufio.NewScanner(fileObj)
    for scanner.Scan() {
        line := scanner.Text()
        words := strings.Fields(line)
        wordCount = wordCount + len(words)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning file: ", err)
        os.Exit(1)
    }

    return wordCount
}

func getFileObjForStdin(stdin *os.File) *os.File{
    tmpFile, err := ioutil.TempFile("", "stdin-*.txt")
    if err != nil {
		fmt.Println("Error creating temporary file:", err)
		os.Exit(1)
	}
	defer os.Remove(tmpFile.Name()) // Remove the temporary file when done

	// Copy input from os.Stdin to the temporary file
	_, err = io.Copy(tmpFile, stdin)
	if err != nil {
		fmt.Println("Error copying input to temporary file:", err)
		os.Exit(1)
	}
    return tmpFile
}

func getScanner(fileObj *os.File) *bufio.Scanner {
    // ensure file seek at the beginning
    if _, err := fileObj.Seek(0, 0); err != nil {
        // Handle error or provide a message if seeking is not supported
        fmt.Println("Warning: Seeking not supported on this input stream.")
    }
    
    stat, _ := fileObj.Stat()

    if (stat.Mode() & os.ModeCharDevice) == 0 {
        // Input is from a pipe or file
        return bufio.NewScanner(fileObj)
    }
    os.Exit(1)
    return nil
}