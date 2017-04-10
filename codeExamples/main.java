package com.company;

public class Main {
    // START1 OMIT
    public static void main(String[] args) {
        new Thread(() -> { connectAndGetData("Server 1"); }).start(); // HL
        new Thread(() -> { connectAndGetData("Server 2"); }).start(); // HL
        new Thread(() -> { connectAndGetData("Server x"); }).start(); // HL
        new Thread(() -> { connectAndGetData("Server y"); }).start(); // HL
        //Wait for all data to be returned and do something
        //...
        try { // OMIT
            Thread.sleep(5000); // OMIT
        } catch(InterruptedException ex) { // OMIT
            //Handle Exception // OMIT
        } // OMIT
        System.out.println("Shutting down");
    }
    // END1 OMIT
    // START2 OMIT
    public static void connectAndGetData(String endpoint) {
        // Simulate a slow data download or connection
        // ...
        try { // OMIT
            System.out.println("Connecting to " + endpoint + "..."); // OMIT
            Thread.sleep(1000); //simulate response time // OMIT
            System.out.println("Connected to " + endpoint); // OMIT
            Thread.sleep(1000); //simulate response time // OMIT
            System.out.println("Getting data from " + endpoint + "..."); // OMIT
            Thread.sleep(1000); //simulate getting slow data // OMIT
            for(int i = 1; i < 100; i++){ // OMIT
                System.out.println("Downloading from: " + endpoint + " Percent done: " + i); // OMIT
            } // OMIT
            System.out.println("Done.  Closing connection with" + endpoint);
        } catch(InterruptedException ex) { // OMIT
            //Handle Exception // OMIT
        } // OMIT
    }
    // END2 OMIT
}

	