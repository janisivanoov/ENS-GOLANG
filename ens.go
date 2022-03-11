package ens

import (
	ethereum "github.com/ethereum/go-ethereum"
)

func ens() {
	// Can dial up a connection through either IPC or HTTP/HTTPS
	client, _ := ethereum.Dial("/home/ethereum/.ethereum/geth.ipc")
	registry, _ = ens.Registry(client)

	// Encoding
	bin, _ := ens.StringToContenthash("/ipfs/QmayQq2DWCkY3d4x3xKh4suohuRPEXe2fBqMBam5xtDj3t")
	// Setting contenthash
	resolver.SetContenthash(opts, data)
	// Getting contenthash
	resolver.Contenthash()
	// Decoding
	_, _ = ens.ContenthashToString(bin)

	// Getting Multicoin
	_, _ = resolver.MultiAddress(0)
	// Setting Multicoin
	resolver.SetMultiAddress(opts, address)

	// Setting text
	resolver.SetText(opts, name, value)
	// Getting text
	resolver.Text(name)
}
