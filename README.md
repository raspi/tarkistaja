# tarkistaja
List file checksums inside of compressed archive

# Examples

    % tarkistaja -m md5 test.zip 
    d41d8cd98f00b204e9800998ecf8427e empty.txt
    746308829575e17c3331bbcb00c0898b hello.txt

Add archive file name as a directory:

    % tarkistaja -m md5 -a test.zip
    d41d8cd98f00b204e9800998ecf8427e test.zip/empty.txt
    746308829575e17c3331bbcb00c0898b test.zip/hello.txt

Write checksums to a file:

    % tarkistaja -o checksums.sha256 -a test.zip
    2020/07/07 19:31:37 hashing "empty.txt"
    2020/07/07 19:31:37 hashing "hello.txt"
    2020/07/07 19:31:37 Done.
    % cat checksums.sha256
    e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855 test.zip/empty.txt
    d9014c4624844aa5bac314773d6b689ad467fa4e1d1a50a1b8a99d5a95f72ff5 test.zip/hello.txt
    
    
# Usage
```
```

# Get source

    git clone https://github.com/raspi/tarkistaja
    
# License

MIT
    