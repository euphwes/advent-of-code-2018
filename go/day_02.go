package main

import (
    "aocUtils/inputUtils"
    "errors"
    "fmt"
)

// ------------------------------------------------------------------------------------------------

const DayNumber = 2
const RunSampleInput = false

// ------------------------------------------------------------------------------------------------

// Returns a map containing the count of each rune in the supplied string
func getLetterCounts(input string) map[rune]int {
    letterCounts := make(map[rune]int)
    for _, val := range input {
        count, found := letterCounts[val]
        if !found {
            letterCounts[val] = 1
        } else {
            letterCounts[val] = count + 1
        }
    }
    return letterCounts
}

// Calculates the "distance" between two supplied strings.
// The distance is defined as the count of the number of indices where
// the letter at that index differs between the two strings.
func calculateStringDistance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("Strings must be the same length")
    }

    distance := 0
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            distance += 1
        }
    }

    return distance, nil
}

// Returns a string which is comprised of only those characters which
// are the same between both supplied input strings at the same index.
func determineCommonLetters(a, b string) string {
    common := ""
    for i := 0; i < len(a); i++ {
        if a[i] == b[i] {
            common += string(a[i])
        }
    }
    return common
}

func part1(input []string) int {
    numHasDoubleLetters := 0
    numHasTripleLetters := 0

    for _, line := range input {
        haveIncrementedDouble := false
        haveIncrementedTriple := false
        
        letterCounts := getLetterCounts(line)
        for _, count := range letterCounts {
            if count == 2 && !haveIncrementedDouble {
                numHasDoubleLetters += 1
                haveIncrementedDouble = true
                continue
            } else if count == 3 && !haveIncrementedTriple {
                numHasTripleLetters += 1
                haveIncrementedTriple = true
                continue
            }
        }        
    }

    return numHasDoubleLetters * numHasTripleLetters
}

func part2(input []string) string {
    for _, a := range input {
        for _, b := range input {
            distance, err := calculateStringDistance(a, b)
            if err != nil {
                panic(err)
            }

            if distance == 1 {
                return determineCommonLetters(a, b)
            }
        }
    }

    return ""
}

// ------------------------------------------------------------------------------------------------

func main() {
    lines, err := inputUtils.GetInputForDay(DayNumber, RunSampleInput)
    fmt.Println(len(lines))
    if err != nil {
        panic(err)
    }

    part1Answer := part1(lines)
    part2Answer := part2(lines)

    fmt.Printf("\nPart 1: %d",   part1Answer)
    fmt.Printf("\nPart 2: %s\n", part2Answer)
}