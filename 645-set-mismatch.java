class Solution {
    public int[] findErrorNums(int[] nums) {
        HashMap<Integer, Integer> memo = new HashMap<Integer, Integer>();
        for (int num: nums) {
            memo.put(num, memo.getOrDefault(num, 0) + 1);
        }
        int dup = -1, missing = -1;
        for (int i = 1; i <= nums.length; i++) {
            if (memo.containsKey(i)) {
                if (memo.get(i) == 2) {
                    dup = i;
                }
            } else {
                missing = i;
            }
        }
        return new int[]{dup, missing};
    }
}
