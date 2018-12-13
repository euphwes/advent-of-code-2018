package main

import (
    "aocUtils/inputUtils"
    "aocUtils/mathUtils"
    "fmt"
    "strconv"
    "sort"
    "strings"
)

// ------------------------------------------------------------------------------------------------

const DayNumber = 6
const RunSampleInput = false

// ------------------------------------------------------------------------------------------------

type Point struct {
    x int
    y int
}

func ManhattanDistanceBetween(p1, p2 Point) int {
    return mathUtils.Abs(p2.x - p1.x) + mathUtils.Abs(p2.y - p1.y)
}

func buildPoints(input []string) []Point {
    points := []Point{}
    for _, line := range input {
        splitLine := strings.Split(line, ", ")
        x, _ := strconv.Atoi(splitLine[0])
        y, _ := strconv.Atoi(splitLine[1])
        newPoint := Point{
            x: x,
            y: y,
        }
        points = append(points, newPoint)
    }

    return points
}

func determineBoundaries(points []Point) (int, int, int, int) {
    minX, maxX, minY, maxY := 1000000, 0, 1000000, 0
    for _, p := range points {
        if (p.x > maxX) {
            maxX = p.x
        } else if (p.x < minX) {
            minX = p.x
        }
        if (p.y > maxY) {
            maxY = p.y
        } else if (p.y < minY) {
            minY = p.y
        }
    }
    return minX, maxX, minY, maxY
}

func determineClosestTargetPoint(p Point, targetPoints []Point) (Point, bool) {
    var winningPoint Point
    currentShortestDistance := 100000000

    var lowestDistanceCount int

    for _, tp := range targetPoints {
        distance := ManhattanDistanceBetween(p, tp)
        if distance == currentShortestDistance {
            lowestDistanceCount += 1
        }
        if distance < currentShortestDistance {
            winningPoint = tp
            currentShortestDistance = distance
            lowestDistanceCount = 1
        }
    }

    if lowestDistanceCount > 1 {
        return Point{}, true
    }

    return winningPoint, false
}

func getDistancesToTargets(p Point, targets []Point) []int {
    distances := []int{}
    for _, tp := range targets {
        distances = append(distances, ManhattanDistanceBetween(p, tp))
    }
    return distances
}

func buildPointsBoundedBy(minX, maxX, minY, maxY int) []Point {
    gridPoints := []Point{}
    for x := mathUtils.Min(minX, 0); x <= (maxX); x++ {
        for y := mathUtils.Min(minY, 0); y <= (maxY); y++ {
            newPoint := Point {
                x: x,
                y: y,
            }
            gridPoints = append(gridPoints, newPoint)
        }
    }
    return gridPoints
}

func part1(input []string) int {
    targetPoints := buildPoints(input)
    minX, maxX, minY, maxY := determineBoundaries(targetPoints)

    gridPoints := buildPointsBoundedBy(minX, maxX, minY, maxY)

    surroundingAreaMap := map[Point]int{}
    for _, p := range targetPoints {
        surroundingAreaMap[p] = 0
    }

    excludedTargets := map[Point]bool{}
    for _, gp := range gridPoints {
        closestTargetPoint, wasEquidistant := determineClosestTargetPoint(gp, targetPoints)
        if wasEquidistant { continue }

        if (gp.x == minX || gp.y == minY || gp.x == maxX || gp.y == maxY) {
            excludedTargets[closestTargetPoint] = true
        }

        surroundingAreaMap[closestTargetPoint] += 1
    }

    surroundingAreaList := []int{}
    for p, area := range surroundingAreaMap {
        _, found := excludedTargets[p]
        if found { continue }

        surroundingAreaList = append(surroundingAreaList, area)
    }

    sort.Ints(surroundingAreaList)

    return surroundingAreaList[len(surroundingAreaList)-1]
}

func part2(input []string) int {
    targetPoints := buildPoints(input)
    minX, maxX, minY, maxY := determineBoundaries(targetPoints)

    gridPoints := buildPointsBoundedBy(minX, maxX, minY, maxY)

    safeCellCount := 0
    for _, p := range gridPoints {
        distancesToTargets := getDistancesToTargets(p, targetPoints)
        if mathUtils.SumInts(distancesToTargets) < 10000 {
            safeCellCount += 1
        }
    }

    return safeCellCount
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