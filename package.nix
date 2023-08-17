{ buildBazelPackage, lib, stdenv, pkgs }:
buildBazelPackage {
  name = "gophrland";

  bazel = pkgs.bazel_6;
  bazelTargets = ["//cmd/gophrland"];
  bazelFlags = ["--explain=nixbuild.log" "--verbose_explanations"];

  nativeBuildInputs = with pkgs; [nix go];
  buildInputs = with pkgs; [cacert];

  buildAttrs = {
    installPhase = ''
      install -Dm755 bazel-bin/cmd/gophrland/gophrland_/gophrland $out/bin/gophrland
    '';
  };

  src = pkgs.nix-gitignore.gitignoreSource [] (lib.cleanSource ./.);
  fetchAttrs.sha256 = {
    x86_64-linux = "sha256-245DbmVlLVumOw6oNS3IcOdBbREgDPRfmlnw8aXrX3U=";
    aarch64-linux = "";
  }.${stdenv.hostPlatform.system} or (throw "unsupported system ${stdenv.hostPlatform.system}");
}

