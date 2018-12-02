package inputUtils

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
)

const (
    INPUTS_DIR = "inputs"
    INPUT_FILE_TEMPLATE = "day_%02d.txt"
    SAMPLE_FILE_TEMPLATE = "day_%02d_sample.txt"
)

// Reads the input file for the specified day, and returns the file contents as a string array.
// Expects the input files to be in a directory named `input` which resides in the current
// working directory, and the files to be named like `day_01.txt`, `day_07_sample.txt`,
// `day_14.txt`, etc
//
// dayNumber int - the Advent of Code challenge day number
// getSample bool - whether or not to return the specified day's sample input
func GetInputForDay(dayNumber int, getSample bool) ([]string, error) {
    var lines []string

    cwd, err := os.Getwd()
    if err != nil {
        return nil, err
    }

    fileNameTemplate := INPUT_FILE_TEMPLATE
    if (getSample) {
        fileNameTemplate = SAMPLE_FILE_TEMPLATE
    }

    fileName := fmt.Sprintf(fileNameTemplate, dayNumber)
    filePath := filepath.Join(cwd, INPUTS_DIR, fileName)

    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    return lines, nil
}
