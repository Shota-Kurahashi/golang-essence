package paths

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func Main() {
	println(filepath.Base(`C:/path/to/file.txt`)) // file.txt

	println(filepath.Dir(`C:/path/to/file.txt`)) // C:\path\to

	println(filepath.Clean(`C:/path/to/..\file.txt`)) // C:\path\file.txt

	println(filepath.Ext(`C:/path/to/file.txt`)) // .txt

	println(filepath.IsAbs(`C:\path\to\file.txt`)) // true
	println(filepath.IsAbs(`.\file.txt`))          // false

	println(filepath.Join(`C:/path`, `to/file.txt`)) // C:\path\to\file.txt

	absolute, err := filepath.Abs(`../file.txt`)
	if err == nil {
		println(absolute) // カレントディレクトリがC:\path\toであればC:\path\file.txt
	}

	absolute, err = filepath.Rel(`C:\path`, `C:\path\to\file.txt`)
	if err == nil {
		println(absolute) // to\file.txt
	}

	println(filepath.VolumeName(`C:\path\to\file.txt`))     // C:
	println(filepath.VolumeName(`\\server\share\file.txt`)) // \\server\share

	s := `C:\path\to\file.txt`
	s = filepath.ToSlash(s)
	println(s) // C:/path/to/file.txt
	s = filepath.FromSlash(s)
	println(s) // C:\path\to\file.txt

	search()
}

func search() {
	files := []string{}

	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	err = filepath.WalkDir(cwd, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
