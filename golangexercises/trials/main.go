package main

//HOW CAN IN CONVERT ALL THESE COMMENT DOCUMENTATION TO READABLE AND EDITABLE TEXT INSTEAD OF COMMENTS
//Golang composition and reflection

//you can only get the keys of your map using a for loop range or you use reflection
//a struct can have no state but have behaviors

//referential transparential involves pure funcs(defined below) and is good for testing (unit),
// because the func can be called anytime without affecting outer environment/scope
//clousures can be made by returning a function that remebers state of its outer
// function scope or using func expressions inside a func that remembers scope of
// outer func
//higher order funcs - func that can receive or/and return functions as params/args
//and return value/statement respectively
//first class funcs - langs that treat func like other variables
//immutable data types
//NO SIDE EFFECTS
//persistent data storage/strutures
// pure function - input and output, expected input gives expected output,
// parameters and return statements, receive and return something withou changing a state, no side effects
//anonymous funcs - lambda, self executing, func literal, func expressions
// func express are achieved with func literals?
// recursing
// currying
//lazy evaluation
//chaining functions (part of func programming)?
//function composition/composition of funcs/pipelining?/collection pipeline?
//r transparency
//monad
// tail call optimization
// map , reduce. map collection to a function
//statements usually mutate state unlike expressions which returns a value
// functional-style programming encourages us to use expressions
//function composition and collection pipeline can be achieved with filter, map,
// and reduce funcs
//with traditional primite func comps, the code is read R-L e.g. describe(clarify("whiskey"));
//^clarify "whiskey" first then invoke describe func on returned value?
//with builtin func compositions like lambda's it is read L-R?

//see a recursion as a loop, so keep calling itself until a base case condition is met

//currying. Like a func requiring 3 args but you invoke the func with 1 arg each time bcos it
//remembers previous args. "taking a generic func and making a specific func out of it"
//function addNumbers(num1, num2) {
// 	return num1 + num2;
// }
// //R.curry is a package func?
// const curriedAdd = R.curry(addNumbers);
// const add4 = curriedAdd(4)
// add(2);//6
// add(7);//11

//A predicate func returns a boolean usually based on func name e.g isBlue(arg)

// Of course there are situations in which the use of mutable state is necessary, e.g. IO.
// Functional programming languages do allow the use of mutable state (via monads, impurity etc.)

// alternates to traditional (for) looping(for, for each, etc) - recursion, currying?, func composition,
// collection pipeline

//Method chaining
//e.g -> PizzaBuilder.AddSauce().AddDough().AddTopping() Builder pattern

// In programming language theory, lazy evaluation, or call-by-need is an evaluation
// strategy which delays the evaluation of an expression until its value is needed
//  (non-strict evaluation) and which also avoids repeated evaluations (sharing).

// In computer science, function composition (not to be confused with object
// 	composition) is an act or mechanism to combine simple functions to build more
// 	complicated ones. Like the usual composition of functions in mathematics,
// 	the result of each function is passed as the argument of the next(pipelining?), \
// 	and the result of the last one is the result of the whole.

// This pattern(func composition nd collection pipelining) is common in functional programming, and also in object-oriented
// languages which have lambdas e.g filter, map, reduce

// func main() {
// //func literal with no identifier(anonymous), also self executing
// 	func() {
// 		fmt.Println("Hello, playground")
// 	}()
// to define a func in a func use a func expresion i.e assign a func literal to a variable
// funcexp := func() {
// 	fmt.Println("Hello, playground")
// }

// funcexp()
// }

// A string is made up of runes, and a rune is a byte. So a string is a slice of bytes
//UTF-8 is 1-4 bytes coding scheme

//variables declared at package-level are accessible to modules/go files in the same package?
//To make the variables accessible to other pacakages, start them with an Uppercase letter?

//We have values, references, and pointers
//& applied on value, *applied on already declared identifier?
//Use * during definition/declaration/inialization, use & during call/implementation?
//* (get) pointer, & get memory address?
//& points to an address
// deference an address/pointer, to receive the value stored in the address? * (on) a &
//& stores the memory address of a  variable (in a variable)?
//In Java, references are basically pointers
// A pointer is a variable that holds memory address of another variable.
// A reference variable is an alias, that is, another name for an already existing variable. A reference, like a pointer is also implemented by storing the address of an object.
//A pointer can be reassigned, a reference can't
// stack vs heap
//threads each have their their own stack memory inside a process. the process has the heap memory

// polymorphism can be achieved by passing related objects(related by inheritance)? e.g a function taking a base class object as a parameter, so any derivative
// of the base object(child) can be passed as an argument (incl. the base object).
// it can also be achieved by passing an interface object to a func as a parameter, so any call to the func will be polymorphic depending on the argument passed
// which has to be an implementation of the interface

// //store a value in variable s. s would be assigned a memory address
// s := 7
// //stores the address of s in variable t. t's value is now the memory address of s
// t := &s
// //sice t is a memory address holding a particular value, we now print that value by deferencing the address
// fmt.Println(*t)
// //we print the address of s
// fmt.Println(&s)
// //we print the address of t which has a value similar to the address of s
// fmt.Println(t)
// fmt.Println(&t)

//A byte type is a sequence of runes
// []byte("abcd") //This actually converts the string "abcd" to a slice of bytes
// []byte{'a', 'b', 'c', 'd'}//A slice of byte is created with each rune as an element
// character, code point, rune. rune = code point?
// casting === conversion?
// fmt.Println(string(25105)) //我
// fmt.Println('我') // 25105

//byte        alias for uint8
//rune        alias for int32. Represents a unicode code point []rune has more bit representation than []byte
//Printing an index of a string returns its code point/rune(binary digit) in UTF8
//fmt.Prin
//Rune
// fmt.Println('我') //Output is the code point of the character

// type Int int
type IntSlice []int

func main() {

	// listOfNumber := []int{1, 2, 3, 4}
	listt := []int{1, 2}
	// print(listt)

	//normally after appending, new copy is generated so not
	// append(listOfNumber, 5)
	// memory := &listt

	// var neww = append(listt, 5)
	listt = append(listt, 5)

	//If you need to use both index(key?), and value of a data structure
	for i, j := range listt {
		println("At index", i, "is", j)
		println()
	}

	//If you only need to use value
	for k := range listt {
		println(listt[k])
	}

	//indexing a string at index i doesn;t return the character in that index
	// fmt.Println("A"[0])
	//although for other collections/data types?? the character at the index position is returned
	// arr := [2]string{"A","B"}
	// fmt.Println(arr[0])

	println("**************************")

	//Another way of doing above. Notice how body of function is more concise than above
	for _, l := range listt {
		println(l)
	}

	println("**************************")

	for m := 0; m < len(listt); m++ {
		println(listt[m])
	}

	// //while
	// for <expression> {

	// }

	// runCode(listt...)

	// print(*listOfNumber)
	// runCode(listOfNumber...)

}

func runCode(numbers ...int) {

	for _, number := range numbers {

		println(number)

	}

}

// func (s IntSlice) removeInt(n int) {
// 	capp := cap(s)
// 	neww := make([]int, 0, capp)

// 	tt := s
// 	s = s[n : n+1]

// 	for i, j := range s {
// 		if n != i
// 		neww = append(neww, j)

// 	}
//
// }

// //	fmt.Print([]byte{'我'})
// fmt.Println(string(25105))
// fmt.Printf("%T\n", string(25105))
// fmt.Println('我') //Output is b
// fmt.Println([]byte{'A'})
// fmt.Println([]byte{'ÿ'})
// fmt.Println([]rune{'A'})
// fmt.Println([]rune{'Ā'})
// fmt.Println([]byte("我")) //[230 136 145]
// fmt.Println([]rune("我")) //[25105]
// fmt.Println(byte('我')) // overflow error because 25105 character code exceeds capacity of uint8/byte
// fmt.Println(rune('我'))

// Strings blog post
// "...(Of course, bytes range from hexadecimal values 00 through FF, inclusive.)..."
// FF in Hex == 255 (byte/code point). byte is uint8, rune is int32
// 2^8 or A byte max value in decimal(byte code point) == 256

//Assertion with comma ok. Assertion applies to interface types. <name.(string)>
//var name interface{} = "Femi"
//str, ok := name.(string)
//OR REPLACE BELOW LINE WITH THIS LINE. if str, ok := name.(string); ok { //ok equals boolean value, if ok/assertion is true, the value is returned to str
// if ok {
// fmt.Printf("%T\n", str)
// } else {
// fmt.Printf("value is not a string\n")
// }

//if <initialization>, <condition> {<body>}

// var val interface{} = 7
//Assertion is performed on an interface below. Like performing non-interface type conversion/casting
//fmt.Println(val.(int) + 6)
//fmt.Println(val + 6) //wont work because val is of type interface int
