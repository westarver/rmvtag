# tagit
> Overview -- tagit is a multi-tool project that
> consists of three primary tools.  proptag will 
> read go source files and generate getters and/or
>setters for basic variable and field types or
> tags can be configured to generate functions that
>require a Set or Get method on more complicated
> user defined types. tags for this tool look like 
> this: //<pgs> </prop> (open and close blocks of 
> variables or fields) or //<pgs/> at the end of a 
> line to affect that line only.
> rmvtag will comment out, uncomment or remove source
> code lines tagged with the appropriately configured 
> tags as so: <rmv>, </rmv> for blocks, or <rmv/> for lines.
> A couple of examples; first proptags:
```go
  	var intvar int //<pg/> makes the unexported
	// variable intvar accessible to outside code 
  	// as a read-only property.
  	type myType struct {
  		// requires myType to have both Get() and
		// Set(structType) methods
  		fld structType //<pGS/>
  	}
```
