### Tiny compiler 

a compiler for small language based on golang.

---

#### Syntax

##### var
var defines a variable
```
var name type = value
```
##### if / else if / else

```
if(boolean condition ){
    ... some statement
}else if(boolean condition / variable){
    ... some statement
}else {
    ... some statement
}

```

##### while / for
 - variable always number
``` 
   for(variable = value ; boolean condition; modify operation){
        ... some statement
   }
   
   while(boolean condition){
         ... some statenment 
   }
   

```

##### comments

```
// comment
```

##### arrays
```
var array []string = {"1","2","3","4"}
var array []num = {1,2,3,4}
var array []num = new([]num) // create empty array
    ...
var arrayLen num = array.len // get array length 
    ...
array.add(5) // add new element to array
```
##### functions

```
// starts from a-z and can include 0-9
func functionName(function args) return type or empty {
    ... some statements
    
    return value | return
}

```

##### types
```
// numbers(int)
var el1 num = 10
var el2 num = el + el
var el3 num = el - el
var el4 num = el * el
var el5 num = el / el
var el6 num = el % el

// strings
var s1 string = "string1"
var s2 string = "string2"
var s3 string = s1+s2
var b1 bool = s1 == s2
var b2 bool = s1 != s2

// logical
var t bool = true
var f bool = false
var r bool = t == f // false
var r bool = t != f // true

void / null 
```

##### console
```
var inputString string = input(text before input)
var inputNum num = input(text before input)
output(text to output)
```

-------

#### Simple example - console calculator

```
/* the function which starts application */
func main(args []string) void {
		
		var isWorking bool = true
		while(isWorking){
			var l num = input("please input a number : ")
			var op string = input("please input an operator like +,-,/,*,% :")
			var r num = input("please input a number :")	
			var res num = operation(op,l,r)
			output(makeStringResult(res))
		}
}

func operation(op string, l num, r num) num {
	
	var result num = 0
	
	if(op == "+"){
		result = l + r
	} else if(op == "-"){
		result = l - r
	}else if(op == "*"){
		result =l * r
	}else if(op == "/"){
		result = l / r
	}else if(op == "%"){
		result =l % r
	}else {
		output("please to ensure you peek correct operator")
		result = 0
	}
	
	return result
}

func makeStringResult(result num) string {
	var nums []num = revert(numToArray(result))
	var strNum string = ""
	
	for(i =0; i < nums.len; i = i+1){
		strNum = strNum + numToString(num[i])
	}
	
	if(strNum == ""){
		strNum = "undefined"
	}
	
	return "result is " + strNum
}

func numToString(number num) string {
	var numbers []string = {"0","1","2","3","4","5","6","7","8","9"}
	return numbers[number] 
}

func numToArray(number num) []num {
	var arrayOfNum []num = new([]num)
	var tempElement num = number
	while(tempelement > 0){
		arrayOfNum.add(tempElement % 10)
		tempElement = tempElement / 10
	}
	return arrayOfNum
}

func revert(array []num) []num {
	var revertArray []num = new([]num)
	var l num = array.len - 1
	for(i = 0; i < array.len; i = I+1){
		revertArray[i] = array[l-i]
	}
	
	return revertArray

}
```

