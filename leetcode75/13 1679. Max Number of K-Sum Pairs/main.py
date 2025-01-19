"""class Solution:
def maxOperations(self, nums: list[int], k: int) -> int:
    start = 0  # count of the values searched
    nums.sort()  # sorting the array, so that we start with lowest, and keep on increasing
    total = 0

    print(nums)
    # remove any values greater than k
    left = 0
    right = len(nums) - 1
    while left <= right:
        mid = (left + right) // 2
        mid_value = nums[mid]
        if mid_value >= k:
            nums = nums[:mid]
            right = mid - 1
        else:
            left = mid + 1
    print(f"New nums becomes {nums}")
    left = 0
    while start < len(nums):
        # print(nums)
        # i is going to be the right pointer
        found_values = []
        for i in range(left + 1, len(nums)):
            # print(i)
            sum = nums[left] + nums[i]
            if sum == k:
                # print(
                #    f"found matching for position {left} and {i} with values {nums[left]},{nums[i]}"
                # )
                found_values.append(left)
                found_values.append(i)
                break
            elif sum > k:
                # since the array is sorted, any greater sum here means, adding it later will never result in k, so just break
                break
        if found_values:
            nums.pop(found_values[0])
            nums.pop(found_values[1] - 1)
            left = 0
            total += 1
        else:
            left += 1
        start += 1
    print(f"resulting array {nums}")
    return total
"""


class Solution:
    def maxOperations(self, nums: list[int], k: int) -> int:
        nums.sort()  # sorting the array, so that we start with lowest, and keep on increasing
        left = 0
        right = len(nums) - 1

        # Use binar search to ignore all the values that are already greater than k, because adding will never equals k
        while left < right:
            mid = (left + right) // 2
            if nums[mid] >= k:
                nums = nums[:mid]
                right = mid
            else:
                left = mid + 1
        left = 0
        right = len(nums) - 1
        total = 0
        while left < right:
            sum = nums[left] + nums[right]
            if sum == k:
                total += 1
                left += 1
                right -= 1
            elif sum < k:
                left += 1
            elif sum > k:
                right -= 1

        return total


s = Solution()
print(s.maxOperations([1, 2, 3, 4], 5))
print(s.maxOperations([3, 1, 3, 4, 3], 6))
print(
    s.maxOperations(
        [
            63,
            10,
            28,
            31,
            90,
            53,
            75,
            77,
            72,
            47,
            45,
            6,
            49,
            13,
            77,
            61,
            68,
            43,
            33,
            1,
            14,
            62,
            55,
            55,
            38,
            54,
            8,
            79,
            89,
            14,
            50,
            68,
            85,
            12,
            42,
            57,
            4,
            67,
            75,
            6,
            71,
            8,
            61,
            26,
            11,
            20,
            22,
            3,
            70,
            52,
            82,
            70,
            67,
            18,
            66,
            79,
            84,
            51,
            78,
            23,
            19,
            84,
            46,
            61,
            63,
            73,
            80,
            61,
            15,
            12,
            58,
            3,
            21,
            66,
            42,
            55,
            7,
            58,
            85,
            60,
            23,
            69,
            41,
            61,
            35,
            64,
            58,
            84,
            61,
            77,
            45,
            14,
            1,
            38,
            4,
            8,
            42,
            16,
            79,
            60,
            80,
            39,
            67,
            10,
            24,
            15,
            6,
            37,
            68,
            76,
            30,
            53,
            63,
            87,
            11,
            71,
            86,
            82,
            77,
            76,
            37,
            21,
            85,
            7,
            75,
            83,
            80,
            8,
            19,
            25,
            11,
            10,
            41,
            66,
            70,
            14,
            23,
            74,
            33,
            76,
            35,
            89,
            68,
            85,
            83,
            57,
            6,
            72,
            34,
            21,
            57,
            72,
            79,
            29,
            65,
            3,
            67,
            8,
            24,
            24,
            18,
            26,
            27,
            68,
            78,
            64,
            57,
            55,
            68,
            28,
            9,
            11,
            38,
            45,
            61,
            37,
            81,
            89,
            38,
            43,
            45,
            26,
            84,
            62,
            22,
            37,
            51,
            15,
            30,
            67,
            75,
            24,
            66,
            40,
            81,
            74,
            48,
            43,
            78,
            14,
            33,
            19,
            73,
            5,
            1,
            2,
            53,
            29,
            49,
            73,
            23,
            5,
        ],
        59,
    )
)
