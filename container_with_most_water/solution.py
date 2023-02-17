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

    def max_area_list(self, heights: List[int]) -> List[int]:
        """returns a list containing the maximum area convered from each height"""
        max_area_covered = []

        i = 0
        while i < (len(heights) - 1):
            current_i = heights[i]
            subsequent_heights = heights[i + 1:]
            max_area_covered.append(self.max_water(current_i, subsequent_heights))
            i += 1
        return max_area_covered

    def maxArea(self, height: List[int]) -> int:
        """final solution that return the largest area water can occupy
        within the line heights"""
        max_area_covered = self.max_area_list(height)
        max_area = 0
        for area in max_area_covered:
            if area > max_area:
                max_area = area
        return max_area
