/* menu.go provides interactive menus for the Gift Exchange-inator
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
 * menu1() provides an interactive menu from which the user can
 *    add members to the group, specify the "rules" they want to use,
 *    and run the Exchange-inator
 */
func menu1(zeGroup *Group) {
   reader := bufio.NewReader(os.Stdin)

   fmt.Println("\nOptions:")
   fmt.Println("1 - add a member to the group")
   fmt.Println("2 - run!")
   fmt.Print("Your choice: ")
   choice1, _ := reader.ReadString('\n')
   choice1 = strings.TrimSpace(choice1)

   if choice1 == "1" {
      fmt.Println("Please enter the following information for the new member.")
      fmt.Print("Name: ")
      aName, _ := reader.ReadString('\n')
      fmt.Print("Boyfriend/Girlfriend: ")
      aSigOth, _ := reader.ReadString('\n')
      fmt.Print("Roommate: ")
      aRoomie, _ := reader.ReadString('\n')
      //trim leading and trailing whitespace from inputs to avoid issues later
      aName = strings.TrimSpace(aName)
      aSigOth = strings.TrimSpace(aSigOth)
      aRoomie = strings.TrimSpace(aRoomie)

      addMember(zeGroup, aName, aSigOth, aRoomie)
      fmt.Println("Member added to the group!")
      menu1(zeGroup)
   } else if choice1 == "2" {
     //get "rules" from user
      var noBfGfs, noRoommates bool

      fmt.Print("\nAllow boyfriends/girlfriends? (y or n) ")
      choice2a, _ := reader.ReadString('\n')
      choice2a = strings.TrimSpace(choice2a)
      if choice2a == "y" {
         noBfGfs = false
      } else {
         noBfGfs = true
      }
      fmt.Print("\nAllow roommates? (y or n) ")
      choice2b, _ := reader.ReadString('\n')
      choice2b = strings.TrimSpace(choice2b)
      if choice2b == "y" {
         noRoommates = false
      } else {
         noRoommates = true
      }

     //run the exchange-inator
      doStuff(zeGroup, noBfGfs, noRoommates)
     //proceed to the results menu
      menu2(zeGroup)
   }
}


/*
 * menu2() provides an interactive menu from which each person can
 *    see their assigned name, but no one else's
 */
func menu2(zeGroup *Group) {
   reader := bufio.NewReader(os.Stdin)

   fmt.Print("\nEnter your name to see your result: ")
   usrName, _ := reader.ReadString('\n')
   usrName = strings.TrimSpace(usrName)

  //find that Person in groupMembers
   for i := range zeGroup.groupMembers {
      if zeGroup.groupMembers[i].name == usrName {
         fmt.Print("You've been assigned: ")
         fmt.Print(zeGroup.groupMembers[i].assignedName)
         break
      }
   }

   fmt.Println("\nWhat do you want to do next?")
   fmt.Println("1 - get another person's result")
   fmt.Println("2 - all done")
   fmt.Print("Your choice: ")
   choice3, _ := reader.ReadString('\n')
   choice3 = strings.TrimSpace(choice3)
   if choice3 == "1" {
     //"hide" result from next person
      for i := 0; i < 50; i++ {
         fmt.Println("    *    *    *    *    *    *    *    *    *    *")
      }
      menu2(zeGroup)
   } else {
     //"hide" last result
      for i := 0; i < 50; i++ {
         fmt.Println("    *    *    *    *    *    *    *    *    *    *")
      }
   }
}


