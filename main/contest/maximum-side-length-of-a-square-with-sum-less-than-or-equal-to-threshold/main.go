package main

import "LeetCodeGo/utils"

func maxSideLength(mat [][]int, threshold int) int {
    result := 0
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[0]); j++ {
            if j != 0 {
                mat[i][j] += mat[i][j-1]
            }

            ll, maxLen := 0, utils.Min(i, j)+1

            for ll < maxLen {
                area := 0
                for k := 0; k < ll+1; k++ {
                    prefix := 0
                    if j-ll-1 > 0 {
                        prefix = mat[i-k][j-ll-1]
                    }
                    area += mat[i-k][j] - prefix
                }
                if area > threshold {
                    break
                }
                ll++
            }
            result = utils.Max(ll, result)
        }
    }
    return result
}
