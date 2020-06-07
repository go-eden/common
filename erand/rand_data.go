package erand

// Bytes return [n]byte which contains random data.
func Bytes(n int) []byte {
	data := make([]byte, n, n)
	gr.Read(data)
	return data
}
