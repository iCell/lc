type Info struct {
    Count int
    Index int
}

func firstUniqChar(s string) int {
    hash := make(map[rune]*Info, len(s))
    for idx, r := range s {
        i, ok := hash[r]
        if ok {
            i.Count += 1
        } else {
            hash[r] = &Info{
                Count: 1,
                Index: idx,
            }
        }
    }
    result := -1
    for _, value := range hash {
        if value.Count != 1 {
            continue
        }
        if result == -1 {
            result = value.Index
        } else if result > value.Index {
            result = value.Index
        }
    }

    return result
}