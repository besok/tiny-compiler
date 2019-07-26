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

	for(i =0; i < Len(nums); i = i+1){
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