package memory

import (
	"fmt"
	"testing"
)

func Test_test(t *testing.T) {

	initMemory(1)
	pB := putBool(true)
	fmt.Printf("bool : %+v \n", pB)
	pI := putInt(1024 * 1024 * 1024)
	fmt.Printf("int : %+v \n", pI)
	pS := putString("Boris, Hello World!")
	fmt.Printf("string : %+v \n", pS)


	fmt.Printf("string : %s \n", getGeneric(pS))
	fmt.Printf("int: %d \n", getGeneric(pI))
	fmt.Printf("bool : %v \n", getGeneric(pB))

	aB := []bool{true,true,true,false,false}
	aI := []int{1,2,3,4,5,6}
	aS := []string{"a","boris","!!!"}

	pAB := putArrayBool(aB)
	pAI := putArrayInt(aI)
	pAS := putArrayString(aS)

	arr1 := getArray(pAB)
	arr2 := getArray(pAI)
	arr3 := getArray(pAS)

	fmt.Println(arr1,arr2,arr3)
	fmt.Println("offset", offset)


}
