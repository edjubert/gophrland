{ pkgs }:
pkgs.mkShellNoCC {
  nativeBuildInputs = with pkgs; [
    go
  ];
}
