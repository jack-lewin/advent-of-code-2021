const fs = require("fs")

const input = fs
    .readFileSync("inputs/day-2.txt", "utf8")
    .split("\n")
    .map(line => {
        const [direction, amount] = line.split(" ")
        return [direction, parseInt(amount, 10)]
    })

const startAt = [0, 0, 0]

const movePart1 = ([horizontal, depth], [direction, amount]) => {
    if (direction === "forward") return [horizontal + amount, depth]
    else if (direction === "down") return [horizontal, depth + amount]
    else if (direction === "up") return [horizontal, depth - amount]
    else return [horizontal, depth]
}

const movePart2 = ([horizontal, depth, aim], [direction, amount]) => {
    if (direction === "forward") return [horizontal + amount, depth + (aim * amount), aim]
    else if (direction === "down") return [horizontal, depth, aim + amount]
    else if (direction === "up") return [horizontal, depth, aim - amount]
    else return [horizontal, depth, aim]
}

const multiplyList = (a, b) => a * b

const part1 = input
    .reduce(movePart1, startAt)
    .reduce(multiplyList, 1)

const part2 = input
    .reduce(movePart2, startAt)
    .slice(0, 2)
    .reduce(multiplyList, 1)

console.log("part 1:", part1)
console.log("part 2:", part2)
