{
  description = "convert-nix-cyclondx uses nix show-derivation --recursive to CycloneDX";

  outputs = { self, nixpkgs }: {

    packages.x86_64-linux.convert-nix-cyclonedx = nixpkgs.legacyPackages.x86_64-linux.buildGoModule {
      pname = "convert-nix-cyclonedx";
      version = "0.0.0";
      src = self;
      vendorSha256 = "sha256-pQpattmS9VmO3ZIQUFn66az8GSmB4IvYhTTCFn6SUmo=";
    };

    defaultPackage.x86_64-linux = self.packages.x86_64-linux.convert-nix-cyclonedx;

  };
}
