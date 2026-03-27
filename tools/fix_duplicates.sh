#!/bin/bash

# Script to find and report duplicate function names in consolidated main.go files

echo "Checking for duplicate function names..."
echo "========================================"
echo ""

find . -name "main.go" -type f ! -path "./_projects/*" ! -path "./tools/*" | while read -r file; do
    dir=$(dirname "$file")
    dir_name=$(basename "$dir")
    
    # Find duplicate function declarations
    duplicates=$(grep -n "^func " "$file" | awk '{print $2}' | cut -d'(' -f1 | sort | uniq -d)
    
    if [ -n "$duplicates" ]; then
        echo "❌ $dir_name has duplicate functions:"
        echo "$duplicates" | while read -r func; do
            echo "  - $func"
            grep -n "^func $func" "$file" | head -5
        done
        echo ""
    fi
done

echo "========================================"
echo "Check complete!"

# Made with Bob
