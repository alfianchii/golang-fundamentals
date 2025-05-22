package main

import "fmt"

func main() {
	// The slice method
	names := [...]string{"Taka", "Ogi", "Kanji", "Nugraha", "Moe"}
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
	days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
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
}