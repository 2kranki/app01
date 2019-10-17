
#!/bin/sh

# vi:nu:et:sts=4 ts=4 sw=4

# This script builds the application, App01ms.

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

if [ -d "$prefix/src/github.com/denisenkom/go-mssqldb" ]; then   # MS SQL Driver
    :
else
    echo "...Fetching github.com/denisenkom/go-mssqldb"
    go get github.com/denisenkom/go-mssqldb
fi



if [ -d "$prefix/src/github.com/shopspring/decimal" ]; then      # Decimal Number Support
    :
else
    echo "...Fetching github.com/shopspring/decimal"
    go get github.com/shopspring/decimal
fi

echo "Compiling the program..."
cd cmd/app01ms
echo "...Formatting files:"
go fmt ./...
echo "...Building Application:"
mkdir -p /tmp/bin
go build -o /tmp/bin/App01ms -v
if [ $? -eq 0 ] ; then
    echo "Built: /tmp/bin/App01ms"
fi
if [ -n "$GOLANG_DOCKER_CONTAINER" ]; then
    cp /tmp/bin/App01ms /go/bin/
fi
cd -

