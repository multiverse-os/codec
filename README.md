[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse: `codec` library
**URL** [multiverse-os.org](https://multiverse-os.org)

#### Introduction 
Codec provides a standardized way to marshal and unmarhsal data from Go structs
to common data ecoding formats, such as `bson`, `json`, `cbor`, and others. In
addition to providign an easy to use interface for the various common encoding
types, it provides incredibly simple way to handle standard compression
algorithms, such as `gzip`, `zstd`, `snappy`, and others. 

The idea is to keep this aspect of the code as simple as possible, abstracting
away the difficulty and allowing any extensibility to support the newest cutting
 edge data encoding, or compression algorithms to obtain gains easily, and with
 essentially no changes to codebases using this library. 


