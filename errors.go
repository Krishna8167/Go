package main

// GO programming expresses error with error values
/* an error is any type that that implements the simple built-in error interface

type error interface {
 Error() string
 }

 if something go wrong in the function, that function should return an error as its last return value
 Any code that calls a function that can return any error should handle errors by testing whether the error is 'nil'

 //Atoi -> convertes stringified number to int
 //strconv -> string converter

     i, err := strconv.Atoi("42b")
    if err != nil {
    fmt.Println("couldn't convert:", err)

      return
     }

// BECAUSE ERROR are just interfaces , you can build your own custom types that implements the error interface.

type userError struct {
name string
}

func (e UserError) Error() string {
return fmt.Sprintf("%v has a problem with their account",e.name)
}

func sensSMS(msg, UserName string) error {
 if !canSendToUser(userName) {
return userError{name:userName}
}
}

// Error package

the go standard library provides an "errors" package that makes it easy to deal with errors.

var err error = errors.New("Something went wrong")

ex:  func divide (x, y float64) (float64, error) {
 if y == 0 {
return 0.0, errors.New("no divison by 0")
 }
return x/y , nil
}

*/
