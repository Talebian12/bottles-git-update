# Maintainer: talebian <talebian@sovietunion.xyz>
# Update package version automatically

#!/bin/bash

VERDIR="$(pwd)/Bottles"
AURDIR="$(pwd)/bottles-git"
GITDIR="$(pwd)/bottles-git-github"

check () {
  if [ ! -d $VERDIR ]
  then
	  git clone https://github.com/bottlesdevs/Bottles
  fi
  if [ ! -d $AURDIR ]
  then
	  git clone ssh://aur@aur.archlinux.org/bottles-git.git bottles-git
  fi
  if [ ! -d $GITDIR ]
  then
	  git clone https://github.com/talebian12/bottles-git bottles-git-github
  fi
}

getver () {
  check
  cd "$VERDIR"
  git pull origin master
  VER="$(git log --format="%h" -n 1)"
  cat $VER
}

updatepkg() {
  getver
  cd "$AURDIR"
  sed -i "s/pkgver=2.*/pkgver=2.$VER/g" PKGBUILD
  sed -i "s/^_ver=.*/_ver=$VER/g" PKGBUILD
  cat PKGBUILD
  sed -i "s/pkgver = 2.*/pkgver = 2.$VER/g" .SRCINFO
  cat .SRCINFO

  cp PKGBUILD $GITDIR
  cp .SRCINFO $GITDIR

  git add .
  git commit -m "Automatic Update to 2.$VER"
  git push

  cd "$GITDIR"
  git add .
  git commit -m "Automatic Update to 2.$VER"
  git push origin master
}

updatepkg
