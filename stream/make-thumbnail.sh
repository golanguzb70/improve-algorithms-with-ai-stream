#!/usr/bin/env bash
# Render the universal stream thumbnail (1280×720) → stream/thumbnail.png
set -euo pipefail

DIR="$(cd "$(dirname "$0")" && pwd)"
CHROME="/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"

"$CHROME" --headless=new --disable-gpu --hide-scrollbars \
  --window-size=1280,720 --screenshot="$DIR/thumbnail.png" \
  "file://$DIR/thumbnail.html" >/dev/null 2>&1

echo "$DIR/thumbnail.png"
