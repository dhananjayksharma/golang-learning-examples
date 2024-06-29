go mod init dkgosql.com/match-prefix-service && go mod tidy


go fmt internal/matchprefixes/matchprefix.go
go fmt cmd/main.go

 
mockgen --destination=../../mocks/matchprefix.go --source=matchprefix.go


