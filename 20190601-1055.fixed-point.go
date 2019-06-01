
package main;
func fixedPoint(A []int) int {
    for i, val := range(A) {
        if i == val {
            return i;
        }
    }
    return -1;
}