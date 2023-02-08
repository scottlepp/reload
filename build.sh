go build -buildmode=plugin -o ./plugin.so -gcflags=all='-N -l' ./pkg/plugin.go

go build -o main -gcflags=all='-N -l' ./pkg