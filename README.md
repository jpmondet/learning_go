# Go quick reminders

- C-style comments, brackets, names (suchAsThisOne) and ... pointers!
- vars & const can be written like imports (with parenthesis)
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
- ``sync.WaitGroup`` can be usefull to wait all goroutines
- https://golang.org/doc/
- https://golang.org/pkg/
- https://golang.org/src/

# Some warnings :
- Variable assignment with ``:=`` works only inside a func
- Shadow variables can be painful. Help : ``go tool vet -shadow your_file.go``
  or ``go-nyet``
- Adding an item (``m['first'] = 1``) to a ``nil`` map (declared with ``var m map[string]int``)
  throws an error
- cap() won't work on a map already initialized
- Unlike in ``C``, Arrays are not pointers
- Unlike ``Python``,  ``range`` always returns an index
- Checking if a key exists on a map can be done with ``if _ ,ok := m['k']; !ok
  {}
- Go is not as easy as ``Python`` for ``Strings`` : 
  - Can't modify a string (immutable). Can convert it to bytes slice instead ``[]byte(s)``
  - On a string, s[2] won't return a char. It will return a byte value.
  - ``len(string)`` returns the number of bytes, not the number of chars...
  - To be encoded, a struct NEEDS to be exportable (upper case first letter)
  - Sending to close channels causes panic
  - ``nil`` channels block infinitely
  - ``http library`` : Even empty body must be closed (/!\ the response must
    not be ``nil``)
  - JSON encoder adds a newline char
  - ``==`` can't compare every types. ``reflect.DeepEqual`` or ``bytes.Equal`` can help.

