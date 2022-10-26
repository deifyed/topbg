# TopBG

## Motivation

TopBG grabs a random image from the top posts of configured subreddits and sets it as the desktop wallpaper

## Usage

```shell
# Set a new random wallpaper
topbg set

# Set wallpaper with image by index
topbg set --index <index> # Find index with `topbg list`

# Permanently save the image previously set
topbg store

# List stored images with indexes
topbg list
```

## Installation

Requires Go 1.19

```shell
# Install topbg into ~/.local/bin
make build && make install

# To adjust installation dir, use INSTALL_DIR
make build && make install INSTALL_DIR=/usr/local/bin
```
