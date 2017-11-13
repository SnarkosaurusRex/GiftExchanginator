/* testinator.go contains all of the Gift Exchange-inator things
 * Author: Lydia Holtrop
 * CS 214 final project, Spring 2017
 */

package main

import (
	"fmt"
//   "List"
//	m "math"     // Math library with local alias m.
//	"strconv"    // String conversions.
)


/*
 *      PERSON
 */
	type Person struct {
		name, sigOth, roomie, assignedName string
		validNames []string
	}


/*
 *      GROUP
 */
   type Group struct {
      groupName    string
      groupMembers []Person
      memberNames  []string
   }


  // ADD MEMBER METHOD
   func addMember(aGroup *Group, aName, aSigOth, aRoomie string) {
      var (
         newb = Person{name: aName, sigOth: aSigOth, roomie: aRoomie}
      )
//      fmt.Println(newb)
      aGroup.groupMembers = append(aGroup.groupMembers, newb)
      aGroup.memberNames = append(aGroup.memberNames, aName)
   }


func main() {

  // TEST PERSON STUFF
   var (
      lya = Person{name: "Lya", sigOth: "Thomas", roomie: "Carolyn"}
   )

   fmt.Println(lya)
   fmt.Println(lya.sigOth)
   fmt.Println(lya.validNames)

   lya.validNames = append(lya.validNames, "DeAnna")
   fmt.Println(lya.validNames)

  // TEST GROUP STUFF
   var (
      testGroup = Group{groupName: "squad"}
   )

   fmt.Println(testGroup)
   fmt.Println(testGroup.groupName)
   fmt.Println(testGroup.groupMembers)
   fmt.Println(testGroup.memberNames)
   fmt.Println("------")
   testGroup.groupMembers = append(testGroup.groupMembers, lya)
   testGroup.memberNames = append(testGroup.memberNames, lya.name)
   fmt.Println(testGroup.groupMembers)

  // TEST addMember()
   fmt.Println("~~~~ Testing addMember() ~~~~")
   addMember(&testGroup, "Carolyn", "Ryan", "Lydia")
   fmt.Println(testGroup.groupMembers)
   fmt.Println(testGroup.memberNames)
}




