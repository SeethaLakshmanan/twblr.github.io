package data_structures

func findPeopleWithCommonInterest(data map[string][]string, interest string) []string {
	result:=[]string{}
	for key, value := range data{
		if(contains(value,interest)){
			result = append(result, key)
		}
	}
	return result
}

func contains(src []string, elem string) bool {
	var response= false
	for _,i:= range src{
		if(i==elem){
			response=true
			break
		}
	}
	return response
}
