rm -f ./output/checRepeat
go mod tidy
go build  -o ./output/checkRepeat ./main/main.go
