My first try of using GoLang 

This program takes the test.csv file and looks for the Age field

bash-3.2$ cat test.csv\
John,Bishop,45,Leeds\
Bob,Smith,66,Bradford\
Joe,Blogs,22,London\
Mark,Jones,12,Paris\
Vicki,Parker,26,Scotland

to run use `go run file.go`

Results are :- 

bash-3.2$ go run file.go\
Outpuing the Age of all the users in the CSV file\
age 45\
age 66\
age 22\
age 12\
age 26\
