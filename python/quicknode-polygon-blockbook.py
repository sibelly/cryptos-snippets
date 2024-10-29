from web3 import Web3
w3 = Web3(Web3.HTTPProvider(""))
resp = w3.provider.make_request('bb_gettx', ["FILL_ME_ARG_1"])
print(resp)