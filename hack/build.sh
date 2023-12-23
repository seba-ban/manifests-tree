#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# kudos to https://github.com/kubernetes-sigs/kustomize/blob/master/releasing/create-release.sh
# Build the release binaries for every OS/arch combination.
# It builds compressed artifacts on $release_dir.
function build_binaries() {
  echo "build binaries"
  version=$1

  release_dir=$2
  echo "build release artifacts to $release_dir"

  mkdir -p "output"
  # build date in ISO8601 format
  build_date=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
  for os in linux darwin windows; do
    arch_list=(amd64 arm64)
    if [ "$os" == "linux" ]; then
      arch_list=(amd64 arm64 s390x ppc64le)
    fi
    for arch in "${arch_list[@]}" ; do
      echo "Building $os-$arch"
      binary_name="manifests-tree"
      [[ "$os" == "windows" ]] && binary_name="manifests-tree.exe"
      CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build \
        -o output/$binary_name \
        -ldflags "-s -w -X github.com/seba-ban/manifests-tree/cmd.version=${version}" \
        main.go
      if [ "$os" == "windows" ]; then
        zip -j "${release_dir}/manifests_tree_${version}_${os}_${arch}.zip" output/$binary_name
      else
        tar cvfz "${release_dir}/manifests_tree_${version}_${os}_${arch}.tar.gz" -C output $binary_name
      fi
      rm output/$binary_name
    done
  done

  # create checksums.txt
  pushd "${release_dir}"
  for release in *; do
    echo "generate checksum: $release"
    sha256sum "$release" >> checksums.txt
  done
  popd

  rmdir output
}

mkdir -p "$2"
build_binaries "$1" "$2"