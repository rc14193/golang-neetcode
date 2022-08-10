import "strings"

func isAnagram(s string, t string) bool {
    if(len(s) != len(t)){
        return false
    }
    m := make(map[string]int)
    for _, value := range s{
        m[string(value)]++
    }
    for _, value := range t{
        count, ok := m[string(value)]
        if(ok == false){
            return false
        }
        count -= 1
        if (count == -1){
            return false
        }
        m[string(value)] = count
    }
    return true
}
