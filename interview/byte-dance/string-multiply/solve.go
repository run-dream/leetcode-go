package string_multiply

func multiply(num1 string, num2 string) string {
    if num1 == "0" ||  num2 == "0" {
        return "0"
    }
    l := len(num1)+len(num2)
    arr := make([]byte, l)
    for i := 0;  i < len(num1); i++{
        n1 := num1[len(num1)-1-i] - '0'
        idx := l-1-i
        var tmp byte
        for j := 0; j < len(num2); j++{
            n2 := num2[len(num2)-1-j] - '0'
            n := arr[idx] + n1*n2 + tmp
            arr[idx] = n%10
            tmp = n/10
            idx--
        }
        arr[idx] += tmp
    }
    var result string
    for i := 0; i <l; i++ {
        if arr[i] == 0 && len(result) == 0{
            continue
        }
        result += string('0'+arr[i])
    }
    return result
}
