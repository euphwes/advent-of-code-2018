package mathUtils

// Takes an array of ints and returns their sum
func SumInts(input []int) int {
    sum := int(0)
    for _, n := range input {
        sum += n;
    }

    return sum
}

// Takes an array of floats and returns their sum
func SumFloats(input []float32) float32 {
    sum := float32(0)
    for _, n := range input {
        sum += n;
    }

    return sum
}