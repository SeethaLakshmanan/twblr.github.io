package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {
	result:= []int32{}
	for _, i := range vals{
		result = append(result, op(i))
	}
	return result

}

func filterInts(op filterOperation, vals []int32) []int32 {
	result:= []int32{}
	for _, i := range vals{
		if(op(i)){
			result = append(result, i)
		}
	}
	return result
}

func concatenate(dest []string, newValues ...string) []string {
	result:= []string{}
	for _, i := range newValues{
		result = append(result, i)
	}
	return result
}

func equals(list1 []string, list2 []string) bool {
	 for i := range list1 {
        if list1[i] != list2[i] {
            return false
        }
    }
    return true
}

func partialReverse(src []int, from, to int) []int {
	result:= []int{}
	for i:=to; i>=from; i--{
		result = append(result, src[i])
	}
	return result
}
