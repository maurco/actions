const { spawn } = require("child_process")
spawn("../bin/hello-world", [], { stdio: "inherit" })
