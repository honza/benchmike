package main

import (
    "fmt"
    "flag"
    "sort"
    "strings"
    "os"
    "os/exec"
    "time"
)

// Command-line arguments
var iterations = flag.Int("i", 100, "Number of iterations to run the command")
var verbose = flag.Bool("verbose", false, "Show iterations")

func Rank(scores []float64 , yourScore float64) int {
    count := 0

    for i := 0; i < len(scores); i++ {
        score := scores[i]
        if score <= yourScore {
            count++
        }
    }

    return 100.0 * count / len(scores)
}

func Percentile(scores []float64, rank int) float64 {
    sort.Float64s(scores)
    for i := 0; i < len(scores); i++ {
        score := scores[i]
        if Rank(scores, score) >= rank {
            return score
        }
    }

    return 0.0
}

func run(cmd exec.Cmd, verbose bool) time.Duration {
    if verbose == true {
        fmt.Println("Running", cmd)
    }

    start := time.Now()
    err := cmd.Run()
    end := time.Now()

    if err != nil && verbose == true {
        fmt.Println(err)
    }

    return end.Sub(start)
}

func usage() {
    fmt.Fprintf(os.Stderr, "usage: benchmike OPTIONS COMMAND\n")
    flag.PrintDefaults()
    os.Exit(2)
}

func parseCommand(cmd string) *exec.Cmd {
    parts := strings.Split(cmd, " ")
    command := parts[0]
    args := parts[1:]
    return exec.Command(command, args...)

}

func main() {
    flag.Usage = usage
    flag.Parse()

    results := make([]float64, *iterations)

    if (len(flag.Args()) < 1) {
        usage()
    }

    cmd := parseCommand(flag.Args()[0])

    for i := 0; i < *iterations; i++ {
        results[i] = float64(run(*cmd, *verbose).Nanoseconds()) / 1000000
    }

    sort.Float64s(results)

    fmt.Println("95th percentile:", Percentile(results, 95))
    fmt.Println("99th percentile:", Percentile(results, 99))
    fmt.Println("Min:            ", results[0])
    fmt.Println("Max:            ", results[len(results) - 1:][0])

}
