func simplifyPath(path string) string {
	var ans []string
	components := strings.Split(path, "/")
	for _, component := range components {
		if len(component) == 0 || component == "." {
			continue
		}     
		if component == ".." {
			if len(ans) != 0 {
				ans = ans[:len(ans)-1]   
			}
			continue
		}
		ans = append(ans, component)
	}
	return "/" + strings.Join(ans, "/")
}