func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for i, value := range nums {
        complement := target - value

        if compIndex, ok := seen[complement]; ok {
            return []int{compIndex, i}
        }

        seen[value] = i
    }

    return nil
}