func groupAnagrams(strs []string) [][]string {
    compositionHashMap := map[[26]int][]string{}
    for _,str:=range strs{
        //1. generate composition key
        var composition [26]int
        for _,char:= range str{
            i:= char-'a'
            composition[i]++
        }
        //2. compares with key in hashmap
        //3. save in hashmap
        compositionHashMap[composition] = append(compositionHashMap[composition],str)
    }

    //4. transfer compositionHashMap into nested array
    //no order promise
    result := [][]string{}
    for _,group:=range compositionHashMap{
result = append(result,group)
    }
    return result
}
