package main

import (
    "fmt"
    "time"
)
// START1 OMIT
func main() {
    go connectAndGetData("Server 1") // HL
    go connectAndGetData("Server 2") // HL
    go connectAndGetData("Server x") // HL
    go connectAndGetData("Server y") // HL
    time.Sleep(5 * time.Second)
    fmt.Println("Shutting down")
}

func connectAndGetData(endpoint string) { // HL
    fmt.Println("Connecting to", endpoint, "...")
    time.Sleep(1 * time.Second) //simulate response time
    fmt.Println("Connected to", endpoint)
    time.Sleep(1 * time.Second) //simulate response time
    fmt.Println("Getting data from", endpoint, "...")
    time.Sleep(1 * time.Second) //simulate getting slow data
    for i := 1; i < 100; i++ {
        fmt.Println("Downloading from: ", endpoint, "Percent done:", i)
    }
    fmt.Println("Done.  Closing connection with", endpoint)
} // HL

//END1 OMIT





