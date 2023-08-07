{ lib, buildGoModule, fetchFromGitHub }:
buildGoModule rec {
  pname = "gophrland";
  version = "v0.0.3.2";

  src = fetchFromGitHub {
    owner = "edjubert";
    repo = "gophrland";
    rev = "v${version}";
    sha256 = "";
  };

  proxyVendor = true;
  vendorHash = "";

  tags = [ "hyprland" "wayland" "gophrland" ];

  meta = with lib; {
    homepage = "https://github.com/edjubert/gophrland";
    description = "Gophrland is a set of tools to manage windows on Hyprland";
    mainProgram = "gophrland";
  };
}