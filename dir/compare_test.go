package dir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	testDir1, _ := ReadDir("testdata/dir1")
	testDir2, _ := ReadDir("testdata/dir2")
	testDir3, _ := ReadDir("testdata/dir3")
	testDir4, _ := ReadDir("testdata/dir4")
	testDir5, _ := ReadDir("testdata/dir5")
	testDir6, _ := ReadDir("testdata/dir6")

	testCases := []struct {
		dir1, dir2        Dir
		expectedOnly1     []string
		expectedOnly2     []string
		expectedDifferent []string
	}{
		{testDir1, testDir1, []string{}, []string{}, []string{}},
		{testDir1, testDir2, []string{}, []string{}, []string{}},
		{testDir1, testDir3, []string{}, []string{}, []string{"file1.txt"}},
		{testDir1, testDir4, []string{"subdir", "subdir/file2.txt"}, []string{"subdirNew", "subdirNew/file2.txt"}, []string{}},
		{testDir1, testDir5, []string{}, []string{"subdirNew", "subdirNew/file3.txt"}, []string{}},
		{testDir1, testDir6, []string{"file1.txt", "subdir", "subdir/file2.txt"}, []string{"file2.txt"}, []string{}},
	}

	for _, tc := range testCases {
		t.Run("TestCompare", func(t *testing.T) {
			r := Compare(tc.dir1, tc.dir2)

			assert.ElementsMatch(t, tc.expectedOnly1, r.onlyDir1)
			assert.ElementsMatch(t, tc.expectedOnly2, r.onlyDir2)
			assert.ElementsMatch(t, tc.expectedDifferent, r.different)
		})
	}
}
