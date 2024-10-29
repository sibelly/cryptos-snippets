from web3 import Web3
w3 = Web3(Web3.HTTPProvider("https://omniscient-alpha-flower.matic.quiknode.pro/21d303fa08c32851dc178f5af7845cdd92b73b26/"))
resp = w3.provider.make_request('bb_gettx', ["FILL_ME_ARG_1"])
print(resp)