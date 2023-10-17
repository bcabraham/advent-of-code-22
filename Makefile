# Usage: make new dir=dayN
dir="newDir"

help:
	echo "usage: make new dir=dayN"

new:
	mkdir $(dir)
	touch ./$(dir)/README.md
	echo "package $(dir)" >> ./$(dir)/$(dir).go
	echo "package $(dir)_test" >> ./$(dir)/$(dir)_test.go
	