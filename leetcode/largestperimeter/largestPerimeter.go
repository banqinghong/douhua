package largestperimeter

import "sort"

func largestPerimeter(A []int) int {
    sort.Ints(A)
    // [3,6,2,55,44,78,65,123]
    // sort.Sort(sort.Reverse(sort.IntSlice(A)))
    max := 0
    for i := len(A) - 1; i >= 2 ; i-- {
        if A[i - 2] <= A[i] - A[i - 1]{
            continue
        }else{
            max = A[i] + A[i - 1] + A[i - 2]
            return max
        }
    }
    return max
}
