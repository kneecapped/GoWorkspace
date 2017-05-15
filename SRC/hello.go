/* author: Kyle Kubis
 */
package main

import ("fmt"
	"io/ioutil"
)

func main() {
	var test string = "Test"
	fmt.Println("Hello, world")
	fmt.Println(test)
	printtest()
	var file_data string = read_file("SRC\\testfile.txt")
	fmt.Println(file_data)
	var kyle person = person{"Kyle",26}
	PrintPerson(kyle)
}

func printtest()  {
	fmt.Print("This came from printest")
}

func check(e error){
	if e != nil {
		panic(e)
	}
}

func read_file(filename string)(string){
	dat, err := ioutil.ReadFile(filename)
	check(err)
	return string(dat)
}

func int_demo(){
	var myint32 int32
	myint32 = 32
	fmt.Println(myint32)
	var my16int1,my16int2 int16
	my16int1 = 5
	my16int2 = 10
	fmt.Println(my16int1+my16int2)
	// Unsigned ints
	var unsigned8 uint8 = 255
	var unsigned16 uint16 = 65535
	fmt.Println(unsigned8)
	fmt.Println(unsigned16)

}

type person struct {
	name string
	age uint8
}

func PrintPerson (h person) {
	fmt.Println("Name:",h.name,"Age: ",h.age)
}
