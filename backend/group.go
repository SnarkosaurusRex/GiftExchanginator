/* group.go declares and defines the Group type
 * Author: Lydia Holtrop
 * CS 214 final project, Spring 2017
 */

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*
 *   declare the Group type
 */
type Group struct {
	groupName    string
	groupMembers []Person
	memberNames  []string
}

/*
 * addMember() creates a new Person and adds them to the Group
 * Receive: aGroup, *Group
 *          aName, aSigOth, aRoomie, strings (the names of the new Person,
 *             their significant other, and their roommate)
 */
func addMember(aGroup *Group, aName, aSigOth, aRoomie string) {
	var (
		newb = Person{name: aName, sigOth: aSigOth, roomie: aRoomie}
	)
	aGroup.groupMembers = append(aGroup.groupMembers, newb)
	aGroup.memberNames = append(aGroup.memberNames, aName)
}

/*
 * removeByString() finds a given name in a given slice of strings
 *       and removes that name from the slice
 * Receive: slicedNames, a slice of strings (the slice to delete from)
 *          target, a string (the name to delete)
 * Return: the slice once the target name has been removed
 */
func removeByString(slicedNames []string, target string) []string {
	targetIndex := -1
	for i := range slicedNames {
		if slicedNames[i] == target {
			targetIndex = i
			break
		}
	}
	if targetIndex == -1 {
		return slicedNames
	} else {
		slicedNames = append(slicedNames[:targetIndex], slicedNames[targetIndex+1:]...)
	}
	return slicedNames
}

/*
 * findValidNames() fills a Person's validNames with the names that Person could be assigned
 * Receive: aPerson (*Person) - a member of the Group
 *          allTheNames - a slice of strings containing all the names in the Group
 *          noSigOths, noRoomies - bools representing the "rules" specified by the user
 */
func findValidNames(aPerson *Person, allTheNames []string, noSigOths, noRoomies bool) {
	aPerson.validNames = make([]string, len(allTheNames))
	copy(aPerson.validNames, allTheNames)
	// remove invalids
	// remove own name
	aPerson.validNames = removeByString(aPerson.validNames, aPerson.name)
	// remove others according to specified "rules"
	if noSigOths == true {
		aPerson.validNames = removeByString(aPerson.validNames, aPerson.sigOth)
	}
	if noRoomies == true {
		aPerson.validNames = removeByString(aPerson.validNames, aPerson.roomie)
	}
}

/*
 * assigninator() assigns each Person in the Group a random name from their validNames
 * Receive: aGroup (*Group)
 *          t (an int) - the number of trials that have been run already
 * Return: a string that indicates whether or not all of the names in the Group
 *             were assigned properly
 */
func assigninator(aGroup *Group, t int, writeFiles bool) string {
	//keep track of who's already been assigned
	assignedNames := []string{}

	for i := range aGroup.groupMembers {
		//use assignedNames to remove names from the Person's validNames
		for j := range assignedNames {
			aGroup.groupMembers[i].validNames = removeByString(aGroup.groupMembers[i].validNames, assignedNames[j])
		}
		if len(aGroup.groupMembers[i].validNames) == 0 {
			if t < 5 {
				assigninator(aGroup, t+1, writeFiles)
			} else {
				errorMessage := "Uh oh..." + aGroup.groupMembers[i].name + " doesn't have any valid names left to choose from...Try again!"
				return errorMessage
			}
		}

		//generate random # to be used as an index
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		x := r1.Intn(len(aGroup.groupMembers[i].validNames))

		//assign name at that index to person i
		aGroup.groupMembers[i].assignedName = aGroup.groupMembers[i].validNames[x]

		//remove that name from the Group's memberNames
		aGroup.memberNames = removeByString(aGroup.memberNames, aGroup.groupMembers[i].assignedName)
		//add that name to assignedNames
		assignedNames = append(assignedNames, aGroup.groupMembers[i].assignedName)
	}

	//check that it worked properly
	if len(aGroup.memberNames) == 0 {
		if writeFiles == true {
			writeResultsFiles(aGroup)
		}
		return "\nAll names have been assigned! Woot!"
	} else if t < 5 { //try assigninator up to 5 times if needed
		assigninator(aGroup, t+1, writeFiles)
	}
	return "\nWhoops, something didn't work quite right...Try again!"
}

/*
 *  writeResultsFiles() outputs each group member's result to a file that can then
 *       be emailed to that person as an attachment without anyone else seeing it
 *  Receive: aGroup (*Group)
 */
func writeResultsFiles(aGroup *Group) {
	for i := range aGroup.groupMembers {
		fileName := aGroup.groupName + "-" + aGroup.groupMembers[i].name + "Assignment" + ".txt"
		fout, _ := os.Create(fileName)
		fwriter := bufio.NewWriter(fout)

		stResult := ", you have been assigned " + aGroup.groupMembers[i].assignedName
		result := append([]byte(aGroup.groupMembers[i].name), stResult...)
		fwriter.Write(result)
		fwriter.Flush()
	}
	fmt.Println("Results files written!")
}

/*
 * doStuff() calls the necessary functions to run the whole program
 * Receive: aGroup (*Group)
 *          noSigOths, noRoomies, writeFiles - bools representing the "rules" specified by the user
 */
func doStuff(aGroup *Group, noSigOths, noRoomies, writeFiles bool) {
	for i := range aGroup.groupMembers {
		findValidNames(&aGroup.groupMembers[i], aGroup.memberNames, noSigOths, noRoomies)
	}
	var outcome = assigninator(aGroup, 0, writeFiles)
	fmt.Println(outcome)
}

// (the end)
