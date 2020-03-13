package main

// 31 tables, kitchen crew,
// x31 waiters (one alt)
// x10 kitchen crew
// x249 seated

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Person struct declaration:
type Person struct {
	Lastname  string
	Firstname string
	Table     string
}

type sendPerson struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Table int    `json:"Table"`
	IsKC  bool   `json:"IsKC"`
	IsW   bool   `json:"IsW"`
	Value int    `json:"Value"`
}

var i = 0

var toSend sendPerson

var sliceToSend []sendPerson

var stringToSend string

// Main function:
func main() {
	// Reading and parsing the original CSV file:
	csvFile, _ := os.Open("seating.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var people []Person
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		people = append(people, Person{
			Firstname: line[0],
			Lastname:  line[1],
		})
	}

	var peopleSlice []Person = people
	var slicedPeople = Shuffle(peopleSlice)
	initFile("first.csv")
	iterateAndChoose(slicedPeople, "first.csv")
	/*
		path := flag.String("path", "./first.csv", "Path of the file")
		flag.Parse()
		fileBytes, fileNPath := ReadCSV(path)
		SaveFile(fileBytes, fileNPath)
		fmt.Println(strings.Repeat("=", 10), "Done", strings.Repeat("=", 10))
	*/
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":6969", nil))
}

// Shuffle function taken from https://www.calhoun.io/how-to-shuffle-arrays-and-slices-in-go/
func Shuffle(slice []Person) []Person {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Person, len(slice))
	n := len(slice)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(slice))
		ret[i] = slice[randIndex]
		fmt.Println(slice[randIndex])
		slice = append(slice[:randIndex], slice[randIndex+1:]...)
	}
	return ret
}

// Chooses the next int people, appends to return crew:
func chooseNext(slice []Person, num int) []Person {
	crew := make([]Person, 0)
	for i := 0; i < num; i++ {
		crew = append(crew, slice[i])
	}
	return crew
}

// Chooses the next int people, removes them from the main slice:
func removeIndex(num int, slice []Person) []Person {
	for i := 0; i < num; i++ {
		slice = append(slice[:0], slice[0+1:]...)
	}
	return slice
}

// File append function, seatType is:
/*
1 = table
2 = KC
3 = waiter
*/
func makeFile(slice []Person, num int, seatType int, title string) {

	var name string

	f, err := os.OpenFile(title, os.O_APPEND|os.O_WRONLY, 0644)
	d := slice
	s := toSend
	for _, v := range d {
		v.Table = strconv.Itoa(num)
		var iS = strconv.Itoa(i)
		if seatType == 1 {
			fmt.Println("1")
			name = v.Lastname + " " + v.Firstname + "," + v.Table + "," + "false" + "," + "false" + "," + iS
			s.First = v.Lastname
			s.Last = v.Firstname
			s.Table = num
			s.IsW = false
			s.IsKC = false
			s.Value = i
		} else if seatType == 3 {
			//v.Table = "Waiter"
			fmt.Println("3")
			name = v.Lastname + " " + v.Firstname + "," + v.Table + "," + "false" + "," + "true" + "," + iS
			s.First = v.Lastname
			s.Last = v.Firstname
			s.Table = num
			s.IsW = true
			s.IsKC = false
			s.Value = i
		} else {
			fmt.Println("2")
			//v.Table = "KC"
			name = v.Lastname + " " + v.Firstname + "," + v.Table + "," + "true" + "," + "false" + "," + iS
			s.First = v.Lastname
			s.Last = v.Firstname
			s.Table = num
			s.IsW = false
			s.IsKC = true
			s.Value = i
		}
		sliceToSend = append(sliceToSend, s)
		fmt.Fprintln(f, name)
		i++
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	/*

	  res2D := &response2{
	        Page:   1,
	        Fruits: []string{"apple", "peach", "pear"}}
	    res2B, _ := json.Marshal(res2D)
	    fmt.Println(string(res2B))


	*/
	jsonToSend, _ := json.MarshalIndent(sliceToSend, "  ", "  ")
	//fmt.Println(string(jsonToSend))
	stringToSend = string(jsonToSend)
}

// Primary init of the file:
func initFile(title string) {
	f, err := os.Create(title)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
}

// Function to move a slice item:
func rearrange(remove int, place int, input []Person) []Person {
	slice := input
	val := slice[remove]
	slice = append(slice[:remove], slice[remove+1:]...)
	newSlice := make([]Person, place+1)
	copy(newSlice, slice[:place])
	newSlice[place] = val
	slice = append(newSlice, slice[place:]...)
	return slice
}

// Function that iterates through the slice and chooses people to go to certain positions based on index.
// Also calls in the file creation function.
func iterateAndChoose(slicedPeople []Person, title string) []Person {
	var originalGroup = slicedPeople

	// choose the first 10 to be kitchen crew:
	var nextGroup = chooseNext(slicedPeople, 10)

	makeFile(nextGroup, 10, 2, title)

	// remove the first 10 from the main list:
	removeIndex(10, slicedPeople)

	// choose the next 31 to be waiters:
	nextGroup = chooseNext(slicedPeople, 32)

	makeFile(nextGroup, 32, 3, title)

	// remove the next 31 from the main list:
	removeIndex(32, slicedPeople)

	// append all tables to CSV file:
	for i := 1; i < 32; i++ {
		var table = chooseNext(slicedPeople, 8)
		makeFile(table, i, 1, title)
		removeIndex(8, slicedPeople)
	}
	fmt.Println("all completed succesfully!")

	return originalGroup
}

func handler(w http.ResponseWriter, r *http.Request) {

	//toSend, err := os.Open("first.json")

	//file, _ := ioutil.ReadFile("first.json")

	//var toSend sendPerson

	//err := json.Unmarshal([]byte(file), &toSend)
	/*
		if err != nil {
			fmt.Println("error to send data")
			return
		}
	*/
	// Let's print the info
	fmt.Println("Incoming Request: ")
	fmt.Println("Method: ", r.Method, " ", r.URL)

	// Collect all the header keys
	headerKeys := make([]string, len(r.Header))
	i := 0
	for k := range r.Header {
		headerKeys[i] = k
		i++
	}
	//indent, _ := json.MarshalIndent(stringToSend, "", "")
	//"<prefix>", "<indent>"
	fmt.Fprintf(w, string(stringToSend))
	fmt.Println(string(stringToSend))
	fmt.Println("endjson")

	// Show Client Headers
	for _, line := range headerKeys {
		fmt.Println("  > ", line, ":", r.Header.Get(line))
	}
	//json.NewEncoder(w).Encode(toSend)

}
