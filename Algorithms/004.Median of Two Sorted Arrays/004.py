class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        m = len(nums1)
        n = len(nums2)
        if m == 0:
            return nums2[n//2]*1.0 if n % 2 == 1 else (nums2[n//2-1]+nums2[n//2])/2
        if n == 0:
            return nums1[m//2]*1.0 if m % 2 == 1 else (nums1[m//2-1]+nums1[m//2])/2
        m_start = 0
        n_start = 0
        target = (m+n)//2
        while target > 1:
            temp = target//2
            if nums1[m_start+temp-1] < nums2[n_start+temp-1]:
                m_start = m_start + temp
            elif nums1[m_start+temp-1] > nums2[n_start+temp-1]:
                n_start = n_start + temp
            else:
                target = target - temp * 2
                m_start = m_start + temp
                n_start = n_start + temp
                break
            target = target - temp
        if (m+n)%2 == 1:
            if target == 1:
                return nums1[m_start]*1.0 if nums1[m_start] > nums2[n_start] else nums2[n_start]*1.0
            else:
                return nums1[m_start]*1.0 if nums1[m_start] < nums2[n_start] else nums2[n_start]*1.0
        else:
            if target == 1:
                if nums1[m_start] < nums2[n_start]:
                    m = nums1[m_start]
                    n = nums2[n_start] if m_start==m-1 or nums2[n_start]<nums1[m_start+1] else nums1[m_start+1]
                else:
                    m = nums2[n_start]
                    n = nums1[m_start] if n_start==n-1 or nums1[m_start]<nums2[n_start+1] else nums2[n_start+1] 
            else:
                m = nums1[m_start-1] if nums1[m_start-1] > nums2[n_start-1] else nums2[n_start-1]
                n = nums1[m_start] if nums1[m_start] < nums2[n_start] else nums2[n_start]
            return (m+n)/2