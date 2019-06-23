package main

type MountainArray []int

func (this *MountainArray) get(index int) int { return (*this)[index] }
func (this *MountainArray) length() int       { return len(*this) }

func main() {
	if true {
		nums := buildIntArray("[1,2,3,4,5,3,1]")
		mountain := MountainArray(nums)
		println(findInMountainArray(3, &mountain))
	}
	if true {
		nums := buildIntArray("[0,1,2,4,2,1]")
		mountain := MountainArray(nums)
		println(findInMountainArray(3, &mountain))
	}
	if true {
		nums := buildIntArray("[1,2,3,4,5,4,3,2,1]")
		mountain := MountainArray(nums)
		println(findInMountainArray(3, &mountain))
	}
	if true {
		nums := buildIntArray("[3,5,3,2,0]")
		mountain := MountainArray(nums)
		println(findInMountainArray(0, &mountain)) //
	}
}
