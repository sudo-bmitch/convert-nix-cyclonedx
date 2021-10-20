package main

type nix map[string]nixEntry

type nixEntry struct {
	Outputs nixOutputs        `json:"outputs,omitempty"`
	Env     map[string]string `json:"env,omitempty"`
}

type nixOutputs struct {
	Out nixOutputsOut `json:"out"`
}

type nixOutputsOut struct {
	Path     string `json:"path"`
	HashAlgo string `json:"hashAlgo"`
	Hash     string `json:"hash"`
}
