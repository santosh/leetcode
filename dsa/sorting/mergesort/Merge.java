import java.util.*;

public class Merge {
    public static void sort(int[] arr, int start, int end) {
        if (start >= end) return;

        int mid = start + (end - start)/2;

        sort(arr, start, mid);
        sort(arr, mid+1, end);

        merge(arr, start, mid, end);
    }

    public static void merge(int[] arr, int start, int mid, int end) {
        int[] temp = new int[arr.length];

        int tempIndex = start;
        int leftStart = start;
        int rightStart = mid+1;

        while (leftStart <= mid && rightStart <= end) {
            if (arr[leftStart] < arr[rightStart]) {
                temp[tempIndex] = arr[leftStart];
                leftStart++;
            } else {
                temp[tempIndex] = arr[rightStart];
                rightStart++;
            }
            tempIndex++;
        }

        while (leftStart <= mid) {
            temp[tempIndex++] = arr[leftStart++];
        }

        while (rightStart <= end) {
            temp[tempIndex++] = arr[rightStart++];
        }

        for (int i = start; i <= end; i++) {
            arr[i] = temp[i];
        }
    }

    public static void main(String[] args) {
        int[] arr = {6, 1, 5, 3, 7, 8, 3, 4, 2, 9, 0};
        sort(arr, 0, arr.length-1);
        System.out.println(Arrays.toString(arr));
    }
}
