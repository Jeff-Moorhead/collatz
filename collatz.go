package main

import ( // factorized import, good form
    "fmt"
    "os"
    "strconv"
    "time"
)


func next(num int) int {
    if num % 2 == 0 {
        return num / 2
    } else {
        return (3 * num) + 1
    }
}


func getNum(args []string) (int, int) { // args is a slice, as indicated by []string
    if len(args) < 2 {
        fmt.Println("Usage: collatz.go [number] --- Please enter a number for the Collatz sequence")
        return 0, 1
    } else if len(os.Args) > 2 {
        fmt.Println("Usage: collatz.go [number] --- Please only enter one number for the Collatz sequence")
        return 0, 2
    }

    num, err := strconv.Atoi(args[1])

    if err != nil {
        fmt.Println("Something went wrong: ", err)
        return 0, 3
    }
    return num, 0
}


func main() {
    num, status := getNum(os.Args)

    if status != 0 {
        os.Exit(0)
    }

    for num > 1 { // this is a golang while loop
        num = next(num) // get the next number in the sequence

        fmt.Println(num)
        time.Sleep(10 * time.Millisecond) // multiply by time.Second because time.Sleep takes nanoseconds as its argument
    }
}
