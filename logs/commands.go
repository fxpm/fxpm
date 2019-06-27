package logs

import (
	"os"
	"strings"
	"time"

	"github.com/gernest/wow/spin"

	"github.com/gernest/wow"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/cheggaaa/pb.v1"
)

// CommandStarting is used to trigger the start of a command.
func CommandStarting(cmd *cobra.Command, args []string) {
	StartTime = time.Now()
	if Log == nil {
		cmd.Println("Logger is not ready, dumping RunPre/RunPost to console.")
		cmd.Printf("Starting command, %s", cmd.Name())

		return
	}

	Log.Info("Starting command", zap.String("command", cmd.Name()))
}

// CommandEnded is used to trigger the end of a command.
func CommandEnded(cmd *cobra.Command, args []string) {
	if Log == nil {
		cmd.Printf("Command completed in %s, %s", time.Since(StartTime).String(), cmd.Name())

		return
	}

	Log.Info("Command completed", zap.String("command", cmd.Name()), zap.String("took", time.Since(StartTime).String()))
}

// CommandProcess represents a single Process within a
// command.
type CommandProcess struct {
	Name        string
	Progress    *wow.Wow
	ProgressBar *pb.ProgressBar
	StartTime   time.Time
	EndTime     time.Time
	Duration    time.Duration
}

// NewCommandProcessFnFunc is the function type passed to the
// NewCommandProcessFn function.
type NewCommandProcessFnFunc func(p *CommandProcess)

// NewCommandProcess creates a new command process, starting
// the process timer.
func NewCommandProcess(args ...string) *CommandProcess {
	var message string
	if len(args) > 1 {
		message = " " + strings.Join(args[:1], " ")
	} else {
		message = " " + args[0]
	}

	p := &CommandProcess{
		Name:     args[0],
		Progress: wow.New(os.Stdout, spin.Get(spin.Dots3), message),
	}

	p.Start()

	return p
}

// NewCommandProcessProgress creates a new command process, starting
// the process bar for a process with max of count.
func NewCommandProcessProgress(name string, count int) *CommandProcess {
	p := &CommandProcess{
		Name:        name,
		ProgressBar: pb.New(count),
	}

	p.Start()

	return p
}

// NewCommandProcessFn provides a means of handling a
// process within the sub-process constructor.
func NewCommandProcessFn(name string, message string, fn NewCommandProcessFnFunc) {
	p := NewCommandProcess(name, message)
	fn(p)
}

// Spinner returns the instance of the wow.Wow.
func (process CommandProcess) Spinner() *wow.Wow {
	return process.Progress
}

// Bar returns the instance of the pb.ProgressBar.
func (process CommandProcess) Bar() *pb.ProgressBar {
	return process.ProgressBar
}

// Start is used to trigger the start of a the named sub-process.
func (process CommandProcess) Start() {
	process.StartTime = time.Now()

	if process.Progress != nil {
		process.Progress.Start()
	}

	if process.ProgressBar != nil {
		process.ProgressBar.Start()
	}

	Log.Info("Starting process", zap.String("process", process.Name))
}

// UpdateMessage is used to update the progress message for the sub-process.
func (process CommandProcess) UpdateMessage() {
	process.EndTime = time.Now()
	process.Duration = time.Since(process.StartTime)

	Log.Info("Completed process", zap.String("process", process.Name), zap.String("took", process.Duration.String()))
}

// Done is used to complete the named sub-process.
func (process CommandProcess) Done(args ...string) {
	process.EndTime = time.Now()
	process.Duration = time.Since(process.StartTime)

	if process.Progress != nil {
		process.Progress.PersistWith(spin.Spinner{Frames: []string{"✓"}}, " "+strings.Join(args, " "))
	}

	if process.ProgressBar != nil {
		process.ProgressBar.FinishPrint(strings.Join(args, " "))
	}

	Log.Info("Completed process", zap.String("process", process.Name), zap.String("took", process.Duration.String()))
}
