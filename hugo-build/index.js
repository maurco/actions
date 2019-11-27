const { spawn } = require("child_process")

process.chdir(__dirname)

spawn("git", ["lfs", "pull"], { stdio: "inherit" })
spawn("bin/hugo-build", { stdio: "inherit" })
