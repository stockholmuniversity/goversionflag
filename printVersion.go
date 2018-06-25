/*
Package suversion gives the compiled binary a "--version" argument
if compiled with correct -ldflags
(The variables must be replaced with actual variables from the build system)

 go build -ldflags "-X suversion.projectName=$project -X suversion.gitCommit=$commit -X suversion.buildTime=$time -X jenkinsBuild=$jenkins"
Primary usage of this packet should be using the function PrintVersionAndExit() which handles
"--version" and "-version" according to GNU Coding standard 4.7.1:
https://www.gnu.org/prep/standards/html_node/_002d_002dversion.html#g_t_002d_002dversion

If the user of "suversion" also use the package "flag" for its own flags, the flags must be declared before PrintVersionAndExit() is called.

There is also the possibility to call GetBuildInformation() to just get the build information without printing or exit the program.
*/
package suversion

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

// Var that is set at compile time by the build system.
// Build arguments:
// go build -ldflags "-X suversion.projectName=$project -X suversion.gitCommit=$var"
var (
	projectName  string
	gitCommit    string
	buildTime    string
	jenkinsBuild string
)

// TODO: Fix comment
var osExit = os.Exit

var fakeFlag = false

/*
PrintVersionAndExit prints information of the build and exits

According to GNU coding standard 4.7.1 argument --version to the binary should:
  "The standard --version option should direct the program to print information about its
  name, version, origin and legal status, all on standard output, and then exit successfully
  Other options and arguments should be ignored once this is seen, and the program should not perform its normal function."
https://www.gnu.org/prep/standards/html_node/_002d_002dversion.html#g_t_002d_002dversion
*/
func PrintVersionAndExit() {

	arg := flag.Bool("version", false, "Prints build information and exits.")
	if flag.Parsed() == false {
		flag.Parse()
	}

    // TODO, also iterate over leftover arguments in flag.Args() if version is there
	if *arg == true || fakeFlag == true {
		buildversion := GetBuildInformation()
		buildSlice := []string{}
		for k, v := range buildversion {
			buildSlice = append(buildSlice, k+": "+v)
		}
		sort.Strings(buildSlice)
		for _, v := range buildSlice {
			fmt.Println(v)
		}

		osExit(0)
	}
}

// GetBuildInformation returns a map with build information from Jenkins at compile time.
//
// In most cases you should not use this function but instead PrintVersionAndExit().
// This function assume that the main program implements its own variation of PrintVersionAndExit()
func GetBuildInformation() (buildversion map[string]string) {
	buildversion = map[string]string{
		"projectName":  projectName,
		"gitCommit":    gitCommit,
		"buildTime":    buildTime,
		"jenkinsBuild": jenkinsBuild,
	}
	return buildversion
}
