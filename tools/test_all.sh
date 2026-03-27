#!/bin/bash

# Script to test all main.go files in the project
# Reports which files compile successfully and which have errors

set +e  # Don't exit on error

ROOT_DIR="${1:-.}"
SKIP_DIRS=("tools" ".git" ".github" "_projects")

echo "Testing all main.go files..."
echo "======================================"
echo ""

success_count=0
error_count=0
declare -a failed_dirs

# Function to check if directory should be skipped
should_skip() {
    local dir="$1"
    local basename=$(basename "$dir")
    
    for skip in "${SKIP_DIRS[@]}"; do
        if [[ "$basename" == "$skip" ]]; then
            return 0
        fi
    done
    return 1
}

# Find all main.go files
find "$ROOT_DIR" -name "main.go" -type f | while read -r main_file; do
    dir=$(dirname "$main_file")
    
    if should_skip "$dir"; then
        continue
    fi
    
    dir_name=$(basename "$dir")
    
    echo -n "Testing $dir_name... "
    
    # Try to build
    if (cd "$dir" && go build -o /tmp/test_binary main.go 2>&1 | grep -q "error"); then
        echo "❌ FAILED"
        error_count=$((error_count + 1))
        failed_dirs+=("$dir_name")
        
        # Show errors
        echo "  Errors:"
        (cd "$dir" && go build main.go 2>&1 | grep "error" | head -5 | sed 's/^/    /')
        echo ""
    else
        echo "✓ OK"
        success_count=$((success_count + 1))
        rm -f /tmp/test_binary
    fi
done

echo ""
echo "======================================"
echo "Summary:"
echo "  Success: $success_count"
echo "  Failed: $error_count"

if [ $error_count -gt 0 ]; then
    echo ""
    echo "Failed directories:"
    printf '  - %s\n' "${failed_dirs[@]}"
    exit 1
fi

echo ""
echo "✓ All tests passed!"

# Made with Bob
