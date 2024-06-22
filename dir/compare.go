package dir

func Compare(dir1 Dir, dir2 Dir) ComparisonResult {
	r := NewComparisonResult(dir1, dir2)

	for path, fileInfo1 := range dir1.Files {
		fileInfo2, ok := dir2.Files[path]
		if !ok {
			r.OnlyDir1(path)
			continue
		}

		// TODO: fileInfo.Mode should coincide as well

		if fileInfo1.Size() != fileInfo2.Size() && !fileInfo1.IsDir() && !fileInfo2.IsDir() {
			r.Different(path)
			continue
		}
	}

	for path := range dir2.Files {
		_, ok := dir1.Files[path]
		if !ok {
			r.OnlyDir2(path)
			continue
		}
	}

	return r
}
