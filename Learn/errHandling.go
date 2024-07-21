// package main

// import (
// 	"fmt"
// 	"os"
// )

// // kill the server in killServer(pidFile String)
// // open file to read server process ID
// // covert pid from str to int
// // simulate kill
// // log the kills

// func readPidFile(pidFile string) ([]string, error) {

// 	file, err := os.Open(pidFile)
// 	if err != nil {
// 		return nil, fmt.Errorf("not able to open the file")
// 	}
// 	defer file.Close()

// 	// data := make([]string, len(file))
// 	// tmp := ""
// 	// for _, val := range file {

// 	// 	if string(val) == "\n" && len(tmp) != 0 {
// 	// 		data = append(data, string(tmp))
// 	// 		tmp = ""
// 	// 	} else if len(string(val)) != 0 {
// 	// 		tmp += string(val)
// 	// 	}
// 	// }

// 	// fmt.Println(data[0])
// }

// func killServer(pidFile string) error {

// 	// open file

// 	return nil
// }

// func main() {
// 	pidFile := "./servers.txt"
// 	file, err := readPidFile(pidFile)
// 	if err != nil {
// 		fmt.Println("ERROR: ", err)
// 	}
// 	fmt.Println(file)
// }
