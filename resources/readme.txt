We use https://github.com/jteeuwen/go-bindata to convert the resources directory and its contents
into compile-able go code, so the files can compiled into an executable as addressable resources
or assets.

Build instructions
------------------

    cd <the resources directory>
    go-ws/bin/go-bindata -debug staticfiles/...

This produces the bindata.go file - which declares itself to be in package main.

You can then use it like this:

    data, err := Asset("pub/style/foo.css")

Or wrap a virtual file system around it using:

    https://github.com/elazarl/go-bindata-assetfs

Important Note
--------------

The instructions above make a debug version that CHEATS and DOES read the real files at runtime
which is obviously more convenient during development, but COMPLETELY defeats the purpose of
doing it for deployment.