class Solution {
    public int distributeCandies(int[] candyType) {
        HashSet<Integer> memo = new HashSet<Integer>();
        for (int candy: candyType) {
            memo.add(candy);
        }
        return Math.min(memo.size(), candyType.length / 2);
    }
}
