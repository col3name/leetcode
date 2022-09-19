package _609

import "strings"

// 609. Find Duplicate File in System
// efgh: root/a2.txt, root/c/d/4.txt, root/4.txt
// abcd: root/a/1.txt, root/c/3.txt

func findDuplicate(paths []string) [][]string {
	mp := make(map[string][]string, 1000)

	var fileContent string
	var data []string
	var fileData []string
	builder := strings.Builder{}
	for _, path := range paths {
		fileData = strings.Split(path, " ")
		for _, file := range fileData[1:] {
			data = strings.Split(file, "(")
			fileContent = data[1][:len(data[1])-1]
			if mp[fileContent] == nil {
				mp[fileContent] = make([]string, 0, 10)
			}
			builder.WriteString(fileData[0])
			builder.WriteString("/")
			builder.WriteString(data[0])
			mp[fileContent] = append(mp[fileContent], builder.String())
			builder.Reset()
		}
	}
	result := make([][]string, 0, len(mp))
	for _, files := range mp {
		if len(files) > 1 {
			result = append(result, files)
		}
	}
	return result
}
