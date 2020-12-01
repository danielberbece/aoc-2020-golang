package main

import (
    "fmt"
    "strconv"
    "log"
    "os"
    "bufio"
)

func ReadInts(filename string) ([]int, error) {
    file, err := os.Open("./input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    var result []int
    for scanner.Scan() {
        x, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, scanner.Err()
}

func part1(array []int) {
    for i := 0; i < len(array); i++ {
        for j := i; j < len(array); j++ {
            if array[i] + array[j] == 2020 {
                fmt.Printf("Part 1: %d\n", array[i] * array[j])
                return
            }
        }
    }
}

func part2(array []int) {
    for i := 0; i < len(array); i++ {
        for j := i; j < len(array); j++ {
            for k := j; k < len(array); k++ {
                if array[i] + array[j] + array[k] == 2020 {
                    fmt.Printf("Part 2: %d\n", array[i] * array[j] * array[k])
                    return
                }
            }
        }
    }
}

func main() {
    array, err := ReadInts("./input")
    if (err != nil) {
        return
    }
    part1(array)
    part2(array)
}
