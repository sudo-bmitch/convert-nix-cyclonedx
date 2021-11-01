# Convert Nix CycloneDX

This converts a Nix derivation to a CycloneDX SBoM

## Run with nix2.4 (with flakes)
```
nix show-derivation .# --recursive | nix run .#
```
