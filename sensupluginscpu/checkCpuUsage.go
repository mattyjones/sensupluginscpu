// Copyright © 2016 Yieldbot <devops@yieldbot.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package sensupluginscpu

import (
	"fmt"

	"github.com/spf13/cobra"
	"time"
)

// checkCPUUsageCmd represents the checkCpuUsage command
var checkCPUUsageCmd = &cobra.Command{
	Use:   "checkCpuUsage",
	Short: "Check the total cpu usage",
	Long:  `Check the total cpu usage on a machine. This will take into account the specific geometry of the given machine.`,
	Run: func(cmd *cobra.Command, args []string) {

		val1 := ReadProc()
		time.Sleep(5)
		val2 := ReadProc()


		fmt.Println(val2[1].User)
    //total := (val2.User + val2.System + val2.Nice + val2.SoftIRQ + val2.Steal) / (val2.User + val2.System + val2.Nice + val2.SoftIRQ + val2.Idle + val2.IOWait)
		//fmt.Println("total: ", total)
		for i := range val1 {
			cpu := val1[i]
			fmt.Println("CPU:", cpu)
		}
		for i := range val2 {
			cpu := val2[i]
			fmt.Println("CPU:", cpu)
		}
	},
}

//CPU_Util = (user+system+nice+softirq+steal)/(user+system+nice+softirq+steal+idle+iowait)

func init() {
	RootCmd.AddCommand(checkCPUUsageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCpuUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCpuUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
