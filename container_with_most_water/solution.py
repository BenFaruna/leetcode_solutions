from typing import List

class Solution:
    def water_area(self, current_height: int, next_height: int, width: int) -> int:
        """checks the height of both lines and return area based on their width"""
        if current_height > next_height:
            return next_height * width
        else:
            return current_height * width

    def max_water(self, current_height: int, heights: list) -> int:
        """checks the maximum area that can be covered by current height and
        subsequent heights"""
        width = 1
        max_area = 0
        for height in heights:
            area = self.water_area(current_height, height, width)
            if area > max_area:
                max_area = area

            width += 1
        return max_area

    def maxArea(self, height: List[int]) -> int:
        """final solution that return the largest area water can occupy
        within the line heights"""
        max_area_covered = 0

        i = 0
        while i < (len(height) - 1):
            current_i = height[i]
            subsequent_heights = height[i + 1:]
            area_covered = self.max_water(current_i, subsequent_heights)

            if area_covered > max_area_covered:
                max_area_covered = area_covered

            i += 1

        return max_area_covered
