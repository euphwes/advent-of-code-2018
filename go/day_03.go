package main

import (
    "aocUtils/inputUtils"
    //"errors"
    "fmt"
    "strconv"
    "strings"
)

// ------------------------------------------------------------------------------------------------

const DayNumber = 3
const RunSampleInput = false

type fabricPointStatus int
const (
    Free    fabricPointStatus = 0
    Claimed fabricPointStatus = 1
    Overlap fabricPointStatus = 2
)

// ------------------------------------------------------------------------------------------------

type point struct {
    X int
    Y int
}

type fabricClaim struct {
    Id     string
    Points []point
}

type fabricState struct {
    Points [][]fabricPointStatus
}

func buildCleanFabricState(length, height int) fabricState {
    newFabricState := fabricState{
        Points: make([][]fabricPointStatus, height),
    }
    for i := range newFabricState.Points {
        newFabricState.Points[i] = make([]fabricPointStatus, length)
    }

    for x := 0; x < length; x++ {
        for y := 0; y < height; y++ {
            newFabricState.Points[x][y] = Free
        }
    }

    return newFabricState
}

func applyClaimToFabricState(claim fabricClaim, state fabricState) fabricState {
    for _, point := range claim.Points {
        x := point.X
        y := point.Y

        currentValue := state.Points[x][y]
        if currentValue == Free {
            state.Points[x][y] = Claimed
        } else if currentValue == Claimed {
            state.Points[x][y] = Overlap
        }
    }

    return state
}

func checkClaimForNoOverlap(claim fabricClaim, state fabricState) bool {
    overlapFree := true
    for _, point := range claim.Points {
        x := point.X
        y := point.Y

        currentValue := state.Points[x][y]
        if currentValue == Overlap {
            overlapFree = false
            break
        }
    }

    return overlapFree
}

func buildFabricClaim(input string) fabricClaim {
    // input looks like:
    // #123 @ 3,2: 5x4

    // firstSplit:
    //      #123
    //      3,2: 5x4
    firstSplit := strings.Split(input, "@")
    id := strings.Trim(firstSplit[0], " ")

    // secondSplit:
    //      3,2
    //      5x4
    secondSplit := strings.Split(firstSplit[1], ":")

    // thirdSplit:
    //      3
    //      2
    thirdSplit := strings.Split(secondSplit[0], ",")
    startingX, _ := strconv.Atoi(strings.Trim(thirdSplit[0], " "))
    startingY, _ := strconv.Atoi(strings.Trim(thirdSplit[1], " "))

    // fourthSplit:
    //      5
    //      4
    fourthSplit := strings.Split(secondSplit[1], "x")
    length, _ := strconv.Atoi(strings.Trim(fourthSplit[0], " "))
    height, _ := strconv.Atoi(strings.Trim(fourthSplit[1], " "))

    var points []point
    for x := 0; x < length; x++ {
        for y := 0; y < height; y++ {
            newPoint := point{
                X: startingX + x,
                Y: startingY + y,
            }
            points = append(points, newPoint)
        }
    }

    return fabricClaim{
        Id:     id,
        Points: points,
    }
}

func part1(input []string) int {
    var claims []fabricClaim
    for _, line := range input {
        claims = append(claims, buildFabricClaim(line))
    }

    santaFabric := buildCleanFabricState(1000, 1000)

    for _, claim := range claims {
        santaFabric = applyClaimToFabricState(claim, santaFabric)
    }

    overlapped := 0
    for _, row := range santaFabric.Points {
        for _, cell := range row {
            if cell == Overlap {
                overlapped += 1
            }
        }
    }

    return overlapped
}

func part2(input []string) string {
    var claims []fabricClaim
    for _, line := range input {
        claims = append(claims, buildFabricClaim(line))
    }

    santaFabric := buildCleanFabricState(1000, 1000)

    for _, claim := range claims {
        santaFabric = applyClaimToFabricState(claim, santaFabric)
    }

    claimWithNoOverlap := ""
    for _, claim := range claims {
        if checkClaimForNoOverlap(claim, santaFabric) {
            claimWithNoOverlap = claim.Id
            break
        }
        santaFabric = applyClaimToFabricState(claim, santaFabric)
    }

    return claimWithNoOverlap
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
    fmt.Printf("\nPart 2: %s\n", part2Answer)
}