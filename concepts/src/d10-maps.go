package main


import "fmt"

func main() {
   
   // map literal  
    stockSymbols := map[string]string{
        "AAPL": "Apple",
        "MSFT":  "Microsoft",
        "FB": "Facebook",
        "ALPHABET": "Google",
    } 

    // Notice the order of printing 
    // fmt sorts the key in "ascending" order 
    //previously fmt used to print in random order
    fmt.Printf("\n%v\n", stockSymbols)


    // nil map.    
    var stockPrice map[string]float64

    //Comment this statement to show the error.
    stockPrice = make(map[string]float64)

    //run the program
    stockPrice["AAPL"] = 199.46

    fmt.Printf("\nThe stock price of APPLE is %.2f\n\n", stockPrice["AAPL"])

    stockPrice["ALPHABET"] = 1113.20

    stockPrice["FB"]= 189.53


    for k, v := range stockPrice {
        fmt.Printf("Stock: %s, Price: %.2f\n", k,v)

    }

    /* nil map, this does't need intializing using make
    as we are initializing with "another" map. */
   
    var stockData map[string]float64

    stockData = stockPrice

    //add another element
    stockData["MSFT"] = 136.95

    // let's print out both the maps 
    fmt.Println("\nMap: StockData")

    for k, v := range stockData {
        fmt.Printf("Stock: %s, Price: %.2f\n", k,v)

    }

    fmt.Println("\nMap: StockPrice")

    for k, v := range stockPrice {
        fmt.Printf("Stock: %s, Price: %.2f\n", k,v)

    }

    // Try inserting the same value again. 
    // duplicates are not allowed in "maps"
    // old value will be overriden by new value.

    stockData["MSFT"] = 140

    for k, v := range stockPrice {
        fmt.Printf("Stock: %s, Price: %.2f\n", k,v)

    }

    // delete 

    delete(stockData, "FB")

    fmt.Println("\n",stockData)

    // key doesn't exist 
    value, found := stockData["AMZN"]

    fmt.Println("\nIs Amazon present in stockdata map? ", found)
    fmt.Println("Default value of Key that is not present in Map is %f",
                 value)


    // When you try to delete value that does NOT exists.
    // Then, nothings happens. No error is thrown
    delete(stockData, "AMZN")


}