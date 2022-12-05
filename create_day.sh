#!/usr/bin/env bash

day_idx=$1

mkdir "./day$day_idx"
file_path="./day$day_idx/day$day_idx.go"

cat << END > $file_path
package day$day_idx

import "github.com/A-Siam/aoc-go/utils"

func Solution(inputPath string) {
	utils.GetInput(inputPath, func(line string) {

    })
}
END
