def find_max_sum(nums):
    if not nums:
        return 0
    
    max_current = max_so_far = nums[0]
    
    for num in nums[1:]:
        max_current = max(num, max_current + num)
        max_so_far = max(max_current, max_so_far)
    
    return max_so_far

# Example usage
nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
print(find_max_sum(nums))  # Output: 6
