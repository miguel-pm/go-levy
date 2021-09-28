{ sources ? import ./sources.nix
, pkgs ? import ./pkgs.nix { inherit sources; }
}:
let
  pre-commit-hooks = import sources.pre-commit-hooks;
in
{
  run = pre-commit-hooks.run {
    src = ../.;
    hooks = {
      nixpkgs-fmt.enable = true;
      shellcheck.enable = true;
      go-lint = {
        enable = true;
        name = "go linter";
        # The command to execute
        entry = "${pkgs.golint}/bin/golint";
        # Pattern of files to run on
        files = "\\.(go)$";
      };
      go-format = {
        enable = true;
        name = "go formatting checks";
        # The command to execute
        entry = "${pkgs.go}/bin/gofmt -l -w";
        # Pattern of files to run on
        files = "\\.(go)$";
      };
    };
  };
  tools = with pre-commit-hooks; [
    nixpkgs-fmt
    shellcheck
  ];
}

