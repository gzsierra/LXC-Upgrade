package main

import (
        "fmt"
        "sync"
        "gopkg.in/lxc/go-lxc.v2"
        "log"
        "os/exec"
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
  cname := c.Name()

  execute(cname, "update")
  execute(cname, "upgrade")
  execute(cname, "clean")
  execute(cname, "autoclean")

  cprogress.Increment()
  cdone++
}

/*
 * Launch command
 */
func execute(cname string, action string){
  defer progressBars.Increment()

  cmd := "lxc-attach"
  cargs := []string{"-n", cname, "--", "apt-get", "-y", action}

  _, err := exec.Command(cmd, cargs...).Output()

  if err != nil {
    fmt.Println(err.Error())
  }
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
