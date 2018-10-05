# Go quick reminders

- C-style comments, brackets, names (suchAsThisOne) and ... pointers!
- stdout -> ``fmt.Println("Hello")``
- Functions stating with an upper case letter are exportable
- ``func name(paramsNames type) (returnsNames type){}``
- ``var x int``  or ``x:= 1``
- lots of var types (even ``complex128``: c:= 3 + 4i !)
- Arrays Versus Slices (~ static versus dynamic)
- ``make`` to allocate memory : ``slice := make([]int, 4)`` (block of 4 ints)
- still on slices : ``slice = append(slice, 21, 22, 13)``
- still on slices : ellipsis at the end of a slice literal means "unpack its elements"
- Hashs/dicts done with ``map``
- Write in a file via ``fmt.Fprint(file, "Things")``
- ``switch-case`` exists in Go. ``fallthrough`` keywork to continue on subsequent case
- ``while`` doesn't exist
- ``for`` can be used on a C-style (``for x := 0; x < 10; x++ {}``) 
or on a Python-style (``for key, value := range``)
- As in ``for``, it is possible to assign in ``if`` statement (``if y := 10; y > x {}``)
- Can assign a function directly to a variable (``oneFunc := func(){return}``)
- ``goto`` exists !  (jump to a "tag" not indented like ``whateverName:``)
- Kind of decorators also exists (func with func parameter)
- ``defer`` to execute code just before the end (to close file for example)
- Object-like can be approached with Interfaces & Structs
- Multiple and variable parameters can be passed to a func with ``...``
  : ``func OneFunc(params ...int). The func receives the params as a Slice
- Basic error handling : ``if x, err := test(); err != nil { fmt.Println(err) } else {}``
- Concurrency with Channels (communication ``<-``) & goroutines (``go``)
- ``Select-case`` statements is like switch but for Channels
- https://golang.org/doc/
- https://golang.org/pkg/
- https://golang.org/src/
