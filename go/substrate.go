package main

import (
	"fmt"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/extrinsic"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/extrinsic/extensions"
)

func main() {
	err := SendData()
	if err != nil {
		fmt.Println("error:", err)
	}
}

func SendData() error {
	//mnemonic
	memo := ""

	api, err := gsrpc.NewSubstrateAPI("WSS ADDRESS HERE")
	if err != nil {
		return err
	}
	fmt.Println("connected")

	keypair, err := signature.KeyringPairFromSecret(memo, 36)
	fmt.Println("keypair", keypair.Address, keypair.PublicKey, keypair.URI)
	if err != nil {
		fmt.Println("keypair", err)
	}
	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return err
	}

	amount := types.NewUCompactFromUInt(1000000)

	ADDR2, err := types.NewMultiAddressFromHexAccountID("OTHER ADDRESS")
	if err != nil {
		return err
	}

	c, err := types.NewCall(meta, "Balances.transfer_keep_alive", ADDR2, amount)
	if err != nil {
		fmt.Println("err call", err)
		return err
	}
	// Create the extrinsic
	ext := extrinsic.NewExtrinsic(c)
	fmt.Println("extrinsic version:", ext.Version)

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		fmt.Println("genesisHash, err", err)
		return err
	}
	fmt.Println("done 2 genesisHash", genesisHash)
	rv, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		fmt.Println("GetRuntimeVersionLatest, err", err)
		return err
	}
	fmt.Println("Runtime version => ", rv)
	key, err := types.CreateStorageKey(meta, "System", "Account", keypair.PublicKey)
	if err != nil {
		return err
	}
	fmt.Println("Key ", key)
	var accountInfo types.AccountInfo
	ok, err := api.RPC.State.GetStorageLatest(key, &accountInfo)
	fmt.Println("accountInfo => ", ok, err, accountInfo)
	if err != nil || !ok {
		return err
	}

	// Sign the transaction using ADDR1's default account
	err = ext.Sign(
		keypair,
		meta,
		extrinsic.WithEra(types.ExtrinsicEra{IsImmortalEra: true}, genesisHash),
		extrinsic.WithNonce(types.NewUCompactFromUInt(uint64(accountInfo.Nonce))),
		extrinsic.WithTip(types.NewUCompactFromUInt(0)),
		extrinsic.WithSpecVersion(rv.SpecVersion),
		extrinsic.WithTransactionVersion(rv.TransactionVersion),
		extrinsic.WithGenesisHash(genesisHash),
		extrinsic.WithMetadataMode(extensions.CheckMetadataModeDisabled, extensions.CheckMetadataHash{Hash: types.NewEmptyOption[types.H256]()}),
		extrinsic.WithAssetID(types.NewEmptyOption[types.AssetID]()),
	)

	if err != nil {
		return err
	}
	fmt.Println("before submission")
	encodedExt, err := codec.EncodeToHex(ext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ext - %s\n", encodedExt)

	sub, err := api.RPC.Author.SubmitExtrinsic(ext)
	if err != nil {
		return err
	}
	fmt.Println("after submission txnHash", sub.Hex())
	return nil
}
