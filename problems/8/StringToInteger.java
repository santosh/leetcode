import java.util.ArrayList;
import java.util.Objects;

public class StringToInteger {
    public static int myAtoi(String s) {
        // definitions
        boolean numIsNegative = false;
        int counter = 0;

        // ignore leading whitespace
        s = s.stripLeading();

        if (Objects.equals(s, "")) {
            return 0;
        }

        char[] charArray = s.toCharArray();
        if (charArray[0] == '-') {
            numIsNegative = true;
            counter = 1;
        } else if (charArray[0] == '+') {
            counter = 1;
        }

        ArrayList<Integer> num = new ArrayList<Integer>();

        for (; counter < charArray.length; counter++) {
            // check if the character is in int digit range from ascii table
            if (charArray[counter] - '0' > 9 || charArray[counter] - '0' < 0) {
                break;
            }
            num.add(charArray[counter] - '0');
        }

        int finalNum = 0;
        int maxLimit = Integer.MAX_VALUE / 10;

        for (int i = 0; i<num.size(); i++) {
            int tempNum = num.get(i);
            if(finalNum > maxLimit  || (finalNum == maxLimit && tempNum > 7)){
                return numIsNegative ? Integer.MIN_VALUE : Integer.MAX_VALUE;
            }
            if (tempNum != 0) {
                while (tempNum > 0) {
                    finalNum *= 10;
                    tempNum /= 10;
                }
            finalNum += num.get(i);
            } else {
                finalNum *= 10;
            }
        }

        return numIsNegative ? -finalNum : finalNum;
    }

    public static void main(String[] args) {
        System.out.println(myAtoi("42"));
        System.out.println(myAtoi("+42"));
        System.out.println(myAtoi("   -42"));
        System.out.println(myAtoi("004193 with words"));
        System.out.println(myAtoi("4193 with words"));
        System.out.println(myAtoi("21474836460"));
        System.out.println(myAtoi("2147483646")); // 2147483646
    }
}
