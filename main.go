/* main.go acts as a driver for the Gift Exchange-inator
 * Author: Lydia Holtrop
 * CS 214 final project, Spring 2017
 */

package main

import (
	"fmt"
   "bufio"
   "os"
   "strings"
)

/*
 * This main function gets some initial information from the user to
 *    establish a new Group, then calls the first interactive menu
 */
func main() {
   reader := bufio.NewReader(os.Stdin)

   fmt.Println("\nWelcome to the Gift Exchange-inator!")
   fmt.Print("Enter a group name: ")
   gpName, _ := reader.ReadString('\n')
   gpName = strings.TrimSpace(gpName)
   var (
      theGroup = Group{groupName: gpName}
   )
   menu1(&theGroup)
}

