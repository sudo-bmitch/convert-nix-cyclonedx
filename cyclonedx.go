package main

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	defCycloneType = "application"
)

type cycloneHashAlg int

const (
	cycloneHashNone cycloneHashAlg = iota
	cycloneHashMD5
	cycloneHashSHA1
	cycloneHashSHA256
	cycloneHashSHA384
	cycloneHashSHA512
	cycloneHashSHA3256
	cycloneHashSHA3384
	cycloneHashSHA3512
	cycloneHashBlake2b256
	cycloneHashBlake2b384
	cycloneHashBlake2b512
	cycloneHashBlake3
)

type cycloneDx struct {
	BomFormat   string             `json:"bomFormat"`
	SpecVersion string             `json:"specVersion"`
	Version     int                `json:"version"`
	Components  []cycloneComponent `json:"components"`
}

type cycloneComponent struct {
	Type    string          `json:"type"`
	Name    string          `json:"name"`
	Version string          `json:"version"`
	Hashes  []cycloneHashes `json:"hashes,omitempty"`
}

type cycloneHashes struct {
	Alg     cycloneHashAlg `json:"alg"`
	Content string         `json:"content"`
}

func NewCycloneFromNix(n *nix) (*cycloneDx, error) {
	var err error
	c := cycloneDx{
		BomFormat:   "CycloneDX",
		SpecVersion: "1.3",
		Version:     1,
		Components:  []cycloneComponent{},
	}
	verRe := regexp.MustCompile("^.*-([0-9\\.]+)(\\.tar\\.gz|\\.tar\\.bz2|)$")

	for _, entry := range *n {
		if entry.Env == nil {
			continue
		}
		cc := cycloneComponent{}
		if pname, ok := entry.Env["pname"]; ok {
			cc.Name = pname
		}
		if version, ok := entry.Env["version"]; ok {
			cc.Version = version
		}
		if name, ok := entry.Env["name"]; ok && (cc.Name == "" || cc.Version == "") {
			if cc.Name == "" {
				cc.Name = name
			}
			verMatch := verRe.FindStringSubmatch(name)
			if cc.Version == "" && verMatch != nil && len(verMatch) >= 2 {
				cc.Version = verMatch[1]
			}
		}
		if cc.Name == "" {
			// should this warn?
			continue
		}
		hash, okHash := entry.Env["outputHash"]
		alg, okAlg := entry.Env["outputHashAlgo"]
		var cAlg cycloneHashAlg
		if okAlg {
			err = (&cAlg).UnmarshalText([]byte(alg))
		}
		if okHash && okAlg && err == nil {
			cc.Hashes = []cycloneHashes{
				{
					Alg:     cAlg,
					Content: hash,
				},
			}
		}
		cc.Type = defCycloneType
		c.Components = append(c.Components, cc)
	}

	return &c, nil
}

func (alg cycloneHashAlg) MarshalText() ([]byte, error) {
	var s string
	switch alg {
	default:
		s = ""
	case cycloneHashMD5:
		s = "MD5"
	case cycloneHashSHA1:
		s = "SHA-1"
	case cycloneHashSHA256:
		s = "SHA-256"
	case cycloneHashSHA384:
		s = "SHA-384"
	case cycloneHashSHA512:
		s = "SHA-512"
	case cycloneHashSHA3256:
		s = "SHA3-256"
	case cycloneHashSHA3384:
		s = "SHA3-384"
	case cycloneHashSHA3512:
		s = "SHA3-512"
	case cycloneHashBlake2b256:
		s = "BLAKE2b-256"
	case cycloneHashBlake2b384:
		s = "BLAKE2b-384"
	case cycloneHashBlake2b512:
		s = "BLAKE2b-512"
	case cycloneHashBlake3:
		s = "BLAKE3"
	}
	return []byte(s), nil
}

func (alg *cycloneHashAlg) UnmarshalText(b []byte) error {
	switch strings.ToLower(string(b)) {
	default:
		return fmt.Errorf("Unknown hash algorithm \"%s\"", b)
	case "":
		*alg = cycloneHashNone
	case "md5":
		*alg = cycloneHashMD5
	case "sha-1", "sha1":
		*alg = cycloneHashSHA1
	case "sha-256", "sha256":
		*alg = cycloneHashSHA256
	case "sha-384", "sha384":
		*alg = cycloneHashSHA384
	case "sha-512", "sha512":
		*alg = cycloneHashSHA512
	case "sha3-256":
		*alg = cycloneHashSHA3256
	case "sha3-384":
		*alg = cycloneHashSHA3384
	case "sha3-512":
		*alg = cycloneHashSHA3512
	case "blake2b-256":
		*alg = cycloneHashBlake2b256
	case "blake2b-384":
		*alg = cycloneHashBlake2b384
	case "blake2b-512":
		*alg = cycloneHashBlake2b512
	case "blake3":
		*alg = cycloneHashBlake3
	}
	return nil
}
