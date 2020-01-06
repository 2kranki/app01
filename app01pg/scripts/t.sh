#!/bin/sh

cd src

echo "Running container..."
if ../dbs/postgres/run.sh ; then
    :
else
    echo "ERROR - Could not load container for postgres!"
    exit 8
fi

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


