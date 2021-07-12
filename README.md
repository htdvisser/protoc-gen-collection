# Collection of Protocol Buffers Generator Plugins

This repository contains a collection of plugins for protoc.

## JSON and YAML files

`protoc-gen-json-files` and `protoc-gen-yaml-files` generate files that describe the messages, enums and services in your proto files.

```
$ protoc -I [your imports ...] \
  --json-files_out=output_path=path/to/data:path/to/data \
  /path/to/*.proto
```
