# libzlintpq
Run zlint from a PostgreSQL function.


```
GOPATH=/root/go go get -u github.com/microo8/plgo/...
GOPATH=/root/go go get -u github.com/zmap/zlint
make
cd build
make install
```

```
$ psql
CREATE EXTENSION libzlintpq;
\i zlint_embedded.fnc
```
