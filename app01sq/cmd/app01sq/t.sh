#!/bin/sh



echo "Testing the package:"
go test -v ./...

echo "Removing created test data if needed:"
files=(*.db)
for file in "${files[*]}"
do
    if test -f "$file"; then
        echo "...Deleting ${file}"
        rm $file
    fi
done

