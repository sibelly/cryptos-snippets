const ethers = require("ethers");
(async () => {
  const provider = new ethers.providers.JsonRpcProvider("https://omniscient-alpha-flower.matic.quiknode.pro/21d303fa08c32851dc178f5af7845cdd92b73b26/");
  const network = await provider.send("bb_gettx", ["0x8301879a2cd385de253c5cc92aaa1b259659c8aa298c6515dfd9f5724d574643"]);
  console.log(network);
})();