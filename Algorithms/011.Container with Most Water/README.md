# [Container with Most Water](https://leetcode-cn.com/problems/container-with-most-water/)

## 题目

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

**Note**: You may not slant the container and n is at least 2.

The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

**Example**:
```
Input: [1,8,6,2,5,4,8,3,7]
Output: 49
```

## 思路
1. 穷举
2. 设立两个指针，一个从头一个从尾，相向而行遍历数组，每次舍弃较短边。（当前Area的短板在较短边， 要想比当前Area大，只能舍弃较短边）