package unix_path

func simplifyPath(path string) string {
	// 先处理 .
	// 规则:
	// 1. 只出现单次的.直接约去
	// 2. 出现两次的，向前找到非./的单词一起约去，找不到的话耶直接约去
	i := 0
	for i < len(path) {
		if path[i] != '.' {
			i++
			continue
		}
		j := i
		for j < len(path) && path[j] == '.' {
			j++
		}
		if j-i > 2 || (j < len(path) && path[j] != '/') {
			i = j
			continue
		}
		if i >= 0 && path[i-1] != '/' {
			i = j
			continue
		}
		if j-i == 1 {
			path = path[0:i] + path[i+1:]
		} else if j-i == 2 {
			path = path[0:i] + path[j:]
			kEnd := i - 1
			for kEnd > 0 && path[kEnd] == '/' {
				kEnd--
			}
			kBegin := kEnd
			for kBegin > 0 && path[kBegin] != '/' {
				kBegin--
			}
			path = path[0:kBegin+1] + path[kEnd+1:]
			i -= (kEnd - kBegin)
		}
	}
	// 再处理 /
	// 1. 重复的只保留一个
	i = 0
	for i < len(path) {
		if path[i] == '/' && i > 0 && path[i-1] == '/' {
			path = path[0:i] + path[i+1:]
		} else {
			i++
		}
	}
	// 2. 末尾的/去掉
	if len(path) > 1 && path[len(path)-1] == '/' {
		path = path[0 : len(path)-1]
	}
	// 3. 开头补上 /
	if len(path) == 0 {
		return "/"
	}
	if path[0] != '/' {
		path = "/" + path
	}
	return path
}
