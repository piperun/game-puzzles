package main


import (
        "fmt"
)


func SolvePuzzle(PuzzleSum int){

        product := 0
        leftover := 0
        mod := PuzzleSum % 9
        if mod == 0 {
                product = (PuzzleSum / 9) - 1
                leftover = PuzzleSum - (product * 9)
        } else {
                product = (PuzzleSum  - mod) / 9
                leftover = mod
        }
        fmt.Printf("Input: %d nine(9) times and then %d\n", product, leftover)

}

func main() {
        fmt.Println("Choose the puzzle to solve:")
        fmt.Println(`
        1: 36
        2: 67
        3: 81
        4: custom value
        `)

        var choices int
        var custom_val int

        fmt.Scan(&choices)

        switch choices {
        case 1:
                SolvePuzzle(36)
        case 2:
                SolvePuzzle(67)
        case 3:
                SolvePuzzle(81)
        case 4:
                fmt.Print("Insert custom value: ")
                fmt.Scan(&custom_val)
                SolvePuzzle(custom_val)
        default:
                fmt.Println("Please input a valid number")
        }

}
