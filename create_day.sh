#!/usr/bin/env bash

set -e

day_idx=$1

if [[ -z $AOC_SESSION ]]; then
	echo "please set AOC_SESSION variable to your aoc sesssion cookie value"
	exit 1
fi

mkdir "./day$day_idx"
file_path="./day$day_idx/day$day_idx.go"

cat << END > "$file_path"
package day$day_idx

import "github.com/A-Siam/aoc-go/utils"

func Solution(inputPath string) {
	utils.GetInput(inputPath, func(line string) {

    })
}
END

input_file="./input/day$day_idx"

curl "https://adventofcode.com/2022/day/$day_idx/input" \
	-H "cookie:session=$AOC_SESSION" > "$input_file"