#!/bin/bash
# chmod +x format.sh

# Fromatting all .go files in project directory
echo "Formatting .go files in directory and subdirectiories"

$HOME/go/bin/goimports -w .

echo "Formatting finished."
