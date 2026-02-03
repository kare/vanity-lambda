#!/bin/bash
set -euo pipefail

REMOTE="${REMOTE:-origin}"
NAME=$(basename "$0")

usage() {
	echo "Usage: $NAME <tag> [message]" >&2
	echo "" >&2
	echo "Examples:" >&2
	echo "	$NAME v1.2.3" >&2
}

if [ "${1:-}" = "" ]; then
	usage
	exit 1
fi

TAG="$1"
MSG="${2:-"Release ${TAG}"}"

if ! git diff --quiet || ! git diff --cached --quiet; then
	echo "$NAME: error: Working tree is not clean." >&2
	exit 1
fi

if git rev-parse -q --verify "refs/tags/${TAG}" >/dev/null; then
	echo "$NAME: error: Tag '${TAG}' exists." >&2
	exit 1
fi

CURRENT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
if [ "${CURRENT_BRANCH}" != "main" ]; then
	echo "$NAME: error: Wrong branch '${CURRENT_BRANCH}'. Release must be made from 'main' branch." >&2
	exit 1
fi

git commit --allow-empty -m "release: github.com/kare/vanity-lambda ${TAG}"
git tag -a "${TAG}" -m "${MSG}"
git push "${REMOTE}" "${CURRENT_BRANCH}"
git push "${REMOTE}" "${TAG}"

