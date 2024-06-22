package dir

import (
	"bytes"
	"path/filepath"
)

func NewComparisonResult(dir1, dir2 Dir) ComparisonResult {
	return ComparisonResult{
		dir1: dir1,
		dir2: dir2,
	}
}

type ComparisonResult struct {
	onlyDir1, onlyDir2, different []string
	dir1, dir2                    Dir
}

func (r *ComparisonResult) OnlyDir1(path string) {
	r.onlyDir1 = append(r.onlyDir1, path)
}

func (r *ComparisonResult) OnlyDir2(path string) {
	r.onlyDir2 = append(r.onlyDir2, path)
}

func (r *ComparisonResult) Different(path string) {
	r.different = append(r.different, path)
}

func (r ComparisonResult) String() string {
	if len(r.onlyDir1)+len(r.onlyDir2)+len(r.different) == 0 {
		return "Directories coincide"
	}

	var buf bytes.Buffer

	// only path1
	for i, p := range r.onlyDir1 {
		if i == 0 {
			buf.WriteString("only in " + r.dir1.Path + ":\n")
		}
		buf.WriteString(p)
		if i < len(r.onlyDir1)-1 {
			buf.WriteString("\n")
		}
	}
	if len(r.onlyDir1) > 0 && len(r.onlyDir2) > 0 {
		buf.WriteString("\n\n")
	}

	// only path2
	for i, p := range r.onlyDir2 {
		if i == 0 {
			buf.WriteString("only in " + r.dir2.Path + ":\n")
		}
		buf.WriteString(p)
		if i < len(r.onlyDir2)-1 {
			buf.WriteString("\n")
		}
	}
	if len(r.onlyDir1)+len(r.onlyDir2) > 0 && len(r.different) > 0 {
		buf.WriteString("\n\n")
	}

	// different
	for i, p := range r.different {
		if i == 0 {
			buf.WriteString("files that differ:\n")
		}
		buf.WriteString(filepath.Join(r.dir1.Path, p) + " and " + filepath.Join(r.dir2.Path, p))
		if i < len(r.different)-1 {
			buf.WriteString("\n")
		}
	}

	return buf.String()
}
