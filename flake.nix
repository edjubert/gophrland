{
  description = "A very basic flake";

  outputs = { self, nixpkgs }: {

    packages.x86_64-linux.hello = nixpkgs.legacyPackages.x86_64-linux.hello;

    packages.x86_64-linux.default = 
      with import nixpkgs {system = "x86_64-linux";};
      stdenv.mkDerivation rec {
        name = "gophrlang-${version}";
        version = "0.0.3";

        src = fetchurl {
          url = "https://github.com/edjubert/gophrland/releases/download/v0.0.3/gophrland-v0.0.3-linux-x86_64.tar.gz";
          sha256 = "sha256-cNIloeBibs+XTFA46kYkkGfDwuykRmDDMVm1Nz5LvEU=";
        };

        sourceRoot = ".";

        installPhase = ''
          install -m 755 -D gophrland $out/bin/gophrland
        '';

        meta = with lib; {
          homepage = "https://github.com/edjubert/gophrland";
          description = "A set of tools to manage windows, workspaces, monitors and scratchpads on Hyprland";
          platforms = platforms.linux;
        };
      };
  };
}
