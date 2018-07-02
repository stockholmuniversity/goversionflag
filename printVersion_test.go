// Test with ldflags:
// go test -ldflags "-X suGoVersion.projectName=suGoVersion -X suGoVersion.gitCommit=gitCommit -X suGoVersion.jenkinsBuild=jBuild01 -X suGoVersion.buildTime=1970-01-01"
package suGoVersion

import (
	"testing"
)

var want = map[string]string{
	"projectName":  "suGoVersion",
	"gitCommit":    "gitCommit",
	"buildTime":    "1970-01-01",
	"jenkinsBuild": "jBuild01",
}
var version string

func TestPrintVersionAndExit(t *testing.T) {
	// Save current function and restore at the end:
	oldOsExit := osExit
	defer func() { osExit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	osExit = myExit
	fakeFlag = true
	// TODO check output from PrintVersionAndExit()
	PrintVersionAndExit()
	if exp := 0; got != exp {
		t.Errorf("Expected exit code: %d, got: %d", exp, got)
	}
}

func TestGetBuildInformation(t *testing.T) {
	got := GetBuildInformation()

	var key = []string{}
	for k := range want {
		key = append(key, k)
	}

	for _, k := range key {
		if want[k] != got[k] {
			t.Error("Expected", want[k], "got ", got[k])
		}
	}
	if len(got) != len(want) {
		t.Error("Number of elements differ. Expected", len(want), "got", len(got))
	}
}
