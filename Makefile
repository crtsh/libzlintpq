export GOPATH=/home/rob/go

libzlintpq.so:
	$(GOPATH)/bin/plgo -v .

clean:
	rm -rf build
