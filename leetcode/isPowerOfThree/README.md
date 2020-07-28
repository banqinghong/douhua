## 3的幂
链接：https://leetcode-cn.com/problems/power-of-three/  
简单描述：判断给定的整数是否为3的幂次方

## 思路
- 判断n是否为1或者3，如果成立直接返回true  
- 判断n对3取模是否为0，如果不为0或者n为0，返回false
- 如果都不成立则取n除以3的余数进行递归
