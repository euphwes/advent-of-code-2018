package main

import (
    "aocUtils/inputUtils"
    "aocUtils/mathUtils"
    "fmt"
    "strconv"
)

// ------------------------------------------------------------------------------------------------

const DayNumber = 1
const RunSampleInput = false

// ------------------------------------------------------------------------------------------------

// this Set interface isn't strictly necessary here, nor is the IntSet struct below, but I might
// as well practice interfaces and types which implement them
type Set interface {
    Add(interface{})
    Contains(interface{}) bool
}

type IntSet struct {
    set map[int]bool
}

func (intSet *IntSet) Add(n int) {
    intSet.set[n] = true
}

func (intSet *IntSet) Contains(n int) bool {
    _, found := intSet.set[n]
    return found
}

// Converts a list of strings representing numbers to a list of integers
func convertStringsToInts(input []string) ([]int, error) {
    var inputAsInt []int
    for _, elem := range input {
        num, err := strconv.Atoi(elem)
        if (err != nil) {
            return nil, err
        }
        inputAsInt = append(inputAsInt, num)
    }

    return inputAsInt, nil
}

// Sums all of the input frequency adjustments to determine the final frequency
// and returns that
func part1(input []string) int {
    frequencies, err := convertStringsToInts(input)
    if err != nil {
        panic(err)
    }
    return mathUtils.SumInts(frequencies)
}

// Continuously sums all the frequency adjustments, tracking the current frequency,
// until the current frequency has already been seen. Returns that value
func part2(input []string) int {
    frequencies, err := convertStringsToInts(input)
    if err != nil {
        panic(err)
    }

    intSet := IntSet{
        set: map[int]bool{},
    }
    currentFrequency := 0
    duplicateFrequency := 0

    for duplicateFrequency == 0 {
        for i := 0; i < len(frequencies); i++ {
            currentFrequency += frequencies[i]
            if intSet.Contains(currentFrequency) {
                duplicateFrequency = currentFrequency
                break
            } else {
                intSet.Add(currentFrequency)
            }
        }
    }

    return duplicateFrequency
}

// ------------------------------------------------------------------------------------------------

func main() {
    lines, err := inputUtils.GetInputForDay(DayNumber, RunSampleInput)
    if err != nil {
        panic(err)
    }

    part1Answer := part1(lines)
    part2Answer := part2(lines)

    fmt.Printf("\nPart 1: %d",   part1Answer)
    fmt.Printf("\nPart 2: %d\n", part2Answer)
}