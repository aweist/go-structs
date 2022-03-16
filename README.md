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