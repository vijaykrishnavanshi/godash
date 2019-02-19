package godash

import "reflect"

// Chunk just takes the input array and returns the chunked arrays
func Chunk(input []interface{}, chunkLength int) [][]interface{} {
	lengthOfArray := len(input)
	var chunks [][]interface{}
	for i := 0; i < lengthOfArray; i += chunkLength {
		end := i + chunkLength
		if end > lengthOfArray {
			end = lengthOfArray
		}
		chunks = append(chunks, input[i:end])
	}
	return chunks
}

// IsFalsy just takes the input and return a boolean if the value is Falsy
func IsFalsy(input interface{}) bool {
	if input == nil {
		return true
	}
	return reflect.DeepEqual(input, reflect.Zero(reflect.TypeOf(input)).Interface())
}

// Compact just takes the input array and returns the array without falsy value
func Compact(input []interface{}) []interface{} {
	lengthOfArray := len(input)
	var chunks []interface{}
	for i := 0; i < lengthOfArray; i++ {
		if IsFalsy(input[i]) {
			chunks = append(chunks, input[i])
		}
	}
	return chunks
}
