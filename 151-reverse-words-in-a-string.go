import "strings"

func reverseWords(s string) string {
    splits := strings.Split(s, " ")

    var arr []string
    for _, s := range splits {
        if len(s) == 0 {
            continue
        }
        arr = append(arr, s)
    }

    if len(arr) <= 1 {
        return strings.Join(arr, " ")
    }

    i, j := 0, len(arr)-1
    for i < j {
        arr[i], arr[j] = arr[j], arr[i]
        i++
        j--
    }

    return strings.Join(arr, " ")
}