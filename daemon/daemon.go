package daemon

import (
	// temporaily disabled for debugging
	"github.com/VividCortex/godaemon"
)

// TProcess defines the call signature of a process
// that can be daemonized
type TProcess func()

// Run runs the given process in a new forked process
func Run(proc TProcess) {
	godaemon.MakeDaemon(&godaemon.DaemonAttr{})
	proc()
}
