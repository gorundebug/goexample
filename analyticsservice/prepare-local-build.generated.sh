#!/bin/sh
set -eu

source_dir="${1:?source directory is required}"
work_dir="${2:?work directory is required}"

test -f "$source_dir/go.mod"
mkdir -p "$work_dir"
find "$work_dir" -mindepth 1 -maxdepth 1 -exec rm -rf -- {} +
cp -a "$source_dir/." "$work_dir/"

cd "$work_dir"
go mod edit -replace github.com/gorundebug/servicelib=/servicelib
go mod edit -replace github.com/gorundebug/model=/modules/model