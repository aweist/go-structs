## Utility to find and output all golang struct definitions within a directory

Use the binary like so:

```shell
go-structs {{directory}}
```

Example

```shell
go-structs /code/go-structs/test_dir
```

Output to file:

```shell
go-structs /code/go-structs/test_dir > /tmp/output.txt
```

The program will only report on files that contain go structs.
The output will print a filename followed by a newlin-delimited list of the structs in that file. 
Example output:

```shell
/code/go-structs/test_dir/file1.go

type Struct1 struct {
	Field1 string
	Field2 int
}

type Struct2 struct {
	Field1 string
	Field2 int
}
```
