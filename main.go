package main

import (
	"fmt"
	"net/http"
)

// PUT taskItems_1 out of func main, to make it a global variable
// this is an array, can be defined in this way: var taskItems or alt way is taskItems :=
var taskItems_1 = []string{
	"Read KT again and again and do Go-lang",
	"Buy aloo 1kg and gobi 3 kg.",
	"Finally, award yourself with BOBS BAR pitcher!! :)",
}

var shortGolang = "Watch go crash course"
var fullGolang = "Watch NANAs go-crash course"
var rewardDesert = "Reward yourself with a donut u idiot sandwich "

func main() {
	// Doing the deployable part right here:
	fmt.Println("######## Welcome to the TODO list APP!! ########")

	// Register HTTP routes
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

	//// first way to declare variable
	//var taskOne = "Read KT again and again and do Go-lang"
	//var taskTwo = "Buy aloo 1kg and gobi 3 kg."
	//
	//// var maxItems = 20
	//
	//// 2nd way to declare variable, no need to use var, instead name :=
	//desert := "Finally, award yourself with BOBS BAR pitcher!! :)"
	//// no_of_tutorials := 5
	//
	//// I can simply print the above 4 var in a list or array -> go gives us arrays, slices
	//// slices are dynamically sized and more flexible than arrays. The taskItems is a slice
	//var taskItems = []string{taskOne, taskTwo, desert}
	//// printing the above slice
	//fmt.Println("Tasks:", taskItems)
	//
	//var taskItems_99 = []string{"one", "two", "three"}
	//
	//// FOR LOOPS:
	//// 1. basic for loop
	//// 2. for loop with condition
	//// 3. for range loops --> gives index along with the data. If index not needed, then type _ in place of it in for loop
	//for index, task := range taskItems {
	//	// fmt.Println(index, task)
	//	fmt.Printf("%d: %s\n", index+1, task)
	//}
	//
	//printTasks(taskItems_99) // This will now correctly print taskItems_99
	//fmt.Println()
	//taskItems_99 = addTask(taskItems_99, "GO for a run brother :/")
	//taskItems_99 = addTask(taskItems_99, "Practice straight drive")
	//
	//fmt.Println("Updated list")
	//printTasks(taskItems_99) // here we see that, there exists only 3 elements, and not 4 as all the variables in addTask func are local.
	//// fmt.Println("List of my TODO'S")
	//// fmt.Println(taskTwo)
	//// fmt.Println(taskOne)
	//// fmt.Println(desert)
	////
	//// fmt.Println("Tutorials")
	//// fmt.Println(taskOne, "no of tutorials =", no_of_tutorials)
	////
	//// fmt.Println("Reward")
	//// fmt.Println(desert)

}

// Handler function to show the task list
func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems_1 {
		// Write the response
		fmt.Fprintln(writer, task)
	}
}

// Handler function to greet the user
func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user! Welcome to our TODO list application!"
	fmt.Fprintln(writer, greeting)
}
