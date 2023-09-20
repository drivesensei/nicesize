Human readable file size for Go programs

### Install

```zsh
go get github.com/drivesensei/nicesize
```
### Usage

```zsh
nicesize.FormatSize(1024)  "1.00KB"

nicesize.FormatSizeWithSpacer(1024, "-")  "1.00-KB"
nicesize.FormatSizeWithSpacer(8000_000_000_000_000_000, " "))  "6.94 EB"
```
