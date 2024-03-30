package main

import (
	"bufio"
	"fmt"
	"strings"

	//	"io"
	"os"
	//	"github.com/aws/aws-sdk-go/aws/defaults"
)

// Get input from the command line
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err:= r.ReadString('\n')

	return strings.TrimSpace(input), err
}

// create a file 
func createFile( fileName string){
	// create the file
	_, errs := os.Create(fileName + ".txt")
	if errs != nil {
		fmt.Println("Failed to create file:", errs)
		return
	 }

}

// append data to an existing file
func appendData(input string, fileName string ){
	 f, err := os.OpenFile(fileName + ".txt", os.O_APPEND|os.O_WRONLY, 0644)
	 if err != nil {
		fmt.Println("failed to append", err)
		return 
	 }
	 defer f.Close()

	 _, err = fmt.Fprintln(f, input)
	 if err != nil{
		fmt.Println(err)
		return 
	 }
	
	 fmt.Println("file appended successfully")
}


func taskManagerFunc(stringOpt string, r *bufio.Reader, name string){

	// use a switch statement to check if the user enters a valid option
		switch stringOpt {
		case "t":
			boolean := fileExists(name)
			// check if the fileExists function returns true
			// if not, create a file and append the tasks
			if !boolean{
				createFile(name)
				task, _ := getInput("Enter a task: ", r)
				// check if user enters blank
				if task == ""{
					fmt.Println("Here are your existing tasks: ")
					readLines(name)
					fmt.Println("Thank you for using Task Manager Hub!")
					return 
				}
				appendData(task, name)
				appendData(name, "TaskManagers")
			}

			// if it does, get the tasks and append it
			task, _ := getInput("Enter a task: ", r)
			// check if user enters blank
			if task == ""{
				fmt.Println("Here are your existing tasks: ")
				readLines(name)
				fmt.Println("Thank you for using Task Manager Hub!")
				return 
			}
			appendData(task, name)
			taskManagerFunc("t", r, name)

		case "d":
			taskManagers := readTaskFile("TaskManagers")

			present := contains(taskManagers, name)

			if present {
				readLines(name)
				taskD, _ := getInput("Please enter a task to delete: ", r)
				tasks := readTaskFile(name)

				contains := contains(tasks, taskD)
					if contains {
						err := os.Remove(name + ".txt")
						if err != nil {
							fmt.Println("Error removing file", err)
							return
						}

					remainingTasks := removeElementFromSlice(tasks, taskD)
					createFile(name)
						
					for _, value := range remainingTasks{
						appendData(value, name)
					}
					fmt.Println("File deleted Succesfully!")
					return 
					} 
			}
			taskManagerFunc("t", r, name)
		default:
			fmt.Println("Thank you for using Task Manager Hub!")
		}
	}

// function to remove an element from a slice
func removeElementFromSlice(arr []string, task string) []string {
    
    // initialize the index of the element to delete
    var index int
	
	for i := range arr {
		if arr[i] == task{
			index = i
		}
	}
    // check if the index is within slice bounds
    if index < 0 || index >= len(arr) {
        fmt.Println("The given index is out of bounds.")
		return []string{}
    } else {
        // delete an element from the slice
        newSlice := append(arr[:index], arr[index+1:]...)
       // fmt.Println("The new slice is:", newSlice)
	   return newSlice
    }  
}
	
// contains function
func contains(elems []string, v string) bool {
		for _, s := range elems {
			if v == s {
				return true
			}
		}
		return false
	}

// function to check if a file exists	
func fileExists(fileName string) bool{
	_, err := os.Stat(fileName + ".txt")

	if os.IsNotExist(err){
		return false
	}else if err != nil{
		return false
	}
	
	return true
}

// function that reads lines from a file and returns an array of strings
func readTaskFile(path string) []string {
	 readFile, err := os.Open(path + ".txt")
 
	 if err != nil {
		 fmt.Println(err)
	 }
	 fileScanner := bufio.NewScanner(readFile)
	 fileScanner.Split(bufio.ScanLines)
	 var fileLines []string
 
	 for fileScanner.Scan() {
		 fileLines = append(fileLines, fileScanner.Text())
	 }
    defer readFile.Close()
	 return fileLines
 }

 // function that reads lines from a file and prints them
func readLines(path string)  {
   // filePath := os.Args[1]
    readFile, err := os.Open(path + ".txt")

    if err != nil {
        fmt.Println(err)
		return
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string

    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

    for _, line := range fileLines {
      fmt.Println(line)
     }

	 defer readFile.Close()
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Welcome to Task Manager Hub. Create your tasks and save it with us!")
	opt, err := getInput("Please pick an option \n't' for creating tasks and \n'd' for deleting tasks: ", reader)
	if err != nil{
		fmt.Println("Error getting input")
	}
	name, _ := getInput("What is your name? ", reader)
			if name == ""{
				fmt.Println("Thank you for using Task Manager Hub!")
				return 
			}
	taskManagerFunc(opt, reader, name)
}

/*

/*



func update(x *string){
	 *x = "Tina"
}
name := "Kingsley"


	fmt.Println("Memmory address of name is ", &name)

	m := &name
	
	fmt.Println("value of m is ", *m )

	update(m)
	fmt.Println(name)
*/