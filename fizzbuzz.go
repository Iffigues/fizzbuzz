package main

import(
	"errors"
	"bytes"
	"strconv"
)

/*
Fizzbuzz Error detector
*/
func (f *FizzBuzz)Error()(err error){	
	if f.string1 == "" || f.string2 == "" {
		return errors.New("string1 or string2 can't be empty")
	}
	if (f.int1 * f.int2) == 0{
		return errors.New("int1 or int2 can't be egal to 0 ")
	}
	if f.limit == 0 {
		return errors.New("limit must be egal or superior to 1 And inferior or egal to "+strconv.Itoa(MaxInt))
	}
	return
}


/*
Return array of string formed following fizzbuzz logic
*/
func (f *FizzBuzz)MakeArray(){
	for i := 1;i<=f.limit;i++{
		var buff bytes.Buffer
		if i%f.int1 == 0{
			buff.WriteString(f.string1)
		}
		if i%f.int2 == 0{
			buff.WriteString(f.string2)
		}
		if buff.Len() == 0{
			buff.WriteString(strconv.Itoa(i))
		}
		f.Res = append(f.Res,buff.String())
	}
}

func (f *FizzBuzz)MakeArrays(){
	var buff bytes.Buffer
	for i := 1;i<=f.limit;i++{
		if i%f.int1 == 0{
			buff.WriteString(f.string1)
		}
		if i%f.int2 == 0{
			buff.WriteString(f.string2)
		}
		if buff.Len() == 0{
			buff.WriteString(strconv.Itoa(i))
		}
		f.Res = append(f.Res,buff.String())
		buff.Reset()
	}
}



/*
return string value from previous fizzbuzz array
*/
func (f *FizzBuzz)String()(final string,er error){
	if len(f.Res) != 0{
		var buff bytes.Buffer
		buff.WriteString(f.Res[0])
		for _,ok := range f.Res[1:]{
			buff.WriteString(",")
			buff.WriteString(ok)
		}
		return buff.String(),nil
	}
	return "",errors.New("array must be not empty")
}
