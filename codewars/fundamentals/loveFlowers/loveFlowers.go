package loveflowers

/**
If one of the flowers has an even number of petals
 and the other has an odd number of petals it means they are in love.
**/
func LoveFunc(flower1, flower2 int) bool {
	f1 := flower1 % 2
	f2 := flower2 % 2
	return f1 != f2
}
