/*

Compiled Languages -> Entired code is read first and then a exe file is created to execute the code
Ex. C++, Golang
Interpreted Languages -> Code is read line by line and executed line by line
Ex. JavaScript

Go tool can run file directly without virtual machine.
There is a lot of control that the GO languageg takes from the operating ststem and gives to us directly, specially in the cases of thread.
Executables are executed at compile time.

You can build anything over the cloud infrastructure.

Golang is not a proper object oriented programming language.
According to the proper defination of OOPS, it should have classes and objects, but go does not have them, it has structs and methods.
Go also does not have operator and method overloading 

try/catch is missing in this language, but it is designed in such a way that we would not need it.

lexer does a lot of work
we may or may not be required to put a semicolon at the end of each line.

// Working with Go

Go to the file in which our go file is present, right click and open the directory in terminal
Run the command "go mod init <module_name>" to create a module

In GO, main function is the entry point of the program

Go automcatically cleans and uncleans apackages required in the code

To run the file:
-> go run <packagename.go>
    -> go run main.goI. I. 
-> This will only run the file and not build it thus will not create an executable file
-> It compiles and runs the program. 

go version
-> to test golang version

go help
-> This will show all the commands that can be used with go
-> These are essentially called as tools

go run
-> compiles and runs the program
-> go run <package>

gopath
-> to determine the path of the go file where it is installed
-> The Go path is used to resolve import statements.
-> The GOPATH environment variable lists places to look for Go code.

We need to set up environment variables path for Golang and we can find it in the variable known as GOPATH. 
We can find the exact path of go by runnig the command: 
-> 'go env GOPATH' 
-> Output: C:\Users\ARYAN GUPTA\go
    -> in go/pkg/mod we can find all the packages that we have installed
        ->  d-----        03-12-2024  05:28 PM                github.com
            d-----        03-12-2024  05:27 PM                golang.org
            d-----        03-12-2024  05:28 PM                honnef.co
            d-----        03-12-2024  05:28 PM                mvdan.cc
    -> in go/bin we can find all the executables that we have created



The Lexer just convert a sequence of characters into a sequence of tokens which follows a specific regular expression, 
Lexer doesn't check if the source code follows the grammar of the language, it is done by the Parser.
In go compiler the lexer automatically inserts a semicolon token into the token stream which is then passed to the parser 
which checks if the token stream follows the grammar of the language.



 So what does lexer does that if it's only delta? This lecture comes in. 



Types in Go:
-> Case Sensitive
-> Variable type should be known in advance
-> Everything is a type
-> fmt.Println -> capital 'P' suggests that it is a publically exported function from whichever class it was built in
-> Basic types:
    String
    Bool
    Integer - int, uint8, uint64, int8, int64, uintptr
    Floating - float32, float64 
    Complex
-> Composite types:
    Arrays
    Slices
    Maps
    StructsI. 
    Pointers

    Functions are also internally treated as a type in Go

If you want to check the upper limit use this formula 2^n where n is number of bits. Example :  uint8 n = 8 so upper limit would be 2^8 = 256, so the range is 0 - 255

uint8 and int8 basically represent unsigned and signed integers having different ranges
For Ex. uint8 -> 0 to 255
        int8 -> -128 to 127
        var x uint8 = 255 // will work fine
        var x uint8 = 256 // will give an error

uint8 - set of unsigned 8 bit integers (0 - 255)
uint16 - set of signed 8 bit integers (0 - 65535)
int8 - set of signed 8 bit integers (-128 - 127)
int8 - set of signed 8 bit integers

byte - alias for uint8
rune - alias for int32

Go is a statically typed language, which means that the type of a variable is known at compile time.

var username string = "aryan"
    -> this is how variable are made in GoI want to have. 

I. OK. var isLoggedIn bool = true/false

fmt.Println("Type of Username: %T \n", username)
    -> Above line will just print '%T' as text and not the type of the variable as these formatting directives are not allowed in the Println function
fmt.Printf("Type of Username: %T \n", username)
    -> this will print the type of the variable

In Go, := is for declaration + assignment, whereas = is for assignment only.
For Ex. var foo int = 10 is the same as foo := 10.

:= -> walrus operator
-> can be used only inside a function or a scope. cant be used for global variables
> responsible for the comma, ok syntax

fmp.Println() -> automatically adds a newline character at the end of the line

# USER INPUT 
We can take input from user in Go using the comma ok or comma error syntax
walrus operator is designed to omit data type of variables because we do not know the data type of the input that we are going to get from the user

we can use the bufio and os package in Go to take input from the user
it has buffer which we can use to read variables and store valeus in variables

Go provides us with a inbuilt time package which we can use to get the current time and date
time package provides with a lot more features like sleep etc that we will use later on


# If for some reason, we want to build a program which whenever opened prints the current time in the console
we can run 'go env' to get the envrionment variables of Go
GOOS=windows
-> we can use this to check the OS that we are using and build the code for any OS irrespective of the OS that we are using

go build -> will automatically find the go file and build it for the OS you are on
But if you want to build the file for a different OS, you can use the command:
GOOS=linux go build
    -> This will build the file for linux OS
GOOS="windows" go build mytime.go
    -> This will build the file for windows OS names mytime.exe


# MEMORY MANAGEMENT
-> Go handles this automatcially and it is not the developers job to handle memory management
Memory allocation and deallocation happens autotmatically in Go
new() and make() are methods that are used to allocate memory in Go and in that memory we can store daa structures

-> new()
    - memory allocated but not initialised
    - you will get a memory address
    - zeroed storage -> we cannot put any data in it intially

-> make()
    - memory allocated and initialised so that we can put any values in it
    - you will get a memory address
    - non-zeroed storage

Garbage Collection happens automatically in Go [GOGC]
-> It is a process of releasing the memory that is no longer in use
Anything that is out of scope or nil is elegible for garbage collection 

Runtine is a go package to tell you the number of available packages etc.

# Pointers
There are situations when we pass on a variable to a function, 
it creates a copy of that variable and we end up not manipulating the original variable. 
To solve this problem, pointers exist which pass on the reference of that particular variable. 
    We can pass on the Memory address of the variable. This way we can make sure that whatever 
    variable we are passing on is the actual variable that we want to pass and not a copy of it. 


Maps, also known as hash tables in other languages, are unordered collections of key-value pairs.
Retrieval is fast in maps as we know the exact key we need to locate
Maps can be converted into io.Writer format useing bufer.NewBuffer()


Defer statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.
Ideally code is executed line by line, but the files we mark as defer are exceuted at the end of the function/file
Deferred statements are executed in Last In First Out order(in the reverse order in which they were invoked)

net/http is the fastest and the most recoomended package to create a web server in Go and handle HTTP requests and responses
-> using this package we can handle cookies, manipulate headers, and much more 
Whenever we make a request to any website, we will get a response object back, which can be studied in detail from the documentation
Response object has a lot of fields and methods that we can use to get the data that we want
-> status(200 ok), status code(200), proto, etc.

We always have to close the response connection





--------------------------------------------------------------------------------------------------------------------------------

Go is a statically typed language, which means either we have to declare variable types explicitly or they have to inferred
These types cannot be changed afterwards, atleast not till type conversion

var myString string -> explicitly mentioned
var myString = "hello" -> inferred

Go is a strongly typed language, which means we can only perform the operations of a particular type in particualr types of variables.
unlike in JS which is a weakly typed language, where we can perform string operations on numbers and vice-versa

Since Go is a strongly and statically typed language, compiler can do more strict checking and force users to fix the bugs before proceeding furthur

Go is compiled, which converts files to binary files which can be ran as a standalone program (program.exe), this comes on contrast with languages such as python which are interpreted and run code line by line causing an extra overhead

Go has a fast compile time which makes the testing process much faster 

Go has inbuild concurrency, we do not need special packages or workarounds to achieve parallism; using Goroutines

Packages in Go are basically just folders which are collection of files. collection of such folders are called modules
Hence, when we initialize a project, we are initializing a module

To initialize a project/ module -> go mod init projectname

Every go file is a part of a package and we identify the name of the package it belongs to by specifying it at the top of the line
'package main' -> the package name has to be the same for all files within the folder 

fmt.PrintLn() takes the string we are passing, attaches a new line character and then print it on the console

go build main.go -> this produces a binary file called main
go run main


--------------------------------------------------------------------------------------------------------------------------------

production and consumption mismatch -> queues used -> production more but consumption less so store in queue


Producer-Consumer Model
Role-Based: In this model, there are two roles—producers and consumers.
Task Queue: Producers generate tasks or data and place them in a queue. Consumers take tasks or data from the queue and process them.
Direct Communication: The communication is usually direct, meaning consumers pull data from the queue when they are ready to process it.
Synchronization: The queue acts as a buffer to synchronize the production and consumption rates. If the queue is empty, consumers wait; if it's full, producers wait.

Publish-Subscribe Model (Pub-Sub)
Event-Based: This model is event-driven and involves publishers and subscribers.
Topics/Channels: Publishers send messages to specific topics or channels. Subscribers receive messages from those topics or channels they are interested in.
Decoupled Communication: The communication is decoupled, meaning publishers don't need to know about subscribers and vice versa. A message broker usually handles the distribution of messages.
Asynchronous: Messages are typically pushed to subscribers asynchronously. Subscribers don't need to request or pull messages; they receive them as they are published.



Concurrency v/s Parallelism
Goroutines are a way to achieve parallelism in Go

Thread - managed by the OS, has its own stack, and is expensive to create
    - fixed stack - 1MB
Goroutine - managed by the Go runtime, has a dynamic stack, and is cheap to create
    - flexible stack - 2KB (lightweight - this increases the number of threads that can be created in a limited amount of memory)
    - which means it can fire up more threads without getting the permission from the OS(however there is a fixed no. of threads that can be created)
    - hence, goroutines are more popular in cloud infrastructure as we can create more threads

Do not communicate by sharing memory; instead, share memory by communicating
http versions 1/2/3
htpp 1.1. text based/blocking call
http 2.0. binary based/non-blocking call



NOw, since goroutines are managed by the Go runtime, there is no one to control their access to the shared memory
It might be posible that their might be 5 goroutines trying to access the same memory location at the same time, which is problematic
This is where Mutexes come into play [Mutex exclusive lock]
Mutex provides a lock on the memory, till one goroutine is done with it, the other goroutine cannot access that particular memory block
Mutex() - lock, unlock
RWMutex() -> it will allow multiple readers, but as soon as anyones come to write on it, it wont allow any other reader or writer to access it

go run --race .


OAuth2 is an open standard for access delegation commonly used to grant websites or applications, limited access to user information without exposing the user's credentials
https://www.oauth.com/playground/





run multiple fgo rotuitnes on same memoery location without mutex
panic and recover use



appliication level and request level context
context.Background() -> application level context

context cis like a lifecycle of a process -> main thread or request
ctrl + c-> closing context of main thread

mongodb _id is a binary field
bson package of ggolang converts it into human redeable string format

equivalent to this? 
same as gin context?



  router := mux.NewRouter()
  router.HandleFunc("/test", TestHandler)
  http.Handle("/v1", router)
  router2 := mux.NewRouter()
  router2.HandleFunc("/test", TestHandler)
  http.Handle("/v2", router2)



Garbage Collection in GO is a process in which GO runtime [Go Schedular + Garbage Collector] automatically finds and frees up the memory 
that is no lomger in use helping to prevent memory leaks and optimize the performance of the program
GC uses the Mark and Sweep algorithm to find and free up the memory that is no longer in use
-> in mark phase, identifies all objects that are still reachable by the program, i.e. they are referenced by pointers in the memory
-> in the sweep phase, once all live objects are marked, garbage collector sweeps through the memory and frees up the memory that is not marked and used by unreachable or dead objects

Go maps objects into short lived and long lived and touches long lived objects less often, and reclaims memory from short lived objects more often
    -> this approach minimizes rhe impact of garbage collection on the performance of the application
Variables used only inside functions can be seen as short lived as they will get out of scope as soon as the function ends

Escape Analysis is a process in which the Go compiler determines whether a variable should be allocated on the heap or the stack
    -> if the variable is used only inside the function, it is allocated on the stack
    -> if the variable is used outside the function, it is allocated on the heap
    -> this is done to optimize the performance of the program and minimize the impact of garbage collection on the performance of the program

-> Function will get inside stack; stack will look like [ main->NewDuck->Reference to Duck struct ], when the function ends, the stack will be popped and the memory will be freed
    -> but the &Duck{} will be allocated on the heap, and will be freed only by the garbage collector
    func main() {
        NewDuck()
    }

    type Duck struct {}

    func NewDuck() *Duck {
        return &Duck{}
    }

    The function is named NewDuck by convention. It returns a pointer to a new instance of the Duck struct.
    &Duck{} creates a new instance of Duck in memory and returns its address (a pointer).

--------------------------------------------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------------------------------------------------------------------------------

    func main() {
        NewDuck()
    }

    type Duck struct {}

    func NewDuck() Duck {
        return Duck{}
    }
    -> everything here will be allocated in the stack and easily removed from there, so no need of garbage collection; hence the cheapest operation

In Escape Analysis, the Go compiler determines whether a variable should be allocated on the heap or the stack, it tries to keep maximum variables in stack

sync.Pool is a way to reuse objects in Go, it is a pool of objects that can be reused by the program
-> It caches objects even if they are dead or alive, the garbage collector will not mark them for cleanup immediately
-> so if we are creating soon-to-be-dead objects, we can use sync.Pool to reuse them and save memory and time

GOGC is a way to set the garbage collection percentage in Go, it is a percentage of the CPU time that is spent on garbage collection


--------------------------------------------------------------------------------------------------------------------------------


func (u *User) updateemail(str string) {
    u.email = email
}

func updateemail(u *User, str string) {
    u.email = email
}


user := User {
    name: "Aryan",
    email: "ss@gmail.com"
}

user.updateemail("new@gail.com)
user.updateemail(&user, "new@gail.com)


--------------------------------------------------------------------------------------------------------------------------------

Channels and Context are passed by reference in Go, which means that they are passed as pointers to the memory location where they are stored
So no need to explicitly pass them as pointers, they are already passed as pointers internally

fmt.Println(runtime.NumCPU()) // 14
The runtime.Numcpus function can be used to get the the number of logical processors available to the GO program. See below program

1. What is Cooperative Scheduling?
In cooperative (non-preemptive) scheduling, a running task (goroutine) does not get forcibly stopped by the scheduler. Instead, it must voluntarily yield execution at specific points, allowing the scheduler to decide whether to switch to another goroutine.


There is no difference in behaviour though when calling a anonymous function using goroutine or calling a normal function using goroutine
go func(){
   //body
}(args..)



Empty select
Select block without any case statement is empty select. The empty select will block forever as there is no case statement  to execute. It is also one of the way for a goroutine to wait indefinitely. But if this empty select is put in the main goroutine then it will cause a deadlock. Let's see a program

package main
func main() {
    select {}
}
-> In the above program we have empty select and hence it results in a deadlock that is why you see the output as below
































*/