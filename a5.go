// a5.go

//
// To run this, type the following at the command-line:
//
//		go run a5.go
//

package main

 	

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"errors"
	"reflect"
)

func main() {
}



func allBitSeqs(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1},{0}}
	}	

	lastSet := allBitSeqs(n-1)
	lastSet = append(lastSet, lastSet...) // this needs to be stored and not call twice to avoid double recursion
	// first loop adds a 1 to half the array, will always divide by 2 perfectly because of the doubling
	index := 0
	for index < len(lastSet)/2 {
		lastSet[index] = append(lastSet[index], 1)
		index++
	}
	// second loop adds a 0 to the other half of the array
	for index < len(lastSet) {
		lastSet[index] = append(lastSet[index], 0)
		index++
	}
	return lastSet
}

// Linear Searching 
func linearSearch(x interface{}, lst interface{}) int{
	// error checking on 
	xType := (reflect.TypeOf(x))
	lstType := (reflect.TypeOf(lst))
	if (lstType.Kind() != reflect.Slice) {
		panic("This is not a Slice")
	}
	if (xType != lstType.Elem()){
		panic("The Slice requires an item of a different type")
	}

	
	anInt, ok := x.(int) // assert that the actual type is int
	anIntArr, okA := lst.([]int)
	if ok && okA {
		return intSearch(int(anInt), []int(anIntArr))
	}
	aString, ok2 := x.(string) // assert that the actual type is string
	aStringArr, okA2 := lst.([]string)
	if ok2 && okA2{
		return stringSearch(string(aString), []string(aStringArr))
	}
	time, ok3 := x.(Time24) // assert that the actual type is Time24
	aTimeArr, okA3 := lst.([]Time24)
	if ok3 && okA3{
		return Time24Search(Time24(time), []Time24(aTimeArr))
	}

	return -1

}

func stringSearch(item string, list []string) int {
	i := 0
	for i < len(list) {
		if (list[i]==item){
			return i
		}
		i++
	}	
	return -1
}

func intSearch(item int, list []int) int {
	i := 0
	for i < len(list) {
		if (list[i]==item){
			return i
		}
		i++
	}	
	return -1
}

func Time24Search(item Time24, list []Time24) int {
	i := 0
	for i < len(list) {
		if (list[i]==item){
			return i
		}
		i++
	}	
	return -1
}

// time problems
type Time24 struct {
    hour, minute, second uint8
	
	
}
// 0 <= hour < 24
// 0 <= minute < 60
// 0 <= second < 60


func  equalsTime24(a Time24, b Time24) bool{
	return a.hour == b.hour && a.minute == b.minute && a.second == b.second
}

func  lessThanTime24(a Time24, b Time24) bool{
	return (a.hour < b.hour) || ((a.hour == b.hour) && (a.minute < b.minute)) || ((a.hour == b.hour) && (a.minute == b.minute) && (a.second < b.second))
}

func (t Time24) String () string {
	return (fmtint(t.hour))+":"+(fmtint(t.minute))+":"+(fmtint(t.second))
}

func fmtint (item uint8) string {
	intCastedItem := int(item)
	if intCastedItem < 10 {
		return "0"+(strconv.Itoa(intCastedItem))
	}
	return (strconv.Itoa(intCastedItem))
}

func (t Time24) validTime24() bool {
	return (t.hour < 24) && (t.minute < 60) && (t.second < 60)
}



func minTime24(times []Time24) (Time24, error) {

	if len(times) == 0 {
		return Time24{hour: 0, minute: 0, second: 0}, errors.New("There is no minimum time out of 0 times")
	}

	minTime := times[0]
	i := 1
	for i < len(times) {
		if lessThanTime24(times[i], minTime){
			minTime = times[i]
		}
		i ++
	}
	return minTime, nil
}



// count words, a much smaller function than the last one
func countWords(filename string) (map[string]int, error){
	data, err := ioutil.ReadFile(filename);
	if err != nil {
		return nil, err
	}
	
	// initialize the file data as an array and the map
	processedData := make(map[string]int)
	fileAsArray := strings.Fields(string(data))

	// fills in the map with words
	i := 0
	for i < len(fileAsArray){
		processedData[(fileAsArray[i])] = 0
		i++	
	}

	// counts the words updating the map as we go
	i = 0
	for i < len(fileAsArray){
		processedData[(fileAsArray[i])] = (processedData[(fileAsArray[i])] + 1)
		i++	
	}
	
	return processedData, err
}






// countEmirps and related helper functions
func countEmirpsLessThan(n int) int {
	numEmirps := 0
	i := 1
	condition := false
	for i < n {
		condition = isPrime(i) && isPrime(reverseInt(i)) && (i != reverseInt(i))
		if condition {
			numEmirps++
		}
		i++
	}
	return numEmirps;
}


func reverseInt(n int) int{
	length := numDigitsInt(n)
	i := 0
	reverse := 0

	for i < length {
		base := power (10, length - 1 - i)
		reverse = reverse + (n/base * power(10, i))
		n = (n % base)
		i++
	}
	return reverse
}

func numDigitsInt(n int) int{
	i:= 0
	for n > 0 {
		n = n/10
		i++
	}
	return i
}

func power(base int, exp int) int {
	i := 0
	returnable := 1
	for i < exp {
		i++
		returnable = base * returnable
	}
	return returnable
}

func isPrime(n int) bool {
	i := 2
	primeTest := true
	for i < (n-1) {
		primeTest = primeTest && (n % i != 0)
		i++
	}
	return primeTest
} 
