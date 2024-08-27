#!/bin/bash

set -euo pipefail

function show_usage() {
    echo "Usage: $0 -v BUILD_VERSION -o TARGET_OS -a TARGET_ARCH -n IMAGE_NAME -t IMAGE_TAG"
    echo "  -o TARGET_OS       Target operating system (e.g., windows, linux)" echo "  -a TARGET_ARCH     Target architecture (e.g., amd64, arm64)"
    echo "  -n IMAGE_NAME      Name of the Docker image"
    echo "  -t IMAGE_TAG       Tag for the Docker image (e.g., latest, stable)"
    exit 1
}

TARGET_OS=""
TARGET_ARCH=""
IMAGE_NAME=""
IMAGE_TAG=""

while getopts "v:o:a:n:t:" option; do
    case "${option}" in
        o) TARGET_OS="${OPTARG}" ;;
        a) TARGET_ARCH="${OPTARG}" ;;
        n) IMAGE_NAME="${OPTARG}" ;;
        t) IMAGE_TAG="${OPTARG}" ;;
        *) show_usage ;;
    esac
done

if [[ -z "${TARGET_OS}" ]]; then
    echo "Error: TARGET_OS is required."
    show_usage
fi

if [[ -z "${TARGET_ARCH}" ]]; then
    echo "Error: TARGET_ARCH is required."
    show_usage
fi

if [[ -z "${IMAGE_NAME}" ]]; then
    echo "Error: IMAGE_NAME is required."
    show_usage
fi

if [[ -z "${IMAGE_TAG}" ]]; then
    echo "Error: IMAGE_TAG is required."
    show_usage
fi

docker buildx build \
    --progress=plain \
    --build-arg "TARGET_OS=${TARGET_OS}" \
    --build-arg "TARGET_ARCH=${TARGET_ARCH}" \
    -f .docker/Dockerfile . \
    -t "${IMAGE_NAME}:${IMAGE_TAG}"

echo "Docker image '${IMAGE_NAME}:${IMAGE_TAG}' built successfully for ${TARGET_OS}/${TARGET_ARCH}."

