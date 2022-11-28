class Solution {
    public int[] sortEvenOdd(int[] nums) {
        if (nums.length < 3) {
            return nums;
        }

        int[] odds = new int[nums.length/2];
        int[] evens = new int[nums.length/2+nums.length%2];

        int oddCounter = 0;
        for (int i = 1; i < nums.length; i=i+2) {
            odds[oddCounter] = nums[i];
            oddCounter++;
        }
        Arrays.sort(odds);

        int evenCounter = 0;
        for (int i = 0; i < nums.length; i=i+2) {
            evens[evenCounter] = nums[i];
            evenCounter++;
        }
        Arrays.sort(evens);

        oddCounter = odds.length-1;
        evenCounter = 0;

        for (int i = 0; i < nums.length; i++) {
            if (i%2==0) {
                nums[i] =  evens[evenCounter];
                evenCounter++;
            } else {
                nums[i] =  odds[oddCounter];
                oddCounter--;
            }
        }
        return nums;
    }
}
