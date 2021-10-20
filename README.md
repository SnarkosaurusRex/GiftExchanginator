# The GiftExchanginator
## What is it?
The idea for this project stemmed from my uncanny tendency to draw my own name from the hat when drawing names for a Christmas gift exchange with friends. In order to keep this from happening, I decided to make a program to take the place of little pieces of paper in a literal hat. Assigning names to people is much faster and easier now. :)
(Oh, and the name "GiftExchanginator" is inspired by Dr. Heinz Doofenshmirtz, if you were wondering.)

## What does it do?
Currently, the GiftExchanginator has a very nice command line interface that first prompts the user for some information about the group. It then prompts for the necessary information for each group member. Once the information for each group member has been entered, options are provided that allow the user to specify whether or not to allow someone to be assigned their roommate and/or significant other. The program then runs, assigning a name to each member of the group. Once the program has finished running, each member of the group can enter their name and see which name they've been assigned (with lots of asterisks between each result so that each person only sees the name they've been assigned, keeping it a surprise for everyone else). Or, if the whole group isn't all together in the same place at the same time, the program also spits out nifty little .txt files that can be emailed out to each group member to inform them which name they've been assigned. This feature is also quite helpful in the event that one or more of the group members later forgets which name they got.

## Things that would be cool to do with/add to this:
* Automatically email result files to people (rather than just outputting the files and leaving them for the person running it to send out to the other people)
* Create an actual user interface that isn't just a command line thing
* Turn the whole thing into a web app
* Add some parallelization (it _is_ written in Go, after all)
* Come up with a way of detecting if no valid results are possible for the given group and alerting users accordingly
* Add an option to prevent "loops" (e.g. two people being assigned to each other)
