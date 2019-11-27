const { spawn } = require("child_process")

process.chdir(`${__dirname}/../`)

spawn("git", ["lfs", "pull"], { stdio: "inherit" })
spawn("hugo-build/bin/hugo-build", { stdio: "inherit" })
