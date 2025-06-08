/*
Creating a folder:
    err := os.Mkdir("/Users/temp", 0755)

\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

os.Chdir() is used to change the current working directory to the named directory in golang. It is similar to the cd command.
    os.Chdir("/Users")
    newDir, err := os.Getwd()
    if err != nil {
    }
    fmt.Printf("Current Working Direcotry: %s\n", newDir)

\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

Read a file into a variable in Go (Golang)
There are various ways of reading a file into a variable in golang. Here are some of the ways
-   Using ReadFile function provided by the ioutil package
-   Using os.Open and then using bytes.Buffer
-   Using os.Open and then using strings.Builder

1.
func main() {
    fileBytes, err := ioutil.ReadFile("test.png")
    if err != nil {
        panic(err)
    }
    fileString := string(fileBytes)
    fmt.Println(fileString)
}

2.
func main() {
	file, err := os.Open("test.png")
	if err != nil {
		log.Fatalf("Error while opening file. Err: %s", err)
	}
	defer file.Close()

	fileBuffer := new(bytes.Buffer)
	fileBuffer.ReadFrom(file)
	fileString := fileBuffer.String()

	fmt.Print(fileString)
}

\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

Buffered write using bufio package:

-> Example: With Default Buffer Size of 4096 bytes
func main() {
    file, err := os.Create("./temp.txt")
    if err != nil {
        log.Fatal(err)
    }
    writer := bufio.NewWriter(file)
    linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line."}
    for _, line := range linesToWrite {
        bytesWritten, err := writer.WriteString(line + "\n")
        if err != nil {
            log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
        }
        fmt.Printf("Bytes Written: %d\n", bytesWritten)
        fmt.Printf("Available: %d\n", writer.Available())
        fmt.Printf("Buffered : %d\n", writer.Buffered())
    }
    writer.Flush()
}
        -> It will write to temp.txt in the current directory. Write happens at the end when writer.Flush() is called as buffer never really gets full in between.

-> Example: With Custom Buffer Size of 10 bytes
We can also create a custom buffer size for our writer using bufio.NewWriterSize(). In the below example, we create a write with a custom buffer size of 10 bytes. If you carefully notice the bytes “Available” and “Buffered” in output, you will find that writes to file keep happening in between when the buffer is full and at the end when writer.Flush() is called.

//Write file using bufio writer
func main() {
    file, err := os.Create("./temp  .txt")
    if err != nil {
        log.Fatal(err)
    }
    writer := bufio.NewWriterSize(file, 10)
    linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line."}
    for _, line := range linesToWrite {
        bytesWritten, err := writer.WriteString(line + "\n")
        if err != nil {
            log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
        }
        fmt.Printf("Bytes Written: %d\n", bytesWritten)
        fmt.Printf("Available: %d\n", writer.Available())
        fmt.Printf("Buffered : %d\n", writer.Buffered())
    }
    writer.Flush()
}

-> Using file.Write()
It does not maintain any buffer and writes to the file immediately as soon as the write is called. See the example below

func main() {
    file, err := os.Create("./test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line."}
    for _, line := range linesToWrite {
        file.WriteString(line + "\n")
    }
}


-> Using ioutil.WriteFile()
ioutil.WriteFile() is kind of a shortcut to writing to a file. It takes in three-parameter – (1) file (2) The string to be written (3) Permission mode of the file. The third parameter is used to create a file with that permission if the file doesn’t already exist. 
One calling ioutil.WriteFile(), it will
-   Create the file if not present with the specified permission
-   Write to the file
-   Close the file

func main() {
    linesToWrite := "This is an example to show how to write to file using ioutil"
    err := ioutil.WriteFile("temp.txt", []byte(linesToWrite), 0777)
    if err != nil {
        log.Fatal(err)
    }
}

\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

Append to a file:

func main() {
    //Write first line
    err := ioutil.WriteFile("temp.txt", []byte("first line\n"), 0644)
    if err != nil {
        log.Fatal(err)
    }
    
    //Append second line
    file, err := os.OpenFile("temp.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Println(err)
    }
    defer file.Close()
    if _, err := file.WriteString("second line"); err != nil {
        log.Fatal(err)
    }
    
    //Print the contents of the file
    data, err := ioutil.ReadFile("temp.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(data))
}


\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

os.Remove() function can be used to delete a file in Golang. Below is the signature of this function
err := os.Remove("sample.txt")
if err != nil {
    log.Fatal(err)
}

os.Remove() function can be used to delete a folder in Golang. Below is the signature of this function
err := os.Remove("sample")


\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


os.Chtimes() function can be used to change the mtime(modified time) or atime(access time) of a file in Golang. Below is the signature of the function.

func main() {
    fileName := "sample.txt"
    currentTime := time.Now().Local()

    //Set both access time and modified time of the file to the current time
    err := os.Chtimes(fileName, currentTime, currentTime)
    if err != nil {
        fmt.Println(err)
    }
}

\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


os.Rename() function can be used to rename a file or folder. Below is the signature of the function.
func Rename(old, new string) error
-> old and new can be fully qualified as well. If the old and new path doesn’t lie in the same directory then os.Rename() function behaves the same as moving a file or folder.
-> A fully qualified path in Golang refers to the absolute path of a file, directory, or package that uniquely identifies its location in the system or module.

-> Rename a file
func main() {
    //Create a file
    file, err := os.Create("temp.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    //Change permission so that it can be moved
    err = os.Chmod("temp.txt", 0777)
    if err != nil {
        log.Println(err)
    }


    err = os.Rename("temp.txt", "newTemp.txt")
    if err != nil {
        log.Fatal(err)
    }
}

-> Rename a folder
func main() {
    //Create a directory
    err := os.Mkdir("temp", 0755)
    if err != nil {
        log.Fatal(err)
    }
    err = os.Rename("temp", "newTemp")
    if err != nil {
        log.Fatal(err)
    }
}


\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


Check if file is a directory or not in Go
See the below code to know if a file is a file or it is a directory
-   If temp is a file output will be = “temp is a file”
-   If temp is a directory output will be = “temp is a directory”

In Golang, os.Stat() is a function from the os package that retrieves information about a file or directory, such as its size, permissions, modification time, and whether it exists.

var (
    fileInfo *os.FileInfo
    err      error
)

func main() {
    info, err := os.Stat("temp")
    if os.IsNotExist(err) {
        log.Fatal("File does not exist.")
    }
    if info.IsDir() {
        fmt.Println("temp is a directory")
    } else {
        fmt.Println("temp is a file")
    }
}


\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


os.Create() can be used to create an empty file in go. The signature of the function is
func Create(name string) (*File, error) 
file, err := os.Create("emptyFile.txt")

Create a named file with mode 0666
It truncates the file if it already exits
In case of path issue, it returns a Path error
It returns a file descriptor which can be used for both reading and write
-> A file descriptor (FD) is a low-level, unique identifier (usually an integer) assigned by the operating system to an open file, socket, or other I/O resources. It helps manage file operations like reading, writing, and closing.


\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


os.Stat and os.IsNotExist() can be used to check whether a particular file or directory exist or not.

1.
fileinfo, err := os.Stat("temp.txt")
if os.IsNotExist(err) {
    log.Fatal("File does not exist.")
}
log.Println(fileinfo)

2.
folderInfo, err := os.Stat("temp")


\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


Iterate over all files and folders for a path in Go (Golang)
-> ‘Walk’ function of ‘filepath’ package can be used to recursively iterate over all files/folder in a directory tree.
-> ‘Walk’ function will walk the entire tree rooted at the root path include all subdirectories.  Below is the signature of the function

The WalkFunc will be called for with the path of the file/folder and fileInfo or the error if any error occurred while walking that file/folder.

Some things to note about Walk function
-   All errors are filtered. An error might arise while opening/visiting the file
-   The function does not follow symbolic links
-   The files are walked in lexical order

func main() {
    currentDirectory, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    iterate(currentDirectory)
}

func iterate(path string) {
    filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Fatalf(err.Error())
        }
        fmt.Printf("File Name: %s\n", info.Name())
        return nil
    })
}


\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


Touch a file in Go (Golang)
-> Touching a file means
    Create an empty file if the file doesn’t already exist
    If the file already exists then update the modified time of the file.

func main() {
    fileName := "temp.txt"
    _, err := os.Stat(fileName)
    if os.IsNotExist(err) {
        file, err := os.Create("temp.txt")
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
    } else {
        currentTime := time.Now().Local()
        err = os.Chtimes(fileName, currentTime, currentTime)
        if err != nil {
            fmt.Println(err)
        }
    }
}

-> When running the first time it will create the file. After the first time, it will update the modified time of the file.


\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


Move file from one location to another in or command mv in Go (Golang)
func Rename(oldpath, newpath string) error

func main() {
    //Create a file
    file, err := os.Create("temp.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    //Change permission so that it can be moved
    err = os.Chmod("temp.txt", 0777)
    if err != nil {
        log.Println(err)
    }

    newLocation := "~/Desktop/temp.txt"
    err = os.Rename("temp.txt", newLocation)
    if err != nil {
        log.Fatal(err)
    }
}
   

\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


Get File Name, Size, Permission Bits, Mode, Modified Time in Go (Golang)
os.Stat() function can be used to the info of a file in go. This function returns stats which can be used to get
- Name of the file
- Size of the file in bytes
- Modified Time of the file
- Permission Bits or Mode of the file

func main() {
    //Create a file
    file, err := os.Create("temp.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    //Write something to the file
    file.WriteString("some sample text" + "\n")

    //Gets stats of the file
    stats, err := os.Stat("temp.txt")
    if err != nil {
        log.Fatal(err)
    }

    //Prints stats of the file
    fmt.Printf("Permission: %s\n", stats.Mode())
    fmt.Printf("Name: %s\n", stats.Name())
    fmt.Printf("Size: %d\n", stats.Size())
    fmt.Printf("Modification Time: %s\n", stats.ModTime())
}

[
    Permission: -rwxrwxrwx
    Name: temp.txt
    Size: 17
    Modification Time: 2020-04-16 22:26:47.080128602 +0530 IST
]

   
\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


Make a copy of a file in Go (Golang)

1.
io.Copy() can be used to create a copy of the file from src to dest. A successful copy will return err != nil instead of err == EOF
Below is the signature of the function. It returns the number of bytes written to the dest
-> func Copy(dst Writer, src Reader) (written int64, err error)

// Copy a file
func main() {
    // Open original file
    original, err := os.Open("original.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer original.Close()

    // Create new file
    new, err := os.Create("new.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer new.Close()

    //This will copy
    bytesWritten, err := io.Copy(new, original)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Bytes Written: %d\n", bytesWritten)
}

2.
If the contents of the file is less then we can also read the contents of the file first and then write all the contents to the new file. See below code.
// Copy a file
func main() {
    //Read all the contents of the  original file
    bytesRead, err := ioutil.ReadFile("original.txt")
    if err != nil {
        log.Fatal(err)
    }

    //Copy all the contents to the desitination file
    err = ioutil.WriteFile("new.txt", bytesRead, 0755)
    if err != nil {
        log.Fatal(err)
    }
}

















































*/