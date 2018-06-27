package main

type cache struct{
	string1 string
	string2 string
	int1 int
	int2 int
	limit int
}

var(
	caches = make(map[cache]*FizzBuzz)
)


/*
fizzbuzz memory system
*/
func (f *FizzBuzz)getCache()(ca cache){
	return cache{
		string1: f.string1,
		string2: f.string2,
		int1: f.int1,
		int2: f.int2,
		limit: f.limit,
	}
}
