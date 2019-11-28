package toolkit

import (
	"fmt"
	"os"
	"os/exec"
)

type ExecOptions struct {
	Cwd    string
	Silent bool

	// /** optional envvar dictionary.  defaults to current process's env */
	// env?: {[key: string]: string}

	// /** optional out stream to use. Defaults to process.stdout */
	// outStream?: stream.Writable

	// /** optional err stream to use. Defaults to process.stderr */
	// errStream?: stream.Writable

	// /** optional. whether to skip quoting/escaping arguments if needed.  defaults to false. */
	// windowsVerbatimArguments?: boolean

	// /** optional.  whether to fail if output to stderr.  defaults to false */
	// failOnStdErr?: boolean

	// /** optional.  defaults to failing on non zero.  ignore will not fail leaving it up to the caller */
	// ignoreReturnCode?: boolean

	// /** optional. How long in ms to wait for STDIO streams to close after the exit event of the process before terminating. defaults to 10000 */
	// delay?: number

	// /** optional. Listeners for output. Callback functions that will be called on these events */
	// listeners?: {
	// stdout?: (data: Buffer) => void
	// stderr?: (data: Buffer) => void
	// stdline?: (data: string) => void
	// errline?: (data: string) => void
	// debug?: (data: string) => void
	// }
}

func Command(val string, args ...[]string) {
	cmd := exec.Command(val, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}

func AddFlagFromEnv(flags *[]string, flagFormat int, name, key string) {
	var format string
	switch flagFormat {
	case 1:
		format = "--%s=%s"
	case 2:
		format = "--%s %s"
	case 3:
		format = "-%s=%s"
	case 4:
		format = "-%s %s"
	default:
		panic(fmt.Sprintf("AddFlagByEnvVar{flagFormat} must be an int between 1-4, got %d", flagFormat))
	}

	if val, ok := os.LookupEnv(key); ok && val != "" {
		*flags = append(*flags, fmt.Sprintf(format, name, val))
	}
}
