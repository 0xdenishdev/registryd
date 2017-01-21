package main

import (
    "fmt"
    "time"

    manager "github.com/0xdenishdev/registryd/eventmanager"
)

func main() {
    sysMonitor := manager.Init()
    for {
        manager.Update(sysMonitor)
        fmt.Println("scanning...")
        time.Sleep(5 * time.Second)
    }
}
