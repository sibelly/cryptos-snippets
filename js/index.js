var Web3 = require('web3');

utils = Web3.utils
let blah = utils.hexToNumber('0x668ee6ed')
console.log(blah)

let blah2 = utils.hexToAscii('0x0000000000000000000000001e213600fa9317feac4ef4087acdf5d0e25d7187')
console.log(blah2)

// let bah3 = utils.hex

// Web3.eth.abi.decodeParameter()