package main

func checkSequential(num string) bool {
    if len(num) < 2 {
        return true
    }

    for i := 1; i < len(num); i++ {
        if num[i-1]+1 != num[i] {
            return false
        }
    }

    return true
}

func findDigits(low, high, curStr string, idx int, digits *[]int) {

}

func sequentialDigits(low int, high int) []int {
    result := make([]int, 0)
    return result
}
