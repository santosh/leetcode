class GroupArrayByParity {
    public int[] sortArrayByParity(int[] nums) {
        int oddIndex =0;
        for (int i = 0; i < nums.length; i++) {
            if(nums[i]%2==0) {
                int t = nums[i];
                nums[i] = nums[oddIndex];
                nums[oddIndex] = t;
                oddIndex++;
            }
        }
        return nums;
    }

    public static void main(String[] args) {
        //        int[] arr = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
        int[] arr = {17, 25, 32, 44, 75, 66, 87, 82, 91, 15};
        //        int[] arr = {0, 2};


        arrange(arr);

        System.out.println(Arrays.toString(arr));
    }
}
