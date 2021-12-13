#!/bin/bash

cp -r template $1
cd $1
go mod init github.com/aoc20i21/$1
