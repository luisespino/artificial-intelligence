package main

import (
    "fmt"
    "math"
    "math/rand"
    "sort"
    "time"
)

func absInt(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func average(arr []int) float64 {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return float64(sum) / float64(len(arr))
}

func sortedByDistance(pop []int, target int) []int {
    sorted := make([]int, len(pop))
    copy(sorted, pop)
    sort.Slice(sorted, func(i, j int) bool {
        return absInt(target-sorted[i]) < absInt(target-sorted[j])
    })
    return sorted
}

func genetic(target int, conv float64) {
    rand.Seed(time.Now().UnixNano())
    population := make([]int, 4)
    for i := 0; i < 4; i++ {
        population[i] = rand.Intn(10000)
    }

    populationAvg := average(population)
    populationConv := 0.0

    fmt.Println("Initial population:", population)
    fmt.Println("Initial convergence:", populationConv)

    for populationConv < conv {
        parents := sortedByDistance(population, target)

        child1 := (parents[0] + parents[1]) / 2
        child2 := 2*int(math.Max(float64(parents[0]), float64(parents[1]))) - int(math.Min(float64(parents[0]), float64(parents[1])))
        child3 := absInt(parents[0] - parents[1])
        child4 := int(float64(parents[0]) * 1.1)
        child5 := int(float64(parents[0]) * 0.9)

        children := []int{child1, child2, child3, child4, child5}
        children = sortedByDistance(children, target)

        population = append(parents[:2], children[:2]...)

        fmt.Println("New population:", population)

        newPopulationAvg := average(population)
        newPopulationConv := math.Min(populationAvg, newPopulationAvg) / math.Max(populationAvg, newPopulationAvg)

        populationAvg = newPopulationAvg
        populationConv = newPopulationConv

        fmt.Println("Convergence:", populationConv)
    }
}

func main() {
    genetic(1234, 1)
}
