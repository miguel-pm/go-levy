let
  sources = import ./nix/sources.nix;
  pkgs = import ./nix/pkgs.nix { inherit sources; };
  pre-commit = import ./nix/pre-commit.nix { inherit sources; inherit pkgs; };
in
pkgs.mkShell {
  name = "go-levy-shell";
  # Prevent change locale warning as suggested in https://gist.github.com/peti/2c818d6cb49b0b0f2fd7c300f8386bc3
  LOCALE_ARCHIVE_2_27 = "${pkgs.glibcLocales}/lib/locale/locale-archive";
  buildInputs = with pkgs; [
    (import sources.niv { }).niv
    go
    golint
    watchexec
  ] ++ pre-commit.tools;
  shellHook = pre-commit.run.shellHook;
}
