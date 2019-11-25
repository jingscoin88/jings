// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Jingscoin Go Bootnodes
	"enode://a5ed591e8842f7d1788934bcb806ee5a9fc67b124c7eaa2cba77aae515a8a3fb18192e95e8bc3786040e3ecb45ea8761c41be9000d089b974f406d1dc6d4c151@3.0.185.98:30303",
	"enode://1406e33d1becaccaa5755cfc58f89d84c629e7eaef0bc3442f1cfab906348fe002ad85a84bac6232ddf9db26da8041100bd763ccbe977265338d40b4550ebbe6@18.140.235.187:30303",
	"enode://ab831ed8c7b2e8d2a049061cb0be18c7b5747fbf55d74b3c60cb828f77f977dd7fbb04bea3a8187a7a3dd26fd9f1a72784fb80f6a8971c1ceb8e366459bc3cfa@18.140.235.187:30303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{

}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	
}
