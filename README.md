# gunk

gunk is a command line utility for encoding and decoding stdin in various encodings/formats

```bash
go install github.com/notwithering/gunk
```

## examples

### encode text into base64

```bash
echo "hello world" | gunk base64
```

### decode base64 encoded text

```bash
echo "aGVsbG8gd29ybGQK" | gunk -d base64
```

### make base91 text into ascii85 text

```bash
echo "TPwJh>Io2Tv!^aB" | gunk -d base91 ascii85
```

### list availible encodings and their aliases

```bash
gunk --list-encodings
```

## supported encodings

- ascii85
- base32
- base32hex
- base58
- base64
- base64url
- base64raw
- base64rawurl
- base91
- hex
