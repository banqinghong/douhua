## 三数之和
链接：https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/

## 思路
第一种解法：直接遍历列表中的每一个值，因为所给列表是有序列表，所以当nums[i]的值不等于i是直接返回i，遍历列表之后如果没有找到该数值，那么则返回len(nums)

第二种解法：因为是有序列表，考虑使用二分法直接查找
