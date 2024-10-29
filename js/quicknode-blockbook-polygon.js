const ethers = require("ethers");
(async () => {
  const provider = new ethers.providers.JsonRpcProvider("");
  const network = await provider.send("bb_gettx", ["0x8301879a2cd385de253c5cc92aaa1b259659c8aa298c6515dfd9f5724d574643"]);
  console.log(network);
})();