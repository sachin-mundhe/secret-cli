package main

import "cmds"

func main() {
	cmds.RootCmd.Execute()
}

//***********Encode - Convert MAP to JSON*************
// kvp := map[string]string{"name": "Searching..."}
// kvp["f1"] = "Rocks"
// var sb strings.Builder
// var sb bytes.Buffer
// enc := json.NewEncoder(&sb)
// enc.Encode(kvp)
// fmt.Println(sb)

// // *************Writing to file*************
// f, err := os.OpenFile("abc.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
// if err != nil {
// 	log.Fatal(err)
// }
// defer func() {
// 	err := f.Close()
// 	if err != nil {
// 		fmt.Println("Unable to close file...")
// 	} else {
// 		fmt.Println("File closed!")
// 	}
// }()

//***********Decode - Convert JSON to MAP*************
// file, _ := os.Open("abc.txt")
// kvp := make(map[string]string)
// dec := json.NewDecoder(file)
// err := dec.Decode(&kvp)
// if err != nil {
// 	log.Fatalln("Error occured while decoding. Error desc:", err)
// }
// fmt.Println(kvp)
