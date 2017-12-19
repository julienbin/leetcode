class Solution(object):
    def findKth(self, nums1, nums2, k):
        len1 = len(nums1)
        len2 = len(nums2)
        if len1 > len2:
            return self.findKth(nums2, nums1, k)
        if len1 == 0:
            return nums2[k - 1]
        if k == 1:
            return min(nums1[0], nums2[0])
        p1 = min(k/2, len1)
        p2 = k - p1
        if nums1[p1 - 1] <= nums2[p2 - 1]:
            return self.findKth(nums1[p1:], nums2, p2)
        else:
            return self.findKth(nums1, nums2[p2:], p1)
        
        
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        len1 = len(nums1)
        len2 = len(nums2)
        if (len1 + len2) % 2 == 1:
            return self.findKth(nums1, nums2, (len1 + len2)/2 + 1)
        else:
            return (self.findKth(nums1, nums2, (len1 + len2)/2) + self.findKth(nums1, nums2, (len1 + len2)/2 + 1)) * 0.5
        
