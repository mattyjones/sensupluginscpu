package sensupluginscpu

import (
	//"fmt"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"

	"github.com/yieldbot/sensuplugin/sensuutil"
	"regexp"
)

// CPUData holds discrete cpu values for a given physical or virtual cpu.
type CPUData struct {
	CPU     string
	User    int
	Nice    int
	System  int
	IOWait  int
	IRQ     int
	SoftIRQ int
	Steal   int
	Guest   int
}

// AcquireNumCPU returns the number of cores available to the system. This will include hyper-threaded cores.
func AcquireNumCPU() int {
	return runtime.NumCPU()
}

// ReadProc will fetch the desired cpu stats from the system
func ReadProc() []*CPUData {
	filename := statFile

	// read in the file or puke to stdout and syslog
	usage, err := ioutil.ReadFile(filename)
	if err != nil {
		syslogLog.WithFields(logrus.Fields{
			"check":   "sensupluginscpu",
			"client":  host,
			//"version": "foo",
			"error":   err,
		}).Error("ReadProc() could not read " + statFile)
		sensuutil.Exit("RUNTIMEERROR")
	}

	//Create empty slice of struct pointers.
	cpus := []*CPUData{}

	// find the needed lines
	lines := strings.Split(string(usage), "\n")
	for i := range lines {
		// exclude the total cpu line, it is easier to calculate the values yourself if needed
		matched, _ := regexp.MatchString("cpu[0-9]+", lines[i])
		if matched {
			m := strings.Split(lines[i], " ")
			// Create struct and append it to the slice.
			cpu := new(CPUData)
			cpu.CPU = m[0]
			cpu.User, _ = strconv.Atoi(m[1])
			cpu.Nice, _ = strconv.Atoi(m[2])
			cpu.System, _ = strconv.Atoi(m[3])
			cpu.IOWait, _ = strconv.Atoi(m[4])
			cpu.IRQ, _ = strconv.Atoi(m[5])
			cpu.SoftIRQ, _ = strconv.Atoi(m[6])
			cpu.Steal, _ = strconv.Atoi(m[7])
			cpu.Guest, _ = strconv.Atoi(m[8])

			cpus = append(cpus, cpu)
		}
	}
	return cpus
}
