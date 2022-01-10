package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	//"strconv"
)

var CHAINS = map[string]string{
	"mainnet": "https://api.mainnet-beta.solana.com",
	"testnet": "https://testnet.solana.com",
	"devnet":  "https://devnet.solana.com",
}

type Sol struct {
	seedphrase interface{}
	pubkey     string
	chain      interface{}
}

func NewSol(seedphrase interface{}, chain interface{}) (S *Sol) {
	S = new(Sol)
	if seedphrase != nil {
		S.seedphrase = seedphrase
		func() {
			defer func() {
				if r := recover(); r != nil {
					if err, ok := r.(error); ok {
						if strings.HasPrefix(err.Error(), "CalledProcessError") {
							grepexc := err
							error := true
							return
						}
					}
					panic(r)
				}
			}()
			S.pubkey = string(check_output(func() string {
				var buf bytes.Buffer
				err := template.Must(template.New("f").Parse("printf \"{{.expr1}}\n\n\" ")).Execute(&buf, map[string]interface{}{"expr1": S.seedphrase})
				if err != nil {
					panic(err)
				}
				return buf.String()
			}()+"| solana-keygen pubkey ASK", true, PIPE))[:-1]
			error := false
		}()
		if error {
			panic(fmt.Errorf("Exception: %v", "keygen error"))
		}
	} else {
		p := Popen("solana-keygen new --no-outfile --no-passphrase", true, PIPE)
		//text = p.stderr.read().decode().split('\n')
		S.pubkey = text[1][8:]
		S.seedphrase = text[4]
	}
	S.chain = chain
	return
}

func set_chain(self interface{}, chain interface{}) {
	self.chain = chain
}

func balance(self interface{}) {
	func() float64 {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					if strings.HasPrefix(err.Error(), "CalledProcessError") {
						grepexc := err
						return
					}
				}
				panic(r)
			}
		}()
		return func() float64 {
			//return float(check_output(#f'solana balance {self.pubkey} ' +\
			//f'--url {CHAINS[self.chain]}',
			i, err := strconv.ParseFloat(string(check_output(true))[:-5], 64)
			if err != nil {
				panic(err)
			}
			return i
		}()
	}()
	panic(fmt.Errorf("Exception: %v", "balance error"))
}

func airdrop(self interface{}, amount interface{}, wait interface{}) {
	if self.chain != "devnet" {
		panic(fmt.Errorf("Exception: %v", "airdrop is not allowed in this chain"))
	}
	if bool(wait) {
		//#f'solana airdrop {amount} {self.pubkey} ' +\
		//f'--url {CHAINS[self.chain]}',
		call(true, PIPE, PIPE)
	} else {
		//f'solana airdrop {amount} {self.pubkey} ' +\
		//f'--url {CHAINS[self.chain]}',
		Popen(true, PIPE, PIPE)
	}
}

func (S *Sol) transfer(to interface{}, amount interface{}, wait interface{}) {
	if to.__class__.__name__ == "Sol" {
		to := to.pubkey
	}
	/*if wait:
		text = f'printf "{S.seedphrase}\n\n{S.seedphrase}' +\
	    '\n\n" | solana transfer --fee-payer ' +\
	    f'ASK --from ASK {to} {amount} ' +\
	    f'--url {CHAINS[S.chain]}'
		else:
		text = f'printf "{S.seedphrase}\n\n{S.seedphrase}' +\
	   '\n\n" | solana transfer --no-wait --fee-payer ' +\
	    f'ASK --from ASK {to} {amount} ' +\
	    f'--url {CHAINS[S.chain]}'
	*/
	func() string {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					if strings.HasPrefix(err.Error(), "CalledProcessError") {
						grepexc := err
						return
					}
				}
				panic(r)
			}
		}()
		return string(check_output(text, true, PIPE))[:-1] //.decode()
	}()
	panic(fmt.Errorf("Exception: %v", "transfer error"))
}

func create_stake_account(self interface{}, amount interface{}) {
	stake_account := Sol()
	/*
	 text = f'printf "{self.seedphrase}\n\n{self.seedphrase}\n\n' +\
	               f'{self.seedphrase}\n\n{self.seedphrase}\n\n' +\
	               f'{stake_account.seedphrase}" | ' +\
	               'solana create-stake-account --from ASK ASK 100' +\
	               ' --stake-authority ASK --withdraw-authority ASK --fee-payer ASK'
	*/
	call(text, true, PIPE, PIPE)
	func() {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					if strings.HasPrefix(err.Error(), "CalledProcessError") {
						grepexc := err
						return
					}
				}
				panic(r)
			}
		}()
		stake_account.pubkey
	}()
	panic(fmt.Errorf("Exception: %v", "stake account error"))
}

func delegate_stake(self interface{}, stake_account interface{}, vote_account interface{}) {
	/*
	   text = f'printf "{self.seedphrase}\n\n{self.seedphrase}\n\n" | ' +\
	               f'solana delegate-stake {stake_account} {vote_account} ' +\
	               '--stake-authority ASK --fee-payer ASK'
	*/
	var error bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					if strings.HasPrefix(err.Error(), "CalledProcessError") {
						grepexc := err
						error = true
						return
					}
				}
				panic(r)
			}
		}()
		call(text, true, PIPE, PIPE)
		error = false
	}()
	if error {
		panic(fmt.Errorf("Exception: %v", "delegate stake error"))
	}
}

func deactivate_stake(self interface{}, stake_account interface{}) {
	/*
			text = f'printf "{self.seedphrase}\n\n{self.seedphrase}\n\n" | ' +\
		f'solana deactivate-stake {stake_account} ' +\
		'--stake-authority ASK --fee-payer ASK'
	*/
	var error bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					if strings.HasPrefix(err.Error(), "CalledProcessError") {
						grepexc := err
						error = true
						return
					}
				}
				panic(r)
			}
		}()
		call(text, true, PIPE, PIPE)
		error = false
	}()
	if error {
		panic(fmt.Errorf("Exception: %v", "deactivate stake error"))
	}
}

func withdraw_stake(self interface{}, stake_account interface{}, amount interface{}) {
	/*
			text = f'printf "{self.seedphrase}\n\n{self.seedphrase}\n\n" | ' +\
		               f'solana withdraw-stake {stake_account} {self.pubkey}' +\
		               f' {amount} --withdraw-authority ASK --fee-payer ASK'
	*/
	var error bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					if strings.HasPrefix(err.Error(), "CalledProcessError") {
						grepexc := err
						error = true
						return
					}
				}
				panic(r)
			}
		}()
		call(text, true, PIPE, PIPE)
		error = false
	}()
	if error {
		panic(fmt.Errorf("Exception: %v", "withdraw stake error"))
	}
}
