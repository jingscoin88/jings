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
	"enode://023297477134b7b9f433578ad4f6cd866c37b94323cb8b6b07ada310facc5157d7ba6b750ecc1b4239492202d9400b0462eb73c2e10f4ad4a774d41a6f32a345@52.221.193.178:30303",
	"enode://c38f8d2283633ae10384a64c3269e71eda03216bdcc65747069333b05743c4c73584a439d7fbe4cf7e3dd6de25922f5977a95b11fe92059062fd5c8e548d587c@18.140.235.187:30303",
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
