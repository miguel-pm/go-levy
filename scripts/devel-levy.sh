#!/bin/sh


# This script sets up a feedback loop that runs the unit tests of the project
# and restarts them when a change happens in the ./src folder.

current_directory=$(basename $(pwd))

if [ $current_directory != "go-levy" ]; then
	echo "ERROR: The current working directory is not set correctly -> $(pwd)"
	echo "to correctly run this script do it from the root of the project."
	exit 1
fi

echo "Setting up feedback loop with UNIT tests"

nix-shell --run "\
	watchexec --clear --watch './levy' 'cd levy && go test ./...'"
