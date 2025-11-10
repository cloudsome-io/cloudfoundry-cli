#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="${SCRIPT_DIR%/scripts}"
DIST_DIR="${ROOT_DIR}/dist"

TARGETS=(
    "linux:amd64"
    "linux:arm64"
    "darwin:amd64"
    "darwin:arm64"
    "windows:amd64"
)

usage() {
    cat <<EOF
Usage: $(basename "$0") [all|<os>:<arch> ...]

Examples:
  $(basename "$0")
  $(basename "$0") all
  $(basename "$0") linux:amd64 windows:amd64
EOF
}

ensure_dist_dir() {
    mkdir -p "$DIST_DIR"
}

build_target() {
    local os="$1"
    local arch="$2"
    local suffix=""

    if [[ "$os" == "windows" ]]; then
        suffix=".exe"
    fi

    local output="${DIST_DIR}/cs-cli-${os}-${arch}${suffix}"

    echo "==> Building ${os}/${arch} -> ${output}"
    CGO_ENABLED=0 GOOS="$os" GOARCH="$arch" go build -o "$output" "$ROOT_DIR"
}

main() {
    if [[ ${1:-} == "-h" || ${1:-} == "--help" ]]; then
        usage
        exit 0
    fi

    ensure_dist_dir

    local targets=("${TARGETS[@]}")

    if [[ $# -gt 0 && $1 != "all" ]]; then
        targets=("$@")
    fi

    for target in "${targets[@]}"; do
        if [[ "$target" == "all" ]]; then
            continue
        fi

        IFS=":" read -r os arch <<< "$target"
        if [[ -z "$os" || -z "$arch" ]]; then
            echo "Invalid target format: $target" >&2
            usage
            exit 1
        fi

        build_target "$os" "$arch"
    done
}

main "$@"

