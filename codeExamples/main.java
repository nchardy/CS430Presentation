package com.company;

public class Main {
    // START1 OMIT
    public static void main(String[] args) {
        new Thread(() -> { connectAndGetData("Server 1"); }).start();
        new Thread(() -> { connectAndGetData("Server 2"); }).start();
        new Thread(() -> { connectAndGetData("Server x"); }).start();
        new Thread(() -> { connectAndGetData("Server y"); }).start();
        
        try {
            Thread.sleep(5000); // OMIT
        } catch(InterruptedException ex) {
            //Handle Exception
        }
        System.out.println("Shutting down");
    }
    // END1 OMIT

    public static void connectAndGetData(String endpoint) {
        try {
            System.out.println("Connecting to " + endpoint + "...");
            Thread.sleep(1000); //simulate response time
            System.out.println("Connected to " + endpoint);
            Thread.sleep(1000); //simulate response time
            System.out.println("Getting data from " + endpoint + "...");
            Thread.sleep(1000); //simulate getting slow data
            for(int i = 1; i < 100; i++){
                System.out.println("Downloading from: " + endpoint + " Percent done: " + i);
            }
            System.out.println("Done.  Closing connection with" + endpoint);
        } catch(InterruptedException ex) {
            //Handle Exception
        }

    }

    // START2 OMIT
    public static void connectAndGetData(String endpoint) {
        ...

    }
    // END2 OMIT
}

	