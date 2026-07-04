package com.array;

public class LargestNum {
    static void main() {

        int arr[]={3,2,4,5,8,1};

        int max=largestNum(arr);
        System.out.println(max);
        int min=smallestNum(arr);
        System.out.println(min);
    }
    @org.jetbrains.annotations.Contract(pure = true)
    private static int largestNum(int []num){
        int largest=num[0];
        for(int i=1;i<num.length;i++){
            if(num[i]>largest){
                largest=num[i];
            }
        }
        return  largest;
    }
    private static int smallestNum(int []num){
        int smallestNum=num[0];
        for (int n:num){
            if(n<smallestNum){
                smallestNum=n;
            }
        }
        return smallestNum;
    }
    private static int secondLargestNum(int []num){
        int largest=num[0];
        int secondLargestNum=0;
        for(int i=1;i<num.length;i++){
            if(num[i]>largest){
                secondLargestNum=largest;
                largest=num[i];

            }
        }
        return  secondLargestNum;
    }
}
