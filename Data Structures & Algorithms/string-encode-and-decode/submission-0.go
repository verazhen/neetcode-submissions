type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    result:=""
    for _,str:=range strs{
        lstr := fmt.Sprintf("%03d", len(str))
        result= fmt.Sprintf("%s%s%s",result,lstr,str)
    }
    return result
}

func (s *Solution) Decode(encoded string) []string {
    i:=0
    results:=[]string{}
    for {
        if i >= len(encoded){
            break
        }
        lstr:=encoded[i:i+3]
        l, err := strconv.Atoi(lstr)
        if err!=nil {
            fmt.Println(err)
            return []string{}
            }
        str := encoded[i+3:i+3+l]
        results = append(results,str)
        i+=3+l
    }
    return results
}
