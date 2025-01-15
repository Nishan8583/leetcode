class Solution:
    def productExceptSelf(self, nums: list[int]) -> list[int]:
        product_array = [1] * len(nums)
        prefix_product = 1
        postfix_product = 1

        # [1,2,3,4]
        # each value will in product_array will have the product of
        # all values in left
        for i in range(1, len(nums)):
            product = prefix_product * nums[i - 1]
            product_array[i] = product
            prefix_product = product

        for i in range(len(nums) - 1, -1, -1):
            print(i)
            product_array[i] = postfix_product * product_array[i]
            postfix_product = nums[i] * postfix_product
        return product_array


s = Solution()
print(s.productExceptSelf([1, 2, 3, 4]))
