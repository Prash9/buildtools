package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    var filepath string

	// remove program path from list
    // filepath := flag.String("filepath", "", "File Path for the program")
    size_flag := flag.Bool("c", false, "Get byte size of the file")
    noOfLines := flag.Bool("l", false, "Get number of lines in the file")
    noOfWords := flag.Bool("w", false, "Get number of words in the file")
    flag.Parse()
    
    filepath = os.Args[len(os.Args)-1]
    
    fmt.Println("File Name: ", filepath)
    if filepath == "" {
        flag.PrintDefaults()
        os.Exit(1)
    }

    fileObj, err := os.Open(filepath)
    if err != nil {
        fmt.Println("Error reading file: ", err)
        os.Exit(1)
    }
    defer fileObj.Close()


    if *size_flag {
        filestat := getFileStat(filepath)
        fmt.Println("File Size(kb):", filestat.Size())
    }
    
    if *noOfLines {
        fmt.Println("Number of lines in file:", getNoOfLines(fileObj))
        // reset file obj pointer
        fileObj.Seek(0,0)
    }

    if *noOfWords {
        fmt.Println("Number of words in file:", getNoOfWords(fileObj))
        // reset file obj pointer
        fileObj.Seek(0,0)
    }

}


func getFileStat(filepath string) os.FileInfo {
    filestat, err := os.Stat(filepath)
    // fmt.Printf("error 1: %T", err)
    if err != nil {
        fmt.Println("Error reading file: ", err)
        os.Exit(1)
    }
    return filestat
}

func getNoOfLines(fileObj *os.File) int {
    
    lineCount := 0
    scanner := bufio.NewScanner(fileObj)
    for scanner.Scan() {
        lineCount++
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning file: ", err)
        os.Exit(1)
    }

    return lineCount
}


func getNoOfWords(fileObj *os.File) int {
    
    wordCount := 0
    scanner := bufio.NewScanner(fileObj)
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