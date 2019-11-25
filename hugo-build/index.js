const { spawn } = require("child_process")
spawn(`${__dirname}/bin/hugo-build`, { stdio: "inherit" })
