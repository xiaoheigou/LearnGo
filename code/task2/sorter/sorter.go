package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"./algorithms/bubblesort"
	"./algorithms/qsort"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile=", *infile, "outfile=", *outfile, "algorithm=", *algorithm)
	}
	values, err := readValues(*infile)
	if err != nil {
		fmt.Println(err)
	}
	t1 := time.Now()
	switch *algorithm {
	case "quicksort":
		qsort.QuickSort(values)
	case "bubblesort":
		bubblesort.BubbleSort(values)
	default:
		fmt.Println("Sorting algorithm", *algorithm, "is either unknown orunsupported.")
	}
	t2 := time.Now()
	fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")

	writeValues(values, *outfile)
}

// 逐行读入
func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Fail to open the input file")
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	//定义int数组，用来存放从文件中读取的待排序的数据
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		// 转换字符数组为字符串
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)

	}
	return

}

//逐行输出
func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}

	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
