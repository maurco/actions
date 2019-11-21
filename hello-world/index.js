const { spawn } = require("child_process")
spawn(`${__dirname}/target/release/hello-world`, [], { stdio: "inherit" })
