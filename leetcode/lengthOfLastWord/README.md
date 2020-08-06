## 最后一个单词的长度
链接：https://leetcode-cn.com/problems/length-of-last-word/

## 思路

从后往前循环字符串，当遇到第一个非空格字符的时候空格标志(spaceFlag)加1，否则不做操作继续下一次循环。
每次循环的时候对当前字符以及空格标志(spaceFlag)进行判断，
如果当前字符是空格字符，存在下面两种情况

- 空格标志(spaceFlag)大于0，那么说明在两次空格字符串之间存在非空格字符，且长度是spaceFlag，那么这个spaceFlag则是我们需要求得的最后一个单词的长度。
- 空格标志(spaceFlag)等于0，那么说明我们还没有找到我们需要的单词，单词的长度依旧为0

举例说明：
s := "hello  world   " (最后面两个空格)，计当前字符串为str，长度为spaceFlag

- 第一次循环 str为空格 spaceFlag = 0
- 第二次循环 str为空格 spaceFlag = 0
- 第三次循环 str为字符"d" spaceFlag = 0 然后 spaceFlag += 1
- 第四次循环 str为字符"l" spaceFlag = 1  然后 spaceFlag += 1
- ......
- 第七次循环 str为字符"w" spaceFlag = 4 然后 spaceFlag += 1
- 第八次循环str为空格 spaceFlag = 5 ，
在第八次循环的时候满足str为空格，且spaceFlag大于0，此时spaceFlag的值就是最后一个单词 world的长度5
