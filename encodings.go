package main

import (
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"catinello.eu/base91"
	"github.com/btcsuite/btcutil/base58"
)

// "encodingName": {"alias1", "alias2"...}
var encodings = map[string][]string{
	"ascii85":      {"a", "base85", "b85", "85"},
	"base32":       {"b32", "32"},
	"base32hex":    {"b32h", "h32"},
	"base58":       {"b58", "58"},
	"base64":       {"b64", "64"},
	"base64url":    {"b64u", "u64"},
	"base64raw":    {"b64r", "r64"},
	"base64rawurl": {"b64ru", "ru64"},
	"base91":       {"b91", "91"},
	"hex":          {"h", "base6", "b6", "6"},
}

func listEncodings() []string {
	var list []string
	for full, others := range encodings {
		list = append(list, full)
		list = append(list, others...)
	}
	return list
}

func findFullName(s string) (string, error) {
	for full, others := range encodings {
		if s == full {
			return full, nil
		}
		for _, name := range others {
			if name == s {
				return full, nil
			}
		}
	}

	return "", fmt.Errorf("unknown encoding: %s", s)
}

func encode(name string, b []byte) (string, error) {
	name, err := findFullName(name)
	if err != nil {
		return "", err
	}

	switch name {
	case "ascii85":
		var dst = make([]byte, ascii85.MaxEncodedLen(len(b)))
		ascii85.Encode(dst, b)
		return string(dst), nil
	case "base32":
		return base32.StdEncoding.EncodeToString(b), nil
	case "base32hex":
		return base32.HexEncoding.EncodeToString(b), nil
	case "base58":
		return base58.Encode(b), nil
	case "base64":
		return base64.StdEncoding.EncodeToString(b), nil
	case "base64url":
		return base64.URLEncoding.EncodeToString(b), nil
	case "base64raw":
		return base64.RawStdEncoding.EncodeToString(b), nil
	case "base64rawurl":
		return base64.RawURLEncoding.EncodeToString(b), nil
	case "base91":
		return string(base91.Encode(b)), nil
	case "hex":
		var dst = make([]byte, hex.EncodedLen(len(b)))
		hex.Encode(dst, b)
		return string(dst), nil
	}

	return "", fmt.Errorf("unknown encoding: %s", name)
}

func decode(name, s string) ([]byte, error) {
	name, err := findFullName(name)
	if err != nil {
		return nil, err
	}

	switch name {
	case "ascii85":
		var dst = make([]byte, ascii85.MaxEncodedLen(len(s)))
		_, _, err := ascii85.Decode(dst, []byte(s), true)
		return dst, err
	case "base32":
		return base32.StdEncoding.DecodeString(s)
	case "base32hex":
		return base32.HexEncoding.DecodeString(s)
	case "base58":
		return base58.Decode(s), nil
	case "base64":
		return base64.StdEncoding.DecodeString(s)
	case "base64url":
		return base64.URLEncoding.DecodeString(s)
	case "base64raw":
		return base64.RawStdEncoding.DecodeString(s)
	case "base64rawurl":
		return base64.RawURLEncoding.DecodeString(s)
	case "base91":
		return base91.Decode([]byte(s)), nil
	case "hex":
		var dst = make([]byte, hex.DecodedLen(len(s)))
		_, err := hex.Decode(dst, []byte(s))
		return dst, err
	}

	return nil, fmt.Errorf("unknown encoding: %s", name)
}
