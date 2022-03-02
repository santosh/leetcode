package dev.santoshk;

public class Array {
    public static void main(String args[]) {
        int []marks = new int[3];
        marks[0] = 97;  // phy
        marks[1] = 98;  // chem
        marks[2] = 95;  // eng

//        System.out.println(marks);  // this won't work; this prints memory address of marks

        System.out.println(marks[0]);
        System.out.println(marks[1]);
        System.out.println(marks[2]);
    }
}
