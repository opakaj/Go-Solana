package main

import (
	"encoding/json"
	"net/http"
)

var CHAINS = map[string]string{
	"mainnet": "https://api.mainnet-beta.solana.com",
	"testnet": "https://testnet.solana.com",
	"devnet":  "https://devnet.solana.com",
}

func call(method interface{}, chain interface{}, params interface{}) {
	var data string
	url := CHAINS[chain]
	headers := map[string]string{"Content-Type": "application/json"}
	if bool(params) {
		data = func() string {
			b, err := json.Marshal(
				map[string]interface{}{
					"jsonrpc": "2.0",
					"id":      1,
					"method":  method,
					"params":  params,
				},
			)
			if err != nil {
				panic(err)
			}
			return string(b)
		}()
	} else {
		data = func() string {
			b, err := json.Marshal(map[string]interface{}{"jsonrpc": "2.0", "id": 1, "method": method})
			if err != nil {
				panic(err)
			}
			return string(b)
		}()
	}
	response, err, err := http.Post(url, headers, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	response
}

func getAccountInfo(pubkey interface{}, chain interface{}) bool {
	response := call("getAccountInfo", chain, []interface{}{pubkey})
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getBalance(pubkey interface{}, chain interface{}) bool {
	response := call("getBalance", chain, []interface{}{pubkey})
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getBlockCommitment(slot interface{}, chain interface{}) bool {
	response := call("getBlockCommitment", chain, []interface{}{slot})
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getBlockTime(slot interface{}, chain interface{}) bool {
	response := call("getBlockTime", chain, []interface{}{slot})
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getClusterNodes(chain interface{}) bool {
	response := call("getClusterNodes", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getConfirmedBlock(slot interface{}, chain interface{}) bool {
	response := call("getConfirmedBlock", chain, []interface{}{slot})
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getConfirmedBlocks(start_slot interface{}, end_slot interface{}, chain interface{}) bool {
	var response interface{}
	if bool(end_slot) {
		response = call("getConfirmedBlocks", chain, []interface{}{start_slot, end_slot})
	} else {
		response = call("getConfirmedBlocks", chain, []interface{}{start_slot})
	}
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getConfirmedSignaturesForAddress(
	pubkey interface{},
	start_slot interface{},
	end_slot interface{},
	chain interface{},
) bool {
	response := call(
		"getConfirmedSignaturesForAddress",
		chain,
		[]interface{}{pubkey, start_slot, end_slot},
	)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getConfirmedTransaction(signature interface{}, chain interface{}) bool {
	response := call("getConfirmedTransaction", chain, []interface{}{signature})
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getEpochInfo(chain interface{}) bool {
	response := call("getEpochInfo", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getEpochSchedule(chain interface{}) bool {
	response := call("getEpochSchedule", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getFeeCalculatorForBlockhash(blockhash interface{}, chain interface{}) bool {
	response := call("getFeeCalculatorForBlockhash", chain, []interface{}{blockhash})
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getFeeRateGovernor(chain interface{}) bool {
	response := call("getFeeRateGovernor", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getFirstAvailableBlock(chain interface{}) bool {
	response := call("getFirstAvailableBlock", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getGenesisHash(chain interface{}) bool {
	response := call("getGenesisHash", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getIdentity(chain interface{}) bool {
	response := call("getIdentity", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getInflation(chain interface{}) bool {
	response := call("getInflation", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getLargestAccounts(chain interface{}) bool {
	response := call("getLargestAccounts", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getLeaderSchedule(slot interface{}, chain interface{}) bool {
	var response interface{}
	if bool(slot) {
		response = call("getLeaderSchedule", chain, []interface{}{slot})
	} else {
		response = call("getLeaderSchedule", chain)
	}
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getMinimumBalanceForRentExemption(datalenght interface{}, chain interface{}) bool {
	response := call("getMinimumBalanceForRentExemption", chain, []interface{}{datalenght})
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getProgramAccounts(pubkey interface{}, chain interface{}) bool {
	response := call("getProgramAccounts", chain, []interface{}{pubkey})
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getRecentBlockhash(chain interface{}) bool {
	response := call("getRecentBlockhash", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getSignatureStatuses(signatures interface{}, chain interface{}) bool {
	response := call("getSignatureStatuses", chain, []interface{}{signatures})
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getSlot(chain interface{}) bool {
	response := call("getSlot", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getSlotLeader(chain interface{}) bool {
	response := call("getSlotLeader", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getSlotsPerSegment(chain interface{}) bool {
	response := call("getSlotsPerSegment", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getStoragePubkeysForSlot(chain interface{}) bool {
	response := call("getStoragePubkeysForSlot", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getStorageTurn(chain interface{}) bool {
	response := call("getStorageTurn", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getStorageTurnRate(chain interface{}) bool {
	response := call("getStorageTurnRate", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getTransactionCount(chain interface{}) bool {
	response := call("getTransactionCount", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getTotalSupply(chain interface{}) bool {
	response := call("getTotalSupply", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getVersion(chain interface{}) bool {
	response := call("getVersion", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func getVoteAccounts(chain interface{}) bool {
	response := call("getVoteAccounts", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func minimumLedgerSlot(chain interface{}) bool {
	response := call("minimumLedgerSlot", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func requestAirdrop(pubkey interface{}, lamports interface{}, chain interface{}) bool {
	response := call("requestAirdrop", chain, []interface{}{pubkey, lamports})
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func sendTransaction(signedtx interface{}, chain interface{}) bool {
	response := call("sendTransaction", chain, []interface{}{signedtx})
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func setLogFilter(logfilter interface{}, chain interface{}) bool {
	response := call("setLogFilter", chain, []interface{}{logfilter})
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}

func validatorExit(chain interface{}) bool {
	response := call("validatorExit", chain)
	if response {
		data := response.json()
		return data["result"]
	} else {
		return false
	}
}
