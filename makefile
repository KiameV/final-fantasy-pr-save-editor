.DEFAULT_GOAL := build
build:
	go build -ldflags -H=windowsgui -o pr_save_editor.exe
	#upx -9 -k pr_save_editor.exe
	#rm pr_save_editor.ex~

