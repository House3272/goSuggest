package sliceDeStrings

import (
	"fmt"
	//"reflect"
	"strings"
	"os"
	"bufio"
)

func NewSlice() *[]string {
	var mySlice []string
	return &mySlice
}

func MakeSlice(filePath *os.File, slicePtr *[]string) {

	r := bufio.NewScanner(filePath)
	for r.Scan() {
		*slicePtr = append(*slicePtr,r.Text())
	}
	if err := r.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done Making Slice Data")
}


// returns pointer to slice of matches
// number of results limited by rCount parameter
func GetMatches(queryS string, rCount uint64, sliceData *[]string) *[]string {
	var tempSlice []string
	for i := 0; ( i < len(*sliceData) )&&( uint64(len(tempSlice)) < rCount ); i++ {
		if strings.HasPrefix(strings.ToLower((*sliceData)[i]), queryS) {
			tempSlice = append(tempSlice,(*sliceData)[i])
		}
	}
	return &tempSlice
}