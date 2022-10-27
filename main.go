// This package is to create a string manipulation
//
// This is to create the manipulation on a string such as Cat() and Join() 
package stringmanipulation

import("strings")

// this is a function to concatinate a string
func Cat(xs []string) string { 
    s := " " 
    for _, v := range xs { 
        s += v
        s += " "
    }
    return s 
}

// this is a function to join a string  
func Join(xs []string) string {
    return strings.Join(xs, " ")
} 
