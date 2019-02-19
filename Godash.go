package godash

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
