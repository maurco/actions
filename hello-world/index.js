const fs = require("fs")
// const { spawn } = require("child_process")
// spawn("./hello-world", [], { stdio: "inherit" })
// spawn("pwd && ls -a", [], { stdio: "inherit" })
// spawn("cd .. && ls -a", [], { stdio: "inherit" })

console.log(process.cwd())
console.log(__dirname)
console.log(fs.readdirSync(__dirname))
