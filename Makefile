export GOPATH=/home/rob/go

libzlint.so:
	$(GOPATH)/bin/plgo .

clean:
	rm -rf build
