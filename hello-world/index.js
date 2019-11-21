const { spawn } = require("child_process")
// spawn("./hello-world", [], { stdio: "inherit" })
spawn("pwd && ls -a", [], { stdio: "inherit" })
spawn("cd .. && ls -a", [], { stdio: "inherit" })
