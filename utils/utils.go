package utils

//RemoveUlidDuplicate 重複を削除
func RemoveUlidDuplicate(args []string) []string {
	results := make([]string, 0, len(args))
	uniqueEncoutered := map[string]bool{}
	for i := 0; i < len(args); i++ {
		if !uniqueEncoutered[args[i]] {
			uniqueEncoutered[args[i]] = true
			results = append(results, args[i])
		}
	}
	return results
}

//GetUlidDuplicateargs 重複した値のみ取得
func GetUlidDuplicateargs(args []string) []string {
	results := make([]string, 0, len(args))
	uniqueEncoutered := map[string]bool{}
	for i := 0; i < len(args); i++ {
		if !uniqueEncoutered[args[i]] {
			uniqueEncoutered[args[i]] = true
		} else {
			results = append(results, args[i])
		}
	}
	return results
}
