package main

func main() {
	{
		values := buildIntArray("[5,4,3,2,1]")
		labels := buildIntArray("[1,1,2,2,3]")
		println(largestValsFromLabels(values, labels, 3, 1))
	}
	{
		values := buildIntArray("[5,4,3,2,1]")
		labels := buildIntArray("[1,3,3,3,2]")
		println(largestValsFromLabels(values, labels, 3, 2))
	}
	{
		values := buildIntArray("[9,8,8,7,6]")
		labels := buildIntArray("[0,0,0,1,1]")
		println(largestValsFromLabels(values, labels, 3, 1))
	}
	{
		values := buildIntArray("[3,2,3,2,1]")
		labels := buildIntArray("[1,0,2,2,1]")
		println(largestValsFromLabels(values, labels, 2, 1))
		printIntArray(values)
		printIntArray(labels)
	}

}
