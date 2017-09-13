# alGOtask
Algorithmic task for technical interview
###Description
Your program will be given a list of accepted words (vocabulary.txt). For each word W in the input file, you must find word W' from the list of accepted words such that the number of changes from W to W' is minimized. It is possible that W is already W' and thus the number of changes necessary is zero. A change is defined as replacing a single letter with another letter, adding a letter in any position, or removing a letter from any position. The total score that you need to output is the minimum number of changes necessary to make all words acceptable.
###Input Specification
Your program must take a single string argument, representing the file name of the content to analyze. The input file consists entirely of lower case letters and space characters. You are guaranteed that the input file will start with a lower case letter, and that all words are separated by at least one space character. The file may or may not end with a new line character.
Example input file:
tihs sententcnes iss nout varrry goud
You are guaranteed that your program will run against well formed input files and that the accepted word list is identical to the one provided for testing.
###Output Specification
Your program must print out the minimum number of changes necessary to turn all words in the input file into accepted words as defined by the vocabulary. Words may not be joined together, or separated into multiple words. A change in a word is defined as one of the following: Replacing any single letter with another letter. Adding a single letter in any position. Removing any single letter. This score must be printed out as an integer and followed by a single new line.
Example Command:
go run example.go example_input
Example Output (newline after number):
8
###Programming Language and Efficiency
Your solution has to be fast in order for us to consider you for this role. You can test your solution with file 187 to make sure that: a) your solution is correct (it should output 187) b) it’s fast enough (it should run under 2 seconds in Go on your regular machine).

