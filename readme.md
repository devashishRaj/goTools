> You have learnt GO, now what can you do with it ?


## Packages 
>All design decisions start and end with the package. The purpose of a package is to provide a 
>solution to a specific problem domain. To be purposeful, packages must provide, not contain.  
>The more focused each package’s purpose is, the more clear it should be what the package provides.

> A module in Go is a collection of (usually related) packages. 
- Good programmers, then, are always thinking in terms of writing importable packages, not mere dead‐end programs.
- If we can figure out how to break down our unsolved problems into a bunch of mini‐ problems that have  already been solved by existing packages, then we’re 90% done.
- pkg.go.dev site lets you search and browse the whole universal library
- to tackle some specific problem, a good place to check first is awesome‐go

>Writing a test forces you to think clearly and precisely about the observable behaviour of the component under test, rather than the irrelevant details of its implementation.

1. If we want to design a package, a great way to begin is by pretending it already exists, and writing code that uses it to solve our problem.

## PaperWork
>Designing packages instead of programs means we need to think about how other people might use our code.

- Mandatory arguments are annoying : hello.PrintTo(os.Stdout) in packages chapter, here mentioning  
os.Stdout is like __useless paperwork__.
    - we can make os.stdout default if argument passed is nil but this makes it worse but 
