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
    time.Sleep(5 * time.Second) // OMIT
    //Wait for all data to be returned and do something
    //...
    fmt.Println("Shutting down")
}
//END1 OMIT
//START2 OMIT
func connectAndGetData(endpoint string) { // HL
    // Simulate a slow data download or connection
    // ...
    fmt.Println("Connecting to", endpoint, "...") // OMIT
    time.Sleep(1 * time.Second) //simulate response time // OMIT
    fmt.Println("Connected to", endpoint) // OMIT
    time.Sleep(1 * time.Second) //simulate response time // OMIT
    fmt.Println("Getting data from", endpoint, "...") // OMIT
    time.Sleep(1 * time.Second) //simulate getting slow data // OMIT
    for i := 1; i < 100; i++ { // OMIT
        fmt.Println("Downloading from: ", endpoint, "Percent done:", i) // OMIT
    } // OMIT
    fmt.Println("Done.  Closing connection with", endpoint) 
} // HL

//END2 OMIT





