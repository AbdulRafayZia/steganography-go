package main

import (
	"fmt"
	"github.com/DimitarPetrov/stegify/steg"
)

func main(){
	
	encodedFile := "encoded_girl.jpg" // Encoded File Name
	ResultFile := "horse.jpg"  // Desired Result File Name
	
	err := steg.DecodeByFileNames(encodedFile, ResultFile)
	if err != nil{
		fmt.Errorf("[-] Can't Decode The File")
	}else{
		fmt.Println("[*] File Has Been Succesfully Decoded")
	}
}
