package sensupluginscpu

import (
	"fmt"
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
	IOWait  int64
	IRQ     int64
	SoftIRQ int64
	Steal   int64
	Guest   int64
}

// AcquireNumCPU returns the number of cores available to the system. This will include hyper-threaded cores.
func AcquireNumCPU() int {
	return runtime.NumCPU()
}

// ReadProc will fetch the desired cpu stats from the system
func ReadProc() string {
	filename := statFile

	// read in the file or puke to stdout and syslog
	usage, err := ioutil.ReadFile(filename)
	if err != nil {
		syslogLog.WithFields(logrus.Fields{
			"check":   "sensupluginscpu",
			"client":  host,
			"version": "foo",
			"error":   err,
		}).Error("ReadProc() could not read " + statFile)
		sensuutil.Exit("RUNTIMEERROR")
	}

	//Create empty slice of struct pointers.
	cpus := []*CPUData{}
	//
	//    //// Create another struct.
	//    //loc = new(Location)
	//    //loc.x = 5
	//    //loc.y = 8
	//    //loc.valid = true
	//    //
	//    //places = append(places, loc)
	//
	//    // Loop over all indexes in the slice.
	//    // ... Print all struct data.
	//    for i := range(places) {
	//	place := places[i]
	//	fmt.Println("Location:", place)
	//    }
	//}
	//
	//

  limitExp := regexp.MustCompile("[0-9]+")
  if strings.Contains(lines[i], "open files") {
    limits := limitExp.FindAllString(lines[i], 2)
  }
	// find the needed lines
	lines := strings.Split(string(usage), "\n")
	for i := range lines {
		if strings.Contains(lines[i], "cpu") {
			m := strings.Split(lines[i], " ")
			// Create struct and append it to the slice.
			cpu := new(CPUData)
			fmt.Println(m)
			cpu.CPU = m[0]
			cpu.User, _ = strconv.Atoi(m[1])
			cpu.Nice, _ = strconv.Atoi(m[2])

			cpus = append(cpus, cpu)
			//fmt.Print(m[0])
			//for j := range m {
			//	fmt.Println(m[j])
			//}
		}
	}
	// Loop over all indexes in the slice.
	// ... Print all struct data.
	for i := range cpus {
		cpu := cpus[i]
		fmt.Println("CPU:", cpu)
	}
	return "test"
}

//
//type CPUData struct {
//	CPU     string
//	User    int64
//	Nice    int64
//	System  int64
//	IOWait  int64
//	IRQ     int64
//	SoftIRQ int64
//	Steal   int64
//	Guest   int64
//}
