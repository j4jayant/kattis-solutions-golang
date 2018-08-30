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



type Pixel struct {
    C string
    Used bool
}

func main() {
	
	testCaseNo := 1
	
	
	for {
		//fmt.Printf("testCaseNo: %d\n", testCaseNo)
		
		if testCaseNo > 50 {
			
			break
		}
		var row, col int64
		
        _, err := fmt.Scanf("%d%d\n", &row, &col)
        if err == io.EOF {
			//fmt.Printf("err == io.EOF | break")
            break
        }
		if(row == 0) {
			break
		}
		//fmt.Printf("row: %d\n", row)
		//fmt.Printf("col: %d\n", col)
		
		var content [][]Pixel
		var i, j int64
		
		for i = 0; i < row; i++ {
			//fmt.Printf("i: %d\n", i)
			var rowData []byte
			
			_,err = fmt.Scanf("%s\n", &rowData)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
			
			var rowContent []Pixel
			for _, v := range rowData {
				pixel := Pixel{}
				
				pixel.C = fmt.Sprintf("%c", v)
				
				if pixel.C == "#" {
					pixel.Used = true
				} else {
					pixel.Used = false
				}
				
				rowContent = append(rowContent, pixel)
			}
			
			content = append(content, rowContent)
			
		}
		
		//fmt.Printf("content: \n%v\n", content)
		
		
		
		starOccurances := 0
		
		for i = 0; i < row; i++ {
			for j = 0; j < col; j++ {
				starOccurances += floodfill(row, col, i, j, content);
			}
		}
		
		
		fmt.Printf("Case %d: %d\n", testCaseNo, starOccurances)
		
		testCaseNo++
    }
	
}

func floodfill(row, col, i, j int64, content [][]Pixel) int {
    
    if(i < 0 || i >= row || j < 0 || j >= col) {
        return 0
    }

    
    if content[i][j].Used {
        return 0
    }

    
    content[i][j].Used = true

    
    floodfill(row, col, i, j+1, content)
    floodfill(row, col, i, j-1, content)
    floodfill(row, col, i+1, j, content)
    floodfill(row, col, i-1, j, content)

    return 1
}
