#!/bin/bash

# Script to consolidate multiple Go files in each directory into a single main.go
# Each original file's main() function becomes a separate example function

set -e

ROOT_DIR="${1:-.}"
SKIP_DIRS=("tools" ".git" ".github" "_projects" "node_modules")

echo "Starting consolidation of Go project..."
echo "Root directory: $ROOT_DIR"
echo ""

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

# Function to consolidate a directory
consolidate_dir() {
    local dir="$1"
    local go_files=($(ls "$dir"/*.go 2>/dev/null | grep -v "/main.go$" || true))
    
    if [[ ${#go_files[@]} -le 1 ]]; then
        return
    fi
    
    echo "Consolidating $(basename "$dir") (${#go_files[@]} files)..."
    
    local main_file="$dir/main.go"
    local temp_file="$dir/main.go.tmp"
    
    # Start building main.go
    echo "package main" > "$temp_file"
    echo "" >> "$temp_file"
    
    # Collect all imports
    echo "import (" >> "$temp_file"
    for file in "${go_files[@]}"; do
        grep "^import" "$file" | sed 's/import//' | tr -d '()' | grep -v "^$" | sort -u >> "$temp_file.imports" 2>/dev/null || true
    done
    
    if [[ -f "$temp_file.imports" ]]; then
        sort -u "$temp_file.imports" | sed 's/^/\t/' >> "$temp_file"
        rm "$temp_file.imports"
    fi
    echo ")" >> "$temp_file"
    echo "" >> "$temp_file"
    
    # Process each file
    local example_funcs=()
    local counter=1
    
    for file in "${go_files[@]}"; do
        local filename=$(basename "$file" .go)
        local safe_name=$(echo "$filename" | tr '-' '_' | tr ' ' '_' | tr '[:upper:]' '[:lower:]')
        local func_name="example${counter}_${safe_name}"
        example_funcs+=("$func_name")
        
        echo "// ============================================================================" >> "$temp_file"
        echo "// Example $counter: $(basename "$file")" >> "$temp_file"
        echo "// ============================================================================" >> "$temp_file"
        echo "" >> "$temp_file"
        
        # Extract everything except package and import statements
        awk '
            BEGIN { in_import=0; in_main=0; brace_count=0 }
            /^package / { next }
            /^import \(/ { in_import=1; next }
            in_import==1 && /^\)/ { in_import=0; next }
            in_import==1 { next }
            /^import / { next }
            /^func main\(\)/ { 
                print "func '"$func_name"'() {"
                in_main=1
                brace_count=1
                next
            }
            in_main==1 {
                if ($0 ~ /{/) brace_count++
                if ($0 ~ /}/) brace_count--
                if (brace_count == 0) {
                    in_main=0
                    print "}"
                    print ""
                    next
                }
                print
                next
            }
            { print }
        ' "$file" >> "$temp_file"
        
        echo "" >> "$temp_file"
        counter=$((counter + 1))
    done
    
    # Create main function
    echo "// ============================================================================" >> "$temp_file"
    echo "// Main Function - Run All Examples" >> "$temp_file"
    echo "// ============================================================================" >> "$temp_file"
    echo "" >> "$temp_file"
    echo "func main() {" >> "$temp_file"
    echo "	fmt.Println(\"Running ${#example_funcs[@]} examples from $(basename "$dir")\")" >> "$temp_file"
    echo "	fmt.Println(strings.Repeat(\"=\", 60))" >> "$temp_file"
    echo "" >> "$temp_file"
    
    for func in "${example_funcs[@]}"; do
        echo "	fmt.Println(\"\\n--- $func ---\")" >> "$temp_file"
        echo "	$func()" >> "$temp_file"
        echo "" >> "$temp_file"
    done
    
    echo "	fmt.Println(\"\\n\" + strings.Repeat(\"=\", 60))" >> "$temp_file"
    echo "	fmt.Println(\"All examples completed!\")" >> "$temp_file"
    echo "}" >> "$temp_file"
    
    # Move temp file to main.go
    mv "$temp_file" "$main_file"
    
    # Delete old files
    for file in "${go_files[@]}"; do
        rm "$file"
        echo "  ✓ Removed $(basename "$file")"
    done
    
    echo "  ✓ Created main.go"
    echo ""
}

# Find all directories with Go files
find "$ROOT_DIR" -type d -not -path "*/\.*" | while read -r dir; do
    if should_skip "$dir"; then
        continue
    fi
    
    consolidate_dir "$dir"
done

echo "✓ Consolidation complete!"

# Made with Bob
