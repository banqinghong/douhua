## 存在重复元素
链接：https://leetcode-cn.com/problems/contains-duplicate/

## 思路
- 新建一个map用来存放已经做了判断的number
- 循环判断nums中的值是否在map中存在。如果不存在则继续下一次循环，如果存在则代表该num在列表中已经出现过一次。直接返回true
