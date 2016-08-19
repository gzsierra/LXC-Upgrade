package main

import (
        "fmt"
        "sync"
        "github.com/gzsierra/go-lxc"
        "log"
        "gopkg.in/cheggaaa/pb.v1"
)

var (
      containers    []lxc.Container

      wg            sync.WaitGroup
      ctotal        int
      cdone         int

      cprogress     *pb.ProgressBar
      progressBars  *pb.ProgressBar
      pool          *pb.Pool
)

/*
 * From lxc host, we launch update in a specific container
 */
func upgradeVM(c lxc.Container)  {
  defer wg.Done()

  execute(c, "update")
  execute(c, "upgrade")
  execute(c, "clean")
  execute(c, "autoclean")

  cprogress.Increment()
  cdone++
}

func execute(c lxc.Container, arg string)  {
    progressBars.Increment()
    c.Execute("apt-get", "-y", arg)
}

/*
 * Launch routine to update containers
 */
func main() {
  // GET VM
  getC()
  // PROGRESSION
  progressionStart()
  // ROUTINE
  routine()
  // FINISHING
  progressionStop()
}

// GET CONTAINERS
func getC()  {
  containers = lxc.ActiveContainers("/var/lib/lxc")
  for i := range containers {
    log.Printf("%s\n", containers[i].Name())
  }
  // COUNT VM
  ctotal = len(containers)
}

// LAUNCH ROUTINE
func routine()  {
  wg.Add(ctotal)
  for _,element := range containers {
    go upgradeVM(element)
  }

  wg.Wait()
}

// START PROGRESSION BAR
func progressionStart()  {
  cprogress = pb.New(ctotal).Prefix("Containers")
  cprogress.ShowTimeLeft = false

  progressBars = pb.New(ctotal*4).Prefix("Tasks")
  progressBars.ShowTimeLeft = false

  pool, _ = pb.StartPool(cprogress, progressBars)
}

// STOP PROGRESSION BAR
func progressionStop()  {
  progressBars.Finish()
  cprogress.Finish()
  pool.Stop()
  fmt.Println("Work Done!")
}
