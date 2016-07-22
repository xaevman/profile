package profile

import (
    "bufio"
    "fmt"
    "os"
    "runtime/pprof"

    "github.com/xaevman/app"
)

var (
    profile   = false
    cpu       *os.File
    mem       *os.File
    cpuWriter *bufio.Writer
    memWriter *bufio.Writer
)

func Disable() {
    profile = false
}

func Enable() {
    profile = true
}

func End() {
    if !profile {
        return
    }

    pprof.WriteHeapProfile(memWriter)

    pprof.StopCPUProfile()
    cpuWriter.Flush()
    memWriter.Flush()

    cpu.Close()
    mem.Close()
}

func Start() {
    if !profile {
        return
    }

    var err error

    memPath := fmt.Sprintf("%s.mem", app.GetName())
    cpuPath := fmt.Sprintf("%s.cpu", app.GetName())

    fmt.Printf("Enabling profiler (heap: %s, cpu: %s)...\n", memPath, cpuPath)

    cpu, err = os.Create(cpuPath)
    if err != nil {
        panic(err)
    }

    mem, err = os.Create(memPath)
    if err != nil {
        panic(err)
    }

    cpuWriter = bufio.NewWriterSize(cpu, 64*1024)
    memWriter = bufio.NewWriterSize(mem, 64*1024)

    pprof.StartCPUProfile(cpuWriter)
}
