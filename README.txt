FizzBuzz README

FizzBuzz is a REST server.


We have two urls :

URL:
	"http://localhost:9001/":
		execute fizzbuzz with:
    	 	     string1 = "Fizz"
	 	     string2 = "Buzz"
	 	     int1 = 3
	 	     int2 = 5
		     limit = 100

	"http://localhost:9001/<sting1>/<string2>/<int1>/<int2>/<limit>"
		execute fizzbuzz with:
		     string1 = <string1>
		     string2 = <string2>
		     int1 = <int1>
		     int2 = <int2>
		     limit = <limit>



string1 and string2 are simple strings

int1, int2, limit are integer representations


command server:
	"--help" :
		 Give info about fizzbuzz server command

	"--string1":
		Update string1

	"--string2":
		Update string2

	"--int1":
		Update int1

	"--int2":
		Update int2

	"--limit":
		Update limit

	"--socket":
		Update socket /decimal representation number

	"--command":
		Open bash command for change parameter server
