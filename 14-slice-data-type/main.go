package main

import "fmt"

func main() {
	// The Slice Method
	names := [...]string{"Taka", "Ogi", "Kanji", "Nugraha", "Moe"} // This is an array
	firstSlice := names[1:3]
	secondSlice := names[3:]
	thirdSlice := names[:3]
	fourthSlice := names[:]

	fmt.Println("The slice:", names)
	fmt.Println("----------------")
	fmt.Println("First slice:", firstSlice)
	fmt.Println("Second slice:", secondSlice)
	fmt.Println("Third slice:", thirdSlice)
	fmt.Println("Fourth slice:", fourthSlice)
	fmt.Println("================================================================================================================")

	// Funcs
	days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"} // This is an array
	fmt.Println("Days:", days)
	fmt.Println("----------------")

	workDays := days[:5]
	workDaysLength := len(workDays)
	workDaysCapacity := cap(workDays)
	fmt.Println("Work days:", workDays)
	fmt.Println("Work days length:", workDaysLength)
	fmt.Println("Work days capacity:", workDaysCapacity)
	fmt.Println("----------------")

	holiday1 := "Holiday 1"
	holiday2 := "Holiday 2"
	workDays[0] = holiday1
	workDays[1] = holiday2
	fmt.Println("Replace some old data to 'workDays':", workDays) // Changed
	fmt.Println("The 'days' data:", days) // Also affected into the parent
	fmt.Println("----------------")

	workDay1 := "Workday 1"
	workDay2 := "Workday 2"
	/*
		append() also added the data to its parent, 
		but start from the end of the sliced data's index (the reference).
		WOULDN'T add the data into the reference (workDays)
	*/
	weekend := append(workDays, workDay1)
	/* 
		If append tried to add items beyond of days' capacity, 5 + 3 = 8
		but the cap is 7, then append would make a new array, copy all
		elements of workDays + new elements, then returns as a new slice.
		Anyway, the weekend's cap is doubled based on the parent's cap.
		Parent's cap: 7
		Weekend's cap: 7x2 = 14
	*/
	// weekend := append(workDays, workDay1, workDay1, workDay1)
	fmt.Println("Weekend:", weekend)
	fmt.Println("Workdays:", workDays)
	fmt.Println("Days:", days)
	// And affected into the reference and parent too
	fmt.Println("----------------CHANGE THE weekend[0]!----------------")
	weekend[0] = workDay2
	fmt.Println("Weekend:", weekend)
	fmt.Println("Workdays:", workDays)
	fmt.Println("Days:", days)
	fmt.Println("================================================================================================================")

	// Slice Init
	var myName []string = make([]string, 2, 5) // Data type, length, capacity
	myName[0] = "Alfian"
	myName[1] = "Taka"
	// Got error
	// myName[2] = "Ganteng"
	fmt.Println("My name:", myName)
	fmt.Println("Slice's length:", len(myName))
	fmt.Println("Slice's capacity:", cap(myName))
	fmt.Println("----------------")

	/*
		If we append items which exceeded the 
		parent's capacity (5), then it would
		make the new array and any changes
		made in the sliced (myFullName) array wouldn't
		affect the parent (myName)
	*/ 
	// myFullName := append(myName, "Ganteng", "Maksimal") // Change to Chii, affect parent
	myFullName := append(myName, "Ganteng", "Maksimal", "Parah", "Sekali", "Lagi") // Change to Chii, doesn't affect parent (create new array)
	fmt.Println("My name:", myFullName)
	fmt.Println("Slice's length:", len(myFullName))
	fmt.Println("Slice's capacity:", cap(myFullName))

	myFullName[1] = "Chii"
	fmt.Println("My name:", myName)
	fmt.Println("My full name:", myFullName)
	fmt.Println("================================================================================================================")

	// Slice Copy
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
	fmt.Println("================================================================================================================")

	// Array vs. Slice
	array := [...]int{1, 2, 3, 4, 5} // Fixed size: 5
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("This is ARRAY:", array)
	fmt.Println("This is SLICE:", slice)
	fmt.Println("----------------")

	fmt.Println("Array's length:", len(array))
	fmt.Println("Array's capacity:", cap(array))
	fmt.Println("Slice's capacity:", cap(slice))
	fmt.Println("Slice's capacity:", cap(slice))
	fmt.Println("================================================================================================================")

	// Furthermore Explanation
	// Got error
	// newArray := append(array, 1)
	/*
		This would create a new slice from 'array'
		var that have a pointer to the 'array' var
	*/ 
	newArray := array[:]
	newArray[0] = 99
	fmt.Println(array) // Affected
	fmt.Println(newArray)

	// Got error: the 'array' var still array data type
	// appendedData := append(array, 1)
	appendedData := append(newArray, 1)
	fmt.Println("Appended data:", appendedData)
	fmt.Println("----------------")

	fmt.Println("Array's length:", len(array))
	fmt.Println("Array's capacity:", cap(array))
	fmt.Println("New Array's length:", len(newArray))
	fmt.Println("New Array's length:", cap(newArray))
	fmt.Println("Appended Data's capacity:", len(appendedData))
	fmt.Println("Appended Data's capacity:", cap(appendedData))
	fmt.Println("----------------")

	fmt.Println("----------------CHANGE THE appendedData[0]!----------------")
	/*
		Doesn't affect its reference and 
		its parent since the append() method
		would create a new ... (array or slice?)
		First, append() allocates a new array
		(if there's no capacity left).
		Then copies the prior elements.
		Afterward, returns a new slice
		(pointing to the new array).
	*/ 
	appendedData[0] = 100
	fmt.Println("Array:", array)
	fmt.Println("New Array:", newArray)
	fmt.Println("Appended Data:", appendedData)
}