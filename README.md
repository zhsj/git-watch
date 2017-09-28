# git-watch

## Example watch file

```
version=4
opts="filenamemangle=s%(?:.*?)?@ANY_VERSION@\.tar\.gz%@PACKAGE@-$1.tar.gz%" \
    https://watch.zhsj.me/github/<owner>/<repo> \
    (?:.*?/)?([^/]*)\.tar\.gz
```
