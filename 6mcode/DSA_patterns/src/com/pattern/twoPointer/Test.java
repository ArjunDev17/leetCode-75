package com.pattern.twoPointer;
//Given an integer array A and an integer B,
// find the total number of contiguous subarrays whose sum is equal to B.

//        Input: A = [1, 2, 3]
//        B = 3 Output: 2
//        Subarrays are: [1, 2] [3]
//A = [3, 4, -7, 1, 3, 3, 1, -4]
//B = 7 output should be 4
import java.util.HashMap;
import java.util.Map;

public class Test {
    static void main() {
        int []A={3, 4, -7, 1, 3, 3, 1, -4};
        int target=7;
        Map<Integer,Integer> prefixMap=new HashMap<>();
        prefixMap.put(0,1);
        int count=0;
        int prefixSum=0;
        for(int num:A){
            prefixSum +=num;
            if(prefixMap.containsKey(prefixSum-target)){
                count+= prefixMap.get(prefixSum-target);
            }
            prefixMap.put(prefixSum,prefixMap.getOrDefault(prefixSum,0)+1);
        }
        System.out.println(count);
    }
}
