package diff_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/diff"
)

func ExampleDiff_basic() {
	// diff file1.txt file2.txt
	// Note: This requires actual files to compare
	gloo.MustRun(
		Diff("testdata/file1.txt", "testdata/file2.txt"),
	)
	// Output:
	//
}
