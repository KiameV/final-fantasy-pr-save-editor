.DEFAULT_GOAL := build
build:
<<<<<<< HEAD
	go build -ldflags "-s -H=windowsgui -o "FFVIPR_Save_Editor.exe"
	upx -9 -k "FFVIPR_Save_Editor.exe"
	rm "FFVIPR_Save_Editor.ex~"
	
setup:
	go install github.com/tc-hib/go-winres@latest
	go-winres simply --icon.png
=======
	go build -ldflags="-s -w -H=windowsgui" -o "FFPR_Save_Editor.exe"
	upx -9 -k "FFPR_Save_Editor.exe"
	rm "FFPR_Save_Editor.ex~"

setup:
	go install github.com/tc-hib/go-winres@latest
	go-winres simply --icon icon.png
>>>>>>> 478428277c0aa37a0885ff0a94096864d730ff78
