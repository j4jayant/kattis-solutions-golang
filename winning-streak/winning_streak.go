package main

import (
    "fmt"
	//"bufio"
	"io"
	//"os"
	//"math"
	//"strings"
	//"strconv"
	//"log"
)



func main() {
	
	testCaseNo := 1
	
	
	for {
		//fmt.Printf("testCaseNo: %d\n", testCaseNo)
		
		if testCaseNo > 5 {
			break
		}
		var matchCount int64
		var winningProbability float64
		
        _, err := fmt.Scanf("%d%f\n", &matchCount, &winningProbability)
        if err == io.EOF {
			//fmt.Printf("err == io.EOF | break")
            break
        }
		if(matchCount == 0) {
			break
		}
		
		//fmt.Printf("matchCount: %d\n", matchCount)
		//fmt.Printf("winningProbability: %f\n", winningProbability)
		
		calculateLongestWinningStreak(matchCount, winningProbability)
		
		testCaseNo++
    }
	
}

func calculateLongestWinningStreak(matchCount int64, winningProbability float64) {
	lossProbability := 1 - winningProbability
	winProbabilities := make([]float64, matchCount + 1)
	winProbabilities[0] = 1

	var i, match, streak int64
	
	probLen := int64(len(winProbabilities))
	for i = 1; i < probLen; i++ {
		winProbabilities[i] = winProbabilities[i - 1] * winningProbability
	}

	
	streakByMatch := make([][]float64, matchCount + 1) 
	for i := range streakByMatch {
		streakByMatch[i] = make([]float64, matchCount + 1)
	}
	
	//fmt.Printf("streakByMatch: %v\n", streakByMatch)
	
	for match = 0; match < matchCount + 1; match++ {

		for streak = 0; streak < matchCount + 1; streak++ {
			if streak >= match {
				streakByMatch[match][streak] = 1
				continue
			}
			if streak + 1 == match {
				streakByMatch[match][streak] = 1 - winProbabilities[match]
				continue
			}
			streakByMatch[match][streak] = streakByMatch[match - 1][streak]

			if match - streak - 2 >= 0 {
				streakByMatch[match][streak] -= streakByMatch[match - streak - 2][streak] * lossProbability * winProbabilities[streak + 1]
			}
		}
	}

	//fmt.Printf("streakByMatch: %v\n", streakByMatch)
	
	var result float64
	result = 0.0
	//fmt.Printf("matchCount: %d\n", matchCount)
	for i = 1; i < matchCount + 1; i++ {
		result += float64(i) * (streakByMatch[matchCount][i] - streakByMatch[matchCount][i - 1])	
	}
	
	fmt.Printf("%f\n", result);
}
