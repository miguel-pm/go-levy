{ sources ? import ./sources.nix }:
let
  pkgs = import sources.nixpkgs {
    overlays = [
      (final: previous: { inherit (import sources.gitignore { inherit (final) lib; }) gitignoreSource; })
    ];
  };
in pkgs
