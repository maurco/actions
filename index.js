const { spawn } = require("child_process")

const options = {
	stdio: "inherit",
}

spawn(
	process.argv[2],
	process.argv.slice(3),
	options,
)
