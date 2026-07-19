package com.pattern.twoPointer;

import java.util.Arrays;
import java.util.HashMap;

public class TwoSum {

//    #Aprroch 1
    public static int twoSum(int []nums,int target){
        for (int j : nums) {
            for (int num : nums) {
                if (j + num == target) {
                    return 1;
                }
            }
        }
        return 0;

    }
    // Approch 2
    public static int twoSum1(int []nums,int target){
        Arrays.sort(nums);
        int left=0;
        int righ=nums.length-1;
        for (int i = 0; i < nums.length; i++) {
            if (nums[left]+nums[righ]==target){
                System.out.println("target found on this direction :"+left+" "+righ);
                return  1;
            }
        }

        return 0;
    }

    public  int[] twoSum2(int []nums,int target){
        HashMap<Integer,Integer> data=new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
           int need=target-nums[i];
           if (data.containsKey(need)){
               return new int[]{data.get(i),i};
           }else {
               data.put(nums[i],i);
           }

        }

        return  new int[]{0,0};
    }
    static void main() {
//        int res=twoSum(new int[]{1,2,7,6,3,2},9);
        int res = twoSum(new int[]{1, 2, 7, 6, 3, 2}, 9);
        System.out.println(res);
    }
}
