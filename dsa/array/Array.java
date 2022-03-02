package dev.santoshk;

public class Array {
    public static void main(String args[]) {
        int []marks = {97, 98, 95};

//        System.out.println(marks);  // this won't work; this prints memory address of marks

        for (int i=0; i<3; i++) {
            System.out.println(marks[i]);
        }
    }
}
