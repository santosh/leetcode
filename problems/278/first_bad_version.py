# The isBadVersion API is already defined for you.
# @param version, an integer
# @return an integer
# def isBadVersion(version):

class Solution:
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """

        left_index = 0
        right_index = n - 1
        
        while left_index <= right_index:
            mid_index = left_index + (right_index - left_index) // 2

            if isBadVersion(mid_index) == True:
                right_index = mid_index - 1
            else::
                left_index = mid_index + 1

        return left_index
