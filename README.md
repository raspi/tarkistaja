![GitHub All Releases](https://img.shields.io/github/downloads/raspi/tarkistaja/total?style=for-the-badge)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/raspi/tarkistaja?style=for-the-badge)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/raspi/tarkistaja?style=for-the-badge)

# tarkistaja
List file checksums inside of compressed archive

# Features
* Archives
  * Everything that [archiver](https://github.com/mholt/archiver) supports
* Checksum methods
  * sha1
  * sha256
  * sha512
  * md5

# Examples

    % tarkistaja -m md5 test.zip 
    d41d8cd98f00b204e9800998ecf8427e empty.txt
    746308829575e17c3331bbcb00c0898b hello.txt

Add archive file name as a directory as additional information:

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
tarkistaja v1.0.0 - (2020-07-07T19:38:16+03:00)
(c) Pekka Järvinen 2020- [ https://github.com/raspi/tarkistaja ]
List file checksums inside of compressed archive.

  Usage:
    tarkistaja [parameters] <compressed file>

Parameters:
  -a    Add archive's file name as a directory name (as additional information)
  -m string
        Checksum method (sha1, sha256, sha512, md5) (default "sha256")
  -o string
        Output checksums to file <filename> instead of STDOUT

Examples:
  List checksums:
    tarkistaja important_files.zip
  List checksums to file:
    tarkistaja -o checksums.sha256 important_files.zip
  List checksums using md5:
    tarkistaja -m md5 important_files.zip
  Add archive file name as directory for additional information:
    tarkistaja -a important_files.zip
```

# Requirements
* Operating system
  * GNU/Linux 
    * x64 arm arm64 ppc64 ppc64le
  * Microsoft Windows
    * x64
  * Darwin (Apple Mac)
    * x64
  * FreeBSD
    * x64 arm
  * NetBSD
    * x64 arm
  * OpenBSD
    * x64 arm arm64

# Get source

    git clone https://github.com/raspi/tarkistaja
    
# License

MIT
    