package diff_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/diff"
)

func ExampleDiff_unified() {
	// diff -u file1.txt file2.txt
	gloo.MustRun(
		Diff("testdata/file1.txt", "testdata/file2.txt", Unified),
	)
	// Output:
	//
}
