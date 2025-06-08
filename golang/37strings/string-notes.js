/*

strconv.Atoi() 
-> converts a string to an integer
-> returns an error if the string does not represent a valid integer
"1234" -> 1234

strings.ReplaceAll()
-> used to trim all white spaces from a string in Golang
    
    sample := " This is a sample string   "
    noSpaceString := strings.ReplaceAll(sample, " ", "")
    fmt.Println(noSpaceString)
    -> replace all blank space with no space
    -> Thisisasamplestring
    noSpaceString := strings.ReplaceAll(sample, " ", "hello")
    -> replace all blank space with hello text
    -> helloThishelloishelloahellosamplehellostringhellohellohello


Multiline String:
multiline := `This is 
              a multiline 
              string`
A backquote (`) can be used to write a multiline string in Golang. Note that a string encoded in backquotes is a raw literal string and doesn’t honor any kind of escaping. Thus \n, \t are treated as a string literal when using back quotes


String Compare:
In Golang string are UTF-8 encoded. strings package of GO provides a Compare method that can be used to compare two strings in Go. Note that this method compares strings lexicographically.
-> res := strings.Compare("abc", "abc")
 0 if a==b
-1 if a < b
+1 if a > b
[In Golang string are UTF-8 encoded - meaning that every string in Go is internally represented as a sequence of bytes using UTF-8 encoding.]
-> Go Strings are Immutable Byte Sequences
-> A Go string is a read-only slice of bytes ([]byte).
-> Go doesn't store characters (rune) directly but instead stores their UTF-8 byte representation.


    s := "Hello, 世界" // 世界 (Chinese for "world") has multi-byte characters

    fmt.Println("Length in bytes:", len(s)) // Output: 13 (not 9, because 世界 takes more than 1 byte each)

    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i]) // Prints raw bytes in hex
    }
    fmt.Println()

    // Correct way to iterate over runes
    for _, r := range s {
        fmt.Printf("%c ", r) // Prints actual characters
    }

    Length in bytes: 13
    48 65 6c 6c 6f 2c 20 e4 b8 96 e7 95 8c 
    H e l l o ,   世 界 

    Indexing (s[i]) accesses bytes, not characters.
    Use rune slices ([]rune(s)) to work with Unicode characters


String contains another string:
    present := strings.Contains("abc", "ab")

    //Case 1 s contains sep. Will output slice of length 3
    res := strings.Split("ab$cd$ef", "$")
    fmt.Println(res)

    //Case 2 s doesn't contain sep. Will output slice of length 1
    res = strings.Split("ab$cd$ef", "-")
    fmt.Println(res)

    //Case 3 sep is empty. Will output slice of length 8
    res = strings.Split("ab$cd$ef", "")
    fmt.Println(res)

    //Case 4 both s and sep are empty. Will output empty slice
    res = strings.Split("", "")
    fmt.Println(res)


Get all words from a sentence in GO:
-> it returns a slice of substrings of the input string
    //Case 1 Input string contains single spaces
    res := strings.Fields("ab cd ef")
    fmt.Println(res)
    [ab cd ef]

    //Case 2 Input string doesn't contain white spaces
    res = strings.Fields("abcdef")
    fmt.Println(res)
    [abcdef]

    //Case 3 Input string contains double white spaces and spaces at end too.
    res = strings.Fields("ab  cd   ef ")
    fmt.Println(res)
    [ab cd ef]

    //Case 4: Input string contain white spaces only. Will output a empty slice
    res = strings.Fields("  ")
    fmt.Println(res)
    []

    //Case 5: Input string is empty. Will output a empty slice
    res = strings.Fields("")
    fmt.Println(res)
    []


Join a string by delimiter or a separator in Go:
-> this function takes a slice of string and a delimiter and it returns a combined string joined by a delimiter
    //Case 1 s contains sep. Will output slice of length 3
    res := strings.Join([]string{"ab", "cd", "ef"}, "-")
    fmt.Println(res)

    //Case 2 slice is empty. It will output a empty string
    res = strings.Join([]string{}, "-")
    fmt.Println(res)

    //Case 3 sep is empty. It will output a string combined from the slice of strings
    res = strings.Join([]string{"ab", "cd", "ef"}, "")
    fmt.Println(res)

    











































*/