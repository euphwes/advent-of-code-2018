package main

import (
    "aocUtils/inputUtils"
    "fmt"
    "strconv"
    "strings"
    "time"
    "sort"
)

// ------------------------------------------------------------------------------------------------

const DayNumber = 4
const RunSampleInput = false

// ------------------------------------------------------------------------------------------------

type watchRecord struct {
    Timestamp   time.Time
    Description string
    ElfId       int
}

type watchRecords []watchRecord

func (records watchRecords) Less(i, j int) bool {
    return records[i].Timestamp.Before(records[j].Timestamp)
}

func (records watchRecords) Swap(i, j int) {
    records[i], records[j] = records[j], records[i] 
}

func (records watchRecords) Len() int {
    return len(records)
}

// ------------------------------------------------------------------------------------------------

func buildWatchRecords(input []string) watchRecords {
    records := make(watchRecords, 0)

    for _, line := range input {
        lastIndexBracket := strings.Index(line, "]")
        dateStr := line[1:lastIndexBracket]
        layout := "2006-01-02 15:04"
        timestamp, _ := time.Parse(layout, dateStr)

        newWatchRecord := watchRecord{
            Timestamp: timestamp,
            Description: line[lastIndexBracket+2:],
        }

        if strings.Contains(newWatchRecord.Description, "begins") {
            hashIndex := strings.Index(newWatchRecord.Description, "#")
            endIndex := strings.Index(newWatchRecord.Description, " begins")
            elfId, _ := strconv.Atoi(newWatchRecord.Description[hashIndex + 1 : endIndex])
            newWatchRecord.ElfId = elfId
        }

        records = append(records, newWatchRecord)
    }

    sort.Sort(records)
    return records
}

func buildElfSleepMinuteMaps(records watchRecords) (map[int]int, map[int]map[int]int) {
    currentElfId := 0
    sleepStopMinute := 0
    sleepStartMinute := 0

    elfTotalSleepMinutes := map[int]int{}
    elfMinuteSleepMap := map[int]map[int]int{}

    for _, r := range records {
        if strings.Contains(r.Description, "begins") {
            currentElfId = r.ElfId
            _, found := elfMinuteSleepMap[currentElfId]
            if !found {
                elfMinuteSleepMap[currentElfId] = map[int]int{}
            }
            _, found = elfTotalSleepMinutes[currentElfId]
            if !found {
                elfTotalSleepMinutes[currentElfId] = 0
            }
        }
        if strings.Contains(r.Description, "asleep") {
            sleepStartMinute = r.Timestamp.Minute()
        }
        if strings.Contains(r.Description, "wakes") {
            sleepStopMinute = r.Timestamp.Minute()
            minuteMap, _ := elfMinuteSleepMap[currentElfId]

            for i := sleepStartMinute; i < sleepStopMinute; i++ {
                _, found := minuteMap[i]
                if !found {
                    minuteMap[i] = 1
                } else {
                    minuteMap[i] += 1
                }
            }
            elfTotalSleepMinutes[currentElfId] += (sleepStopMinute - sleepStartMinute)
            elfMinuteSleepMap[currentElfId] = minuteMap
        }
    }

    return elfTotalSleepMinutes, elfMinuteSleepMap
}

func part1(elfTotalSleepMinutes map[int]int, elfMinuteSleepMap map[int]map[int]int) int {
    sleepiestElfId := 0
    longestSleepCount := 0

    for elfId, sleepCount := range elfTotalSleepMinutes {
        if sleepCount > longestSleepCount {
            longestSleepCount = sleepCount
            sleepiestElfId = elfId
        }
    }

    sleepiestMinute := 0
    longestSleepLength := 0
    for minute, sleepLength := range elfMinuteSleepMap[sleepiestElfId] {
        if sleepLength > longestSleepLength {
            longestSleepLength = sleepLength
            sleepiestMinute = minute
        }
    }

    return sleepiestElfId * sleepiestMinute
}

func part2(elfMinuteSleepMap map[int]map[int]int) int {
    sleepiestElfId := 0
    mostSleptMinute := 0
    longestMinuteSleepLength := 0

    for elfId, innerMap := range elfMinuteSleepMap {
        for minute, minuteSleepLength := range innerMap {
            if minuteSleepLength > longestMinuteSleepLength {
                longestMinuteSleepLength = minuteSleepLength
                sleepiestElfId = elfId
                mostSleptMinute = minute
            }
        }
    }

    return sleepiestElfId * mostSleptMinute
}

// ------------------------------------------------------------------------------------------------

func main() {
    lines, err := inputUtils.GetInputForDay(DayNumber, RunSampleInput)
    if err != nil {
        panic(err)
    }

    records := buildWatchRecords(lines)
    elfTotalSleepMinutes, elfMinuteSleepMap := buildElfSleepMinuteMaps(records)

    part1Answer := part1(elfTotalSleepMinutes, elfMinuteSleepMap)
    part2Answer := part2(elfMinuteSleepMap)

    fmt.Printf("\nPart 1: %d",   part1Answer)
    fmt.Printf("\nPart 2: %d\n", part2Answer)
}