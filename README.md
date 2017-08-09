# git-watch

## Example watch file

```
version=4
opts="filenamemangle=s%(?:.*?)?([^/]*)\.tar\.gz%<project>-$1.tar.gz%" \
    http://<server>/github/<owner>/<repo> \
    (?:.*?/)?([^/]*)\.tar\.gz
```
