package main

import "fmt"

func main()  {
	// Task 1: FooBar
	FooBar()
	
	// Task 2: Weather forecast
	weathers, err := weatherForecast()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Weather Forecast:")

	for _, w := range *weathers {
		fmt.Printf(
			"%s: %.2f°C\n",
			w.Date.Format("Mon, 02 Jan 2006"),
			w.Temp,
		)
	}
}