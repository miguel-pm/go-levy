let
  sources = import ./nix/sources.nix;
  pkgs = import ./nix/pkgs.nix { inherit sources; };
  pre-commit = import ./nix/pre-commit.nix { inherit sources; inherit pkgs; };
in
pkgs.mkShell {
  name = "go-levy-shell";
  buildInputs = with pkgs; [
    (import sources.niv { }).niv
    go
    golint
    watchexec
  ] ++ pre-commit.tools;
  shellHook = pre-commit.run.shellHook;
}
