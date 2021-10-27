# Protobuff 101

Purpose: Learning protobuf with Golang

### Proto Compiler

- For protobuffs that are compiled for Golang, there needs to be the following line in the *.proto file
```protobuf
option go_package = "some_package/path";
```

- The command line argument to create the proto file

```zsh
protoc --go_out=. --go_opt=paths=source_relative model/book.proto
```

### Learnings

- Protoc creates a `proto.marshall` similar to `json.marshall`

The contents of the hexdump can be examined using the following command

```zsh
hexdump -C book.protobuf
```

A comparison to json output was also done with


```zsh
hexdump -C book.json
```

Unsurprising, protobuf produce smaller sized files

### Decoding

Decode by
- Using the model from the generated `book.pb.go` file and reading the bytes with ioutil.ReadFile
- Match to the model in the pb.go file
