{ lib, buildGoModule }:

let
  fs = lib.fileset;
  sourceFiles = fs.unions [
    ./main.go
    ./go.mod
    ./go.sum
  ];
in buildGoModule {
  pname = "dbus-wait";
  version = "1.0.0";

  src = fs.toSource {
    root = ./.;
    fileset = sourceFiles;
  };

  vendorHash = "sha256-y0H6OoxXH+pOEiIzPeSUEAGukrLK0ioRcQXAwnR/J94=";

  meta = {
    description = "Blocks until the specified D-Bus event is fired";
    license = lib.licenses.mit;
    platforms = lib.platforms.linux;
  };
}
