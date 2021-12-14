package compiler

func HashSame(hash1, hash2 []byte) bool {

	for i, hashData := range hash1 {
		if hashData != hash2[i] {
			return false
		}
	}

	return true
}
