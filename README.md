### Tiny compiler 

a compiler for small language based on golang.

---

#### Syntax

##### var
var defines a variable
```
var name type = value
var name type = funcitonInvocation()

```
##### if / else if / else

```
if(boolean condition / variable){
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
/* comment */
```

##### arrays
```
var array []string = {"1","2","3","4"}
var array []num = {1,2,3,4}
var array []num = [10]num // create array with 10 elems
var array []num = [..]num // create empty array 
    ...
var arrayLen num = Len(array) // get array length 
    ...
Add(array,5) // add new element to array
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
var r bool = t > f // false
void / null 
```

##### console
```
var inputString string = Input(text before input)
var inputNum num = Input(text before input)
Output(text to output)
```

-------

#### Simple example - console calculator

```golang
/* the function which starts application */
func main(args []string) void {
		var isWorking bool = true
		while(isWorking){
			var l num = Input("please input a number : ")
			var op string = Input("please input an operator like +,-,/,*,% :")
			var r num = Input("please input a number :")	
			var res num = operation(op,l,r)
			Output(makeStringResult(res))
			var continue string = Input("do you want to continue?")
			if(continue != "y"){
				isWorking = false
			}
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
		Output("please to ensure you peek correct operator")
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
	var arrayOfNum []num = [..]num
	var tempElement num = number
	while(tempelement > 0){
		var nextEl num = tempElement % 10
		Add(arrayOfNum, nextEl)
		tempElement = tempElement / 10
	}
	return arrayOfNum
}

func revert(array []num) []num {
	var revertArray []num = [..]num
	var l num = Len(array) - 1
	for(i = 0; i < Len(array); i = i+1){
		revertArray[i] = array[l-i]
	}
	return revertArray
}
```

