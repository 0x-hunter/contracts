#!/usr/bin/env node

const path = require('path')
const fs = require('fs')

let inFile = process.argv[2]
let outDir = process.argv[3]

let jsonObject = JSON.parse(fs.readFileSync(inFile).toString())
let basename = path.basename(inFile)

fs.writeFileSync(path.join(outDir, basename + ".ts"), `export default ${JSON.stringify({
    abi : jsonObject.abi
}, undefined, 2)};`)
