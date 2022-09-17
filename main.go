package main

import (
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func readData(s string) (*os.File, error) {
	f, err := os.Open(s)
	if err != nil {
		return nil, fmt.Errorf("Wasn't able to read data file")
	}

	return f, nil
}

func dropColumns(df dataframe.DataFrame) {
	newdf := df.Drop([]int{0, 1, 2, 3, 4})

	fmt.Println(newdf)
}

//INPUT CSV
//RETURN AVAILABLE COLUMN TO COMPARE
//INPUT COLUMN TO COMPARE
//RETURN DATAFRAM COMPARING THESE COLUMN

func main() {
	f, _ := readData("./data/data.csv")

	df := dataframe.ReadCSV(f)

	dropColumns(df)
}
