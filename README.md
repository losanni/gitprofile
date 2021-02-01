# Git Profile

A git config manager

## How to use

- Save your configs to ```~/.config/git/``` with the file extension ```.gitprofile```. Example: ```work.gitprofile```

- This program will create a symlink from your ```.gitprofile``` file to the main git config in ```~/.config/git/``` (```~/.config/git/config```).

## Install

Requirements:
- [Go](https://golang.org)
- Git

```
git clone https://github.com/losanni/git-profile.git
cd git-profile
go build main.go
mv main /usr/local/bin/gitprofile
```

## Todo

- Example configs
- Cross platform binaries

