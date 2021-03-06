package arrays

func maxSubArray(nums []int) int {
  n := len(nums)
  f := make([]int, n)
  f[0] = nums[0]

  for i := 1; i < n; i++ {
    f[i] = max(f[i-1]+nums[i], nums[i])
  }

  ans := f[0]
  for i := 1; i < n; i++ {
    ans = max(ans, f[i])
  }

  return ans
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}
