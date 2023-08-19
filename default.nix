{ stdenv
, callPackage
, go
, lib
, makeWrapper
, installShellFiles
, fetchFromGitHub
, buildGoApplication
, mkGoEnv
}:

buildGoApplication {
  pname = "gophrland";
  version = "v0.0.8";

  modules = ./gomod2nix.toml;

  src = lib.cleanSourceWith {
    filter = name: type: builtins.foldl' (v: s: v && ! lib.hasSuffix s name) true [
      "tests"
      "builder"
      "templates"
    ];
    src = lib.cleanSource ./.;
  };

  inherit go;

  allowGoReference = true;

  subPackages = [ "./cmd/gophrland" ];

  nativeBuildInputs = [ makeWrapper installShellFiles ];

  passthru = {
    inherit buildGoApplication mkGoEnv;
  };

  postInstall = lib.optionalString (stdenv.buildPlatform == stdenv.targetPlatform) ''
    wrapProgram $out/bin/gophrland --prefix PATH : ${lib.makeBinPath [ go ]}
  '';

  meta = {
    description = "Gophrland is a set of tools to manage windows on Hyprland.";
    homepage = "https://github.com/edjubert/gophrland";
    license = lib.licenses.mit;
    maintainers = [ lib.maintainers.edjubert ];
  };
}