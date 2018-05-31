#!/usr/bin/env bash
ROOT=$(git rev-parse --show-toplevel)
MAIN_VERSION=$(git describe --always --abbrev=0 --tags || echo "0.1")
BUILD_PLATFORMS="linux windows darwin freebsd"
BUILD_ARCHS="amd64 386"
GO_BUILD_CMD="go build -installsuffix cgo -a"
if [ "$#" -ne 1 ] ; then
	echo "ERROR: ldflags needed"
	exit 1
fi

LDFLAGS="$1"

mkdir -p "${ROOT}/release"

for OS in ${BUILD_PLATFORMS[@]}; do
	for ARCH in ${BUILD_ARCHS[@]}; do

		NAME="acictl-${OS}-${ARCH}"
		if [[ "${OS}" == "windows" ]]; then
			NAME="${NAME}.exe"
		fi

		if [[ "${OS}" == "darwin" && "${BUILD_PLATFORM}" == "darwin" ]]; then
		  CGO_ENABLED=1
		else
		  CGO_ENABLED=0
		fi

		echo "Building for ${OS}/${ARCH} with CGO_ENABLED=${CGO_ENABLED}"

		GOARCH=${ARCH} GOOS=${OS} CGO_ENABLED=${CGO_ENABLED} ${GO_BUILD_CMD} -ldflags "${LDFLAGS}"\
		-o "${ROOT}/release/${NAME}" main.go

		shasum -a 256 "${ROOT}/release/${NAME}" > "${ROOT}/release/${NAME}".sha256
	done
done
