# Usage make dir=dayN
dir="newDir"
new:
	mkdir $(dir)
	touch ./$(dir)/README.md
	echo "package $(dir)" >> ./$(dir)/$(dir).go
	echo "package $(dir)_test" >> ./$(dir)/$(dir)_test.go
	