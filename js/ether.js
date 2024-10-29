const ethers = require('ethers');
const ABI = require('./abi.json'); // Contract ABI
const provider = ethers.getDefaultProvider();

const inter = new ethers.utils.Interface(ABI);

(async() => {
  const tx = await provider.getTransaction('0xa30e9e19967bd3307feeddcf99b26be0d804cdc0ade6929f3b9328a95e388b4c');
    const decodedInput = inter.parseTransaction({ data: tx.data, value: tx.value});

    // Decoded Transaction
    console.log({
        function_name: decodedInput.name,
        from: tx.from,
        to: decodedInput.args[0],
        erc20Value: Number(decodedInput.args[1])
      });        
})();