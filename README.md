# golang_longest_substring_without_repeating_characters

Given a string `s`, find the length of the **longest substring**
 without repeating characters.

## Examples

**Example 1:**

```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

```

**Example 2:**

```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

```

**Example 3:**

```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

```

**Constraints:**

- `0 <= s.length <= 5 * 104`
- `s` consists of English letters, digits, symbols and spaces.

## 解析

給定一個字串 s

要求寫一個演算法找出子字串中不包含重複字元的最長長度

最直接的想法是

當還沒遇到重複字元時就累加長度。遇到重複字元時，就從上一次開始累加的位置start 右移一位， 也就是start+1 當作下一個開始做累計

![](https://i.imgur.com/lQs2Dc0.png)

![](https://i.imgur.com/gitLval.png)

具體作法如下：

初始化最大長度 maxLen  = 0, start = 0, visitMap

然後逐步檢查每個字元 s[i]

每次先檢查 s[i] 是否已經存在 visitMap

如果s[i] 存在 visitMap 且 visitMap[s[i]] ≥ start 更新 start = vistMap[s[i]] +1

如果s[i] 不存在 visitMap ，更新 visitMap[s[i]] = i

更新 maxLen = max(maxLen, i - start +1)

當 start + maxLen ≥ len(s) 代表已經沒有辦法再找到更長不重複字元的子字串，直接回傳 maxLen

最後回傳 maxLen

![](https://i.imgur.com/nYohldn.png)

## 程式碼
```go
package sol

func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	start, maxLen := 0, 0
	visitMap := make(map[byte]int)
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for end := 0; end < sLen; end++ {
		// after start visit same character
		if pos, ok := visitMap[s[end]]; ok && pos >= start {
			start = pos + 1
		}
		visitMap[s[end]] = end
		maxLen = max(maxLen, end-start+1)
		if start+maxLen >= sLen {
			break
		}
	}
	return maxLen
}

```
## 困難點

1. 要看出找出不重複字元子字串每個位置間的關係
2. 要理解需要透過 hashTable 去儲存當下遇到每個字元的最新位置，用來做遇到重複字元開始位置更新

## Solve Point

- [x]  初始化 start = 0, maxLen = 0, visitMap
- [x]  從 i = 0.. len(s)-1 每次做以下運算
- [x]  檢查 s[i] 是否有在 visitMap 中
- [x]  如果有 , 檢查 visitMap[s[i]] ≥ start ， 更新 start =  visitMap[s[i]]+ 1
- [x]  更新 visitMap[s[i]] = i
- [x]  更新 maxLen = max(maxLen, i - start +1)
- [x]  當 start + maxLen ≥ len(s) 代表無法從 start 開始找到比 maxLen 長的不重複字元子字串 直接回傳 maxLen
- [x]  回傳 maxLen