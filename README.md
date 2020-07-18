# Clockwork Base32 for Go

Implementation of Clockwork Base32 for Go.

Clockwork Base32 is a simple variant of Base32 inspired by Crockford's Base32.
See [Clockwork Base32 Specification](https://gist.github.com/szktty/228f85794e4187882a77734c89c384a8)

## Usage

```
import "github.com/szktty/go-clockwork-base32"

encoded := clockwork.Encode([]byte("Hello, world!"))
decoded, err := clockwork.Decode(encoded)
if err != nil {
    fmt.Printf("decode failed => %s\n", err)
}
fmt.Printf("decoded => %s\n", decoded)
```
