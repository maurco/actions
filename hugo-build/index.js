const fs = require("fs")
const { spawn } = require("child_process")

console.log(fs.readdirSync(__dirname))
console.log(fs.readdirSync(`${__dirname}/bin`))
spawn(`${__dirname}/bin/hugo-build`, { stdio: "inherit" })
