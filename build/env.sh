#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
sccdir="$workspace/src/github.com/meachon"
if [ ! -L "$sccdir/go-socicoin" ]; then
    mkdir -p "$sccdir"
    cd "$sccdir"
    ln -s ../../../../../. go-socicoin
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
GO15VENDOREXPERIMENT=1
export GOPATH GO15VENDOREXPERIMENT

# Run the command inside the workspace.
cd "$sccdir/go-socicoin"
PWD="$sccdir/go-socicoin"

# Launch the arguments with the configured environment.
exec "$@"