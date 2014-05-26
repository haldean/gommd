MultiMarkdown in Go.
==

This wraps MMD-4 in Go. You need to clone my [patched version of
MultiMarkdown][mmd] that allows you to build a shared library version of MMD,
then run:

    cd MultiMarkdown-4
    make lib
    sudo make install-lib

Which will build and install the shared library on your system. You can then use
gommd just like you normally would. Patches to add features of MultiMarkdown in
a more Go-like way are welcome!

[mmd]: https://github.com/haldean/MultiMarkdown-4
