// Import the API, Keyring and some utility functions
const { ApiPromise, WsProvider } = require('@polkadot/api');
const { Keyring } = require('@polkadot/keyring');

const ADDR1 = '';
const ADDR2 = '';

async function main () {
  const provider = new WsProvider('');

  // Instantiate the API
  const api = await ApiPromise.create({provider});

  // Construct the keyring after the API (crypto has an async init)
  const keyring = new Keyring({ type: 'sr25519' });

  // Add ADDR1 to our keyring with a hard-derivation path 
  const ADDR1 = keyring.addFromUri('');

  // // Create a extrinsic, transferring 12345 units to ADDR2
  // const transfer = api.tx.balances.transferAllowDeath(ADDR2, 1);


  // // Sign and send the transaction using our account
  // const hash = await transfer.signAndSend(ADDR1);

  // console.log('Transfer sent with hash', hash.toHex());

  const txHash = await api.tx.balances
  .transferKeepAlive(ADDR2, 1000000)
  .signAndSend(ADDR1);

// Show the hash
console.log(`Submitted with hash ${txHash}`);

}

main().catch(console.error).finally(() => process.exit());