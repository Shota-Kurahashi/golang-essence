package hsd

func StringDistance(lhs, rls string) int {
	return Distance([]rune(lhs), []rune(rls))
}

func Distance(a, b []rune) int {
	dist := 0
	if len(a) != len(b) {
		return -1
	}

	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}

	return dist
}

func CreateProfile(filename string) (bool, error) {
	return false, nil
}

// ? 主なテストツールコマンド
//  $ go test -v
// 	$ go test -v -cover -coverprofile cover.tmp
//  $ grep -v _mock.go cover.tmp > cover.out
//  $ go tool cover -html=cover.out -o cover.html
