# Go

[GO DOC SPECS](https://go.dev/ref/spec)

### keywords
The following keywords are reserved and may not be used as identifiers.
```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### Operators
```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~
```

### The zero-value

Each element of such a variable or value is set to the zero value for its type:  
```
false for booleans  
0 for numeric types   
"" for strings  
nil for pointers, functions, interfaces, slices, channels, and maps  
```

