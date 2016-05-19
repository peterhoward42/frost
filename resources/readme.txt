We use https://github.com/jteeuwen/go-bindata to convert the resources directory and its contents
into compile-able go code, so the files can compiled into an executable as addressable resources
or assets.

Build instructions
------------------

    These are scripted in make.bat

To access resources programatically
-------------------------------------
    data, err := Asset("pub/style/foo.css")

    Or wrap a virtual file system around it using:

    https://github.com/elazarl/go-bindata-assetfs
