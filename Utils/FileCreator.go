//package FileCreator
//
//import (
//	"fmt"
//	"os"
//)

//import (
//	"fmt"
//	"io/ioutil"
//)
//
//// create main function to execute the program

package Utils

import (
	"fmt"
	"io/ioutil"
)

// create main function to execute the program
func main() {
	err := ioutil.WriteFile("myfile.py", []byte("print('hello World')"), 0644) //create a new file
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File is created successfully.") //print the success on the console
}
