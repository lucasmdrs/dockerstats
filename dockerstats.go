// Package dockerstats provides the ability to get currently running Docker container statistics,
// including memory and CPU usage.
//
// To get the statistics of running Docker containers, you can use the `Current()` function:
//
// 		stats, err := dockerstats.Current()
//		if err != nil {
//			panic(err)
//		}
//
//		for _, s := range stats {
//			fmt.Println(s.ContainerID) // 9f2656020722
//			fmt.Println(s.Memory) // {Raw=221.7 MiB / 7.787 GiB, Percent=2.78%}
//			fmt.Println(s.CPU) // 99.79%
//		}
//
// Alternatively, you can use the `NewMonitor()` function to receive a constant stream of Docker container stats,
// available on the Monitor's `Stream` channel:
//
// 		m := dockerstats.NewMonitor()
//
// 		for res := range m.Stream {
//			if res.Error != nil {
//				panic(err)
//			}
//
//			for _, s := range res.Stats {
//				fmt.Println(s.ContainerID) // 9f2656020722
//			}
// 		}
package dockerstats

// Current returns the current `Stats` of each running Docker container.
//
// Current will always return a `[]Stats` slice equal in length to the number of
// running Docker containers, or an `error`. No error is returned if there are no
// running Docker containers, simply an empty slice.
func Current(comm Communicator, filters ...string) ([]Stats, error) {
	return comm.Stats(filters...)
}
