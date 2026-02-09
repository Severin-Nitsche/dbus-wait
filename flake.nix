{
  description = "D-Bus wait";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { self, nixpkgs }:
  let
    pkgs = nixpkgs.legacyPackages.x86_64-linux;
  in {

    packages.x86_64-linux = {
      default = self.packages.x86_64-linux.dbus-wait;
      dbus-wait = pkgs.callPackage ./package.nix {};
    };

  };
}
