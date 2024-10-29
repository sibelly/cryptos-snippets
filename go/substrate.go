package main

import (
	"bytes"
	"fmt"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

func main() {
	err := SendData()

	if err != nil {
		fmt.Println("error:", err)

	}

}

func SendData() error {
	//kos private key
	memo := ""

	api, err := gsrpc.NewSubstrateAPI("")
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

	amount := types.NewUCompactFromUInt(1)

	bob, err := types.NewMultiAddressFromHexAccountID("")
	if err != nil {
		return err
	}

	c, err := types.NewCall(meta, "Balances.transfer_keep_alive", bob, amount)
	if err != nil {
		fmt.Println("err call", err)
		return err
	}
	// Create the extrinsic
	ext := types.NewExtrinsic(c)
	fmt.Println("extrinsic version:", ext.Version)
	fmt.Println("extrinsic:", ext)

	// Assuming `extrinsicBytes` is the byte slice of your encoded extrinsic
	var seilah types.Extrinsic
	extrinsicBytes := []byte("0x4d0284006a298320c2b3422006b7896b2f102ff7fad99f26a0d9807642a293667cdac13e0100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000140300dc0c382e18e8765f85170c5f3fb9158b9628c71442d3394f512845ac50a341441700b813d4d5a39e2702")
	decoder := scale.NewDecoder(bytes.NewReader(extrinsicBytes))
	seilah.Decode(*decoder)

	fmt.Println("seilah ", seilah)

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
	fmt.Println("done 3")
	key, err := types.CreateStorageKey(meta, "System", "Account", keypair.PublicKey)
	if err != nil {
		return err
	}
	fmt.Println("done 4 key ", key)
	var accountInfo types.AccountInfo
	ok, err := api.RPC.State.GetStorageLatest(key, &accountInfo)
	fmt.Println("done 5", ok, err, accountInfo)
	if err != nil || !ok {
		return err
	}

	nonce := uint32(accountInfo.Nonce)
	o := types.SignatureOptions{
		BlockHash:          genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
		SpecVersion:        rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(100),
		TransactionVersion: rv.TransactionVersion,
	}
	fmt.Println("before sign:", o)
	// Sign the transaction using Alice's default account
	err = ext.Sign(keypair, o)
	if err != nil {
		return err
	}
	fmt.Println("before submission")
	// Send the extrinsic
	txnHash, err := api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return err
	}
	fmt.Println("after submission txnHash", txnHash)
	return nil
}
