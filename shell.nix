{ pkgs }:
pkgs.mkShellNoCC {
  nativeBuildInputs = with pkgs; [
    go
    gcc
    alejandra
    bazel_6
    bazel-buildtools
  ];
}
