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


  // runTests()

}


/*
 * TEST ALL THE THINGS!!!
 * (this function is what I used to test the program without
 *     re-entering all of the group members every time)
 */
func runTests() {
  // set-up
   var (
      testGroup = Group{groupName: "InsanitySquad"}
   )
   addMember(&testGroup, "Carolyn", "Ryan", "Lydia")
   addMember(&testGroup, "Lydia", "Thomas", "Carolyn")
   addMember(&testGroup, "Kelsey", "Sean", "Janelle")
   addMember(&testGroup, "Janelle", "Will", "Kelsey")
   addMember(&testGroup, "Kimberly", "", "DeAnna")
   addMember(&testGroup, "DeAnna", "", "Kimberly")
   addMember(&testGroup, "Sean", "Kelsey", "Walter")
   addMember(&testGroup, "Walter", "", "Sean")
   addMember(&testGroup, "Will", "Janelle", "")
   addMember(&testGroup, "Ryan", "Carolyn", "")
   addMember(&testGroup, "Noah", "", "")
   addMember(&testGroup, "Steve", "", "")


  // TEST RUN!!!
   doStuff(&testGroup, true, true, true)

/*  // print all results
   for i := range testGroup.groupMembers {
      fmt.Print(testGroup.groupMembers[i].name)
      fmt.Print(" --> ")
      fmt.Println(testGroup.groupMembers[i].assignedName)
   }*/

}


