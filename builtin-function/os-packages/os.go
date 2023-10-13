package ospackages

import (
	"fmt"
	"os"
)

func Main() error {
	f, err := os.Open("file.txt")

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	defer f.Close()

	return nil
}
