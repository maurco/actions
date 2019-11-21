const { spawn } = require("child_process")
spawn("./hello-world", [], { stdio: "inherit" })
