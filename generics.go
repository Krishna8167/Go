package main

/*
Go doesn't support classes,
For long time, that meant that Go code couldn't easily be reused in many
circumstances.
For example, imagine some code that splits a slice into 2 equal parts. the code
that splits the slice doesn't really care about the value stored in the slice.
Unfortunately in Go we would need to write it multiple times for each type, which is a very un-DRY thing to do
(give example...)

Why generics:
it reduces repetitive code.constare used more often in libraries and packages.

Constraints :
(constraints examples... )

interface type lists :
when generics were released, a new way of writing interfaces was also released at the same time.

We can now simply list a bunch of types to get a new interface  constraint.
(example .. )

we would use it if , you know exactly which type satisfy your interface.

Parametric Constraints :
Your interface definitions, which can later be used as constraints, can accept type parameters as well.
(give example .. )

Naming generics type :

func splitAnySlice[T any](s []T) ([]T, []T) {
mid := len(s)/2
return s[:mid], s[mid:]
}

T -> is just a variable name, we could have named the type parameter anything. T is fairly common convention
for type variable, similar to how i is a ocnvention for index variables in loops.

*/
