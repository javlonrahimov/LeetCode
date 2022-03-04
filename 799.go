package main

import "math"

func chanpanageTower(poured, queryRow, queryGlass int) float64 {
    matrix := [102][102]float64{}
    matrix[0][0] = float64(poured)

    for r := 0; r <= queryRow; r++ {
        for c := 0; c <= r; c++ {
            q := (matrix[r][c] - 1.0) / 2.0
            if q > 0 {
                matrix[r+1][c] += q
                matrix[r+1][c+1] +=q;
            }
        }
    }

    return math.Min(1, matrix[queryRow][queryGlass])
}
