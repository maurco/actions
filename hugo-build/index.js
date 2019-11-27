const { spawn } = require("child_process")

process.chdir(`${__dirname}/../`)
console.log(process.cwd())
spawn("ls -a", { stdio: "inherit" })

process.chdir("../")
console.log(process.cwd())
spawn("ls -a", { stdio: "inherit" })

spawn("git", ["lfs", "pull"], { stdio: "inherit" })
spawn("hugo-build/bin/hugo-build", { stdio: "inherit" })
