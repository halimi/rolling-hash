# rollhash
Rolling hash based file diffing tool

An [rdiff](https://librsync.github.io/page_rdiff.html) like tool that can produce 'deltas' between the original and the new file.
To compare the two files not necessary to have the original file. `rollhash` can create a 'signature' file from the original file. This can be used to create the 'deltas'.

# Usage

`rollhash` has the same API as `rdiff`

```bash
Usage:
  rollhash signature [BASIS [SIGNATURE]]
  rollhash delta SIGNATURE [NEWFILE [DELTA]]

```

## Signature

`signature` generates a signature from the base file. Signature describes the hash of the chunks in `map[string][]int` format, where the key is the hash and value is the list of the chunks. This map also contains the chunk size used during the generation. The chunk size depends on the input data length, between the min and max limit value.

```bash
parameters:
  * BASIS - input file path, mandatory parameter
  * SIGNATURE - output file path, optional parameter, if not provided then it prints out to the standard out.
```

## Delta

`delta` generates the diff between the original file and the new file. The 'deltas' shows which chunk has been changed in a `map[int][]byte` format. Where the key is the chunk index and the value the new data. This delta can be used to apply on the original file to get the new file.

```bash
parameters:
  * SIGNATURE - signature file path, mandatory parameter
  * NEWFILE - input file path, mandatory parameter
  * DELTA - output file path, optional parameter, if not provided then it prints out to the standard out.
```

# How to build

You can build `rollhash` with the standard go build command.

```bash
go build -o rollhash cmd/rollhash/*.go
```
