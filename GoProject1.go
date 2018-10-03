//Armaly Albert
/*
Correct input file is war.txt
Correct output file is warOutput.txt
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

//Preconditions: Is ran by main
//Postconditions: Returns the name of the desired text file the user wants to open
func getFile() string{
	fmt.Println("Please enter the name of the text file.")
	var fileName string
	fmt.Scanln(&fileName)

	//sets user entry to fileName and returns
	return fileName
}
//Preconditions: Is ran by main
//Postconditions: Returns a map with the word list and corresponding count
func readFile(fileName string) map[string]int{

	//Creates a map to hold word list and count
	wordList := make(map[string]int)

	//Reads in the file and sets it to the rawData string
	byteArray, err := ioutil.ReadFile(fileName)  //This method was taken from the slides you gave us

	//Error catcher, spits out file not read message and exit program
	if err != nil{

		fmt.Println("File not read")
		os.Exit(1)
	}

	fmt.Println("Removing all non alpha numerics")
	editedArray := ""

	//Regex case, looks for A-z both upper and lower and also includes letter with accent marks. Also keeps numbers and white space
	regexArray, err := regexp.Compile("[^A-z\\sÀ-ÿ0-9]")

	editedArray += regexArray.ReplaceAllString(string(byteArray), "")//Replaces any that are not in the regex case with a "" to omit it
	fmt.Println("Done")
	fmt.Println(editedArray)
	fmt.Println("Counting word usage")

	//Separates the words by white spaces and places them in array
	wordCount := strings.Fields(editedArray)

	//Adds each word to the map or increments by 1 if it is already there
	for j:=0; j<len(wordCount); j++ {
		wordList[wordCount[j]]++
	}

	fmt.Println("Word count:")

	//Creation of output file https://golangcode.com/writing-to-file/
	file, err := os.Create("warOutput.txt")
	if err != nil {
		log.Fatal("Output file error", err)

	}
	defer file.Close();

	fmt.Fprint(file, wordList)
	return wordList
}
func main(){
	//Correct file is war.txt
	var fileName string = getFile()
	//fmt.Println(fileName)

	fmt.Println(readFile(fileName))

	fmt.Println("Word count was also printed to warOutput.txt")
}