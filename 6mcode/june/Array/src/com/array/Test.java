package com.array;//
//Write a program Sorting a HashMap according to value
//        ("Akshay",30);
//("Kinja",22);
//        ("Vrajesh",11);
//        ("Himanshu",45);


//sort map
//        collect this map
//rerurn it sorted map
import java.util.*;
import java.util.HashMap;
import java.util.stream.Collectors;

public class Test {
    static void main() {


        Map<String, Integer> map = new HashMap<>();
        map.put("Akshay", 30);
        map.put("Kinja", 22);
        map.put("Vrajesh", 11);
        map.put("Himanshu", 30);
        Map<String, Integer> smp =
                map.entrySet()
                        .stream()
                        .sorted(Map.Entry.comparingByValue())
                        .collect(Collectors.toMap(
                                Map.Entry::getKey,
                                Map.Entry::getValue,
                                (oldValue, newValue) -> oldValue,
                                LinkedHashMap::new
                        ));
        System.out.println("SortedMap" + smp);
    }
}
