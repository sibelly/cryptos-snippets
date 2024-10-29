// Required imports
const { ApiPromise, WsProvider } = require('@polkadot/api');

const ADDR1 = '';

async function main () {
  // Initialise the provider to connect to the local node
  const provider = new WsProvider('');

  // Create the API and wait until ready
  const api = await ApiPromise.create({ provider });

  // Retrieve the chain & node information via rpc calls
  const [chain, nodeName, nodeVersion] = await Promise.all([
    api.rpc.system.chain(),
    api.rpc.system.name(),
    api.rpc.system.version()
  ]);

  console.log(`### You are connected to chain ${chain} using ${nodeName} v${nodeVersion}`);

  console.log(`### Retrieve last block timestamp, account nonce & balances`)
  const [now, { nonce, data: balance }] = await Promise.all([
    api.query.timestamp.now(),
    api.query.system.account(ADDR1)
  ]);
  console.log(`${now} - ${nonce} - ${balance}`);

  console.log(`### Retrieve chain properties`)
  const chainInfo = api.registry.getChainProperties()
  console.log("tokenDecimals -> ", chainInfo.tokenDecimals.toHuman());
  console.log("tokenSymbol -> ", chainInfo.tokenSymbol.toHuman());

}

main().catch(console.error).finally(() => process.exit());