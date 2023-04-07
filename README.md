go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest

cobra-cli init
cobra-cli add count

export PATH=$PATH:$(go env GOPATH)/bin# cassandra-golang-count
