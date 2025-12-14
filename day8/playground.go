package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"math"
	"os"
	"slices"
	"strings"
)

var distMap = make(map[float64]string)
var sortedDists []float64
var circuitsArray = []string{}
var jPositions = []string{}

func powInt(x, y int) float64 {
	return math.Pow(float64(x), float64(y))
}

func getDistance(j1 string, j2 string) float64 {
	var x1, y1, z1, x2, y2, z2 int
	fmt.Sscanf(j1, "%d,%d,%d", &x1, &y1, &z1)
	fmt.Sscanf(j2, "%d,%d,%d", &x2, &y2, &z2)
	return math.Sqrt(powInt(x2-x1, 2) + powInt(y2-y1, 2) + powInt(z2-z1, 2))
}

func initDistances() {
	for _, j1 := range jPositions {
		for _, j2 := range jPositions {
			if j1 != j2 {
				dist := getDistance(j1, j2)
				distMap[dist] = fmt.Sprintf("%s-%s", j1, j2)
			}
		}
	}
	sortedDists = slices.Sorted(maps.Keys(distMap))
	// for i, dist := range sortedDists {
	// 	fmt.Printf("Distance %d: %f, Pair: %s\n", i+1, dist, distMap[dist])
	// }
}

func addJunctionToCircuit(junction string, circuitIdx int) {
	fmt.Printf("Adding junction %s to circuit %d\n", junction, circuitIdx)

	circuitsArray[circuitIdx] = circuitsArray[circuitIdx] + "-" + junction
}

func addNewCircuit(j1 string, j2 string) {
	fmt.Printf("Creating new circuit with junctions %s and %s\n", j1, j2)
	newCircuit := fmt.Sprintf("%s-%s", j1, j2)
	circuitsArray = append(circuitsArray, newCircuit)
}

func checkJunctionsConnected(j1 string, j2 string) bool {
	for _, circuit := range circuitsArray {
		j1InCircuit := strings.Contains(circuit, j1)
		j2InCircuit := strings.Contains(circuit, j2)
		if j1InCircuit && j2InCircuit {
			return true
		}
	}
	return false
}

func countTotalJunctionsInCircuits() int {
	totalJunctions := 0
	for _, circuit := range circuitsArray {
		junctions := strings.Split(circuit, "-")
		totalJunctions += len(junctions)
	}
	return totalJunctions
}

func setupCircuits() {
	connectionCount := 0
	// find the 10 shortest distances
	for _, dist := range sortedDists {
		fmt.Println("Processing distance:", dist)
		fmt.Println("Pairs:", distMap[dist])
		// Find the next shortest distance
		jPair := distMap[dist]
		// check if either junction is already in a circuit
		junctions := strings.Split(jPair, "-")
		j1InCircuit := -1
		j2InCircuit := -1
		for i, circuit := range circuitsArray {
			if strings.Contains(circuit, junctions[0]) {
				j1InCircuit = i
			}
			if strings.Contains(circuit, junctions[1]) {
				j2InCircuit = i
			}
		}

		if j1InCircuit == -1 && j2InCircuit == -1 {
			// neither junction is in a circuit, create a new one
			addNewCircuit(junctions[0], junctions[1])
			connectionCount++
		} else if j1InCircuit != -1 && j2InCircuit == -1 {
			// j1 is in a circuit, add j2 to it
			addJunctionToCircuit(junctions[1], j1InCircuit)
			connectionCount++
		} else if j1InCircuit == -1 && j2InCircuit != -1 {
			// j2 is in a circuit, add j1 to it
			addJunctionToCircuit(junctions[0], j2InCircuit)
			connectionCount++
		} else {
			fmt.Println("Both junctions already in circuits, combine circuits if different")
			// both junctions are in circuits, check if they are different circuits
			if j1InCircuit != j2InCircuit {
				// combine circuits
				fmt.Printf("Combining circuits %d and %d\n", j1InCircuit, j2InCircuit)
				circuit1 := circuitsArray[j1InCircuit]
				circuit2 := circuitsArray[j2InCircuit]
				combinedCircuit := circuit1 + "-" + circuit2
				// remove the two old circuits
				if j1InCircuit > j2InCircuit {
					circuitsArray = append(circuitsArray[:j1InCircuit], circuitsArray[j1InCircuit+1:]...)
					circuitsArray = append(circuitsArray[:j2InCircuit], circuitsArray[j2InCircuit+1:]...)
				} else {
					circuitsArray = append(circuitsArray[:j2InCircuit], circuitsArray[j2InCircuit+1:]...)
					circuitsArray = append(circuitsArray[:j1InCircuit], circuitsArray[j1InCircuit+1:]...)
				}
				// add the combined circuit
				circuitsArray = append(circuitsArray, combinedCircuit)
				connectionCount++
			} else {
				fmt.Println("Junctions already in the same circuit")
				connectionCount++
			}
		}

		fmt.Println("----------------------")
		if countTotalJunctionsInCircuits() == 1000 {
			fmt.Println("All junctions connected!")
			fmt.Println("lastpair:", jPair)
			break
		}

	}
	slices.Sort(circuitsArray)
	fmt.Println("Circuits len:", len(circuitsArray))
	for i, circuit := range circuitsArray {
		fmt.Printf("Circuit %d: , %d\n", i+1, len(strings.Split(circuit, "-")))
	}

}

func main() {
	file, err := os.Open("junctions")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("failed to close file: %s", err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		jPositions = append(jPositions, line)

	}
	fmt.Println("number of junctions:", len(jPositions))
	initDistances()
	setupCircuits()

}

// part 1:
// create an array of array of strings that represent the circuits
// keep this array sorted by the length of string arrays (shortest first)

// create circuits by going through the list of junction positions and comparing
// the distance between junctions

// after finding the 10 shortest connections, output the circuits
// mulitple the lengths of the three largest circuits and output the result
