package main

import (
    "aocUtils/inputUtils"
    "fmt"
    "strings"
    "sort"
)

// ------------------------------------------------------------------------------------------------

const DayNumber = 5
const RunSampleInput = false

// ------------------------------------------------------------------------------------------------

func buildDestructionPairs() []string {
    pairs := []string{}
    for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
        upper := strings.ToUpper(string(letter))
        pairs = append(pairs, string(letter) + upper)
        pairs = append(pairs, upper + string(letter))
    }
    return pairs
}

func reducePolymer(polymer string) string {
    pairs := buildDestructionPairs()
    for true {
        anyPairsDestroyed := false
        for _, destructionPair := range pairs {
            if strings.Contains(polymer, destructionPair) {
                anyPairsDestroyed = true
                polymer = strings.Replace(polymer, destructionPair, "", -1)
            }
        }
        if !anyPairsDestroyed {
            break
        }
    }
    return polymer
}

func part1(polymer string) int {
    return len(reducePolymer(polymer))
}

func part2(polymer string) int {
    possibleReducedPolymerLengths := []int{}

    for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
        lower := string(letter)
        upper := strings.ToUpper(lower)
        testPolymer := strings.Replace(polymer, lower, "", -1)
        testPolymer = strings.Replace(testPolymer, upper, "", -1)

        reducedPolymer := reducePolymer(testPolymer)
        possibleReducedPolymerLengths = append(possibleReducedPolymerLengths, len(reducedPolymer))
    }

    sort.Ints(possibleReducedPolymerLengths)
    return possibleReducedPolymerLengths[0]
}

// ------------------------------------------------------------------------------------------------

func main() {
    lines, err := inputUtils.GetInputForDay(DayNumber, RunSampleInput)
    if err != nil {
        panic(err)
    }

    part1Answer := part1(lines[0])
    part2Answer := part2(lines[0])

    fmt.Printf("\nPart 1: %d",   part1Answer)
    fmt.Printf("\nPart 2: %d\n", part2Answer)
}