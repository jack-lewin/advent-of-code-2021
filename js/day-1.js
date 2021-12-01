const fs = require("fs")

const input = fs
    .readFileSync("inputs/day-1.txt", "utf8")
    .split("\n")
    .map(_ => parseInt(_, 10))

const isGreaterThanPrevious = (acc, cur, i, arr) => {
    if (i === 0) return acc
    else if (cur > arr[i - 1]) return acc + 1
    else return acc
}

const getRollingAverage = (_, i, arr) => {
    if (!arr[i + 2]) return null
    else return arr[i] + arr[i + 1] + arr[i + 2]
}

const part1 = input
    .reduce(isGreaterThanPrevious, 0)

const part2 = input
    .flatMap(getRollingAverage)
    .reduce(isGreaterThanPrevious, 0)

console.log("part 1:", part1)
console.log("part 2:", part2)
