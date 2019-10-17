
#!/bin/sh

# vi:nu:et:sts=4 ts=4 sw=4

# This script builds the application, App01sq.

prefix="$HOME/go"


if [ -n "$GOLANG_DOCKER_CONTAINER" ]; then
    prefix=/go
fi

echo "Adding needed packages..."
if [ -d "$prefix/src/github.com/2kranki/jsonpreprocess" ]; then  # JSON Comment Remover
    :
else
    echo "...Fetching github.com/2kranki/jsonpreprocess"
    go get github.com/2kranki/jsonpreprocess
fi
if [ -d "$prefix/src/github.com/2kranki/go_util" ]; then         # Utility Functions
    :
else
    echo "...Fetching github.com/2kranki/go_util"
    go get github.com/2kranki/go_util
fi




if [ -d "$prefix/go/src/github.com/mattn/go-sqlite3" ]; then        # SQLite 3 Driver
    :
else
    echo "...Fetching github.com/mattn/go-sqlite3"
    go get github.com/mattn/go-sqlite3
fi
if [ -d "$prefix/src/github.com/shopspring/decimal" ]; then      # Decimal Number Support
    :
else
    echo "...Fetching github.com/shopspring/decimal"
    go get github.com/shopspring/decimal
fi

echo "Compiling the program..."
cd cmd/app01sq
echo "...Formatting files:"
go fmt ./...
echo "...Building Application:"
mkdir -p /tmp/bin
go build -o /tmp/bin/App01sq -v
if [ $? -eq 0 ] ; then
    echo "Built: /tmp/bin/App01sq"
fi
if [ -n "$GOLANG_DOCKER_CONTAINER" ]; then
    cp /tmp/bin/App01sq /go/bin/
fi
cd -

