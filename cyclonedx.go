package main

import "regexp"

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
	Alg     string `json:"alg"`
	Content string `json:"content"`
}

func NewCycloneFromNix(n *nix) (*cycloneDx, error) {
	c := cycloneDx{
		BomFormat:   "CycloneDX",
		SpecVersion: "1.3",
		Version:     1,
		Components:  []cycloneComponent{},
	}
	verRe := regexp.MustCompile("^.*-([0-9\\.]+)(\\.tar\\.gz|\\.tar\\.bz2)$")

	for _, entry := range *n {
		if entry.Env == nil {
			continue
		}
		cc := cycloneComponent{}
		if name, ok := entry.Env["name"]; ok {
			cc.Name = name
			verMatch := verRe.FindStringSubmatch(name)
			if verMatch != nil && len(verMatch) >= 2 {
				cc.Version = verMatch[1]
			}
		}
		cc.Type = "file"
		c.Components = append(c.Components, cc)
	}

	return &c, nil
}
