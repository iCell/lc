func countAndSay(n int) string {
	result := "1"
	
	for i:= 1; i < n; i++ {
		var builder strings.Builder
		pre, count := result[0], 1

		for j := 1; j < len(result); j++ {
			if result[j] == pre {
				count++
				continue
			}
			
			builder.WriteString(strconv.Itoa(count))
			builder.WriteByte(pre)
			
			pre = result[j]
			count = 1
		}
		
		builder.WriteString(strconv.Itoa(count))
		builder.WriteByte(pre)
		
		result = builder.String()
	}
	
	return result
}