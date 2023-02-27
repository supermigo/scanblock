package contractparse

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	erc20 "github.com/nattaponra/go-abi/erc20/contract"
	"scanchainlist/scanblock/eth"
	"scanchainlist/xlog"
	"strings"
)

type Erc20Consumer struct {
	ContractAddr string
	Contract     *erc20.Contract
}

func NewErc20Consumer(addr string, ethClientPtr *ethclient.Client) *Erc20Consumer {
	result := new(Erc20Consumer)
	result.ContractAddr = addr
	result.Contract, _ = erc20.NewContract(common.HexToAddress(addr), ethClientPtr)
	return result
}

func (ge *Erc20Consumer) GetConsumer() ([]*eth.EventConsumer, error) {
	parsed, _ := abi.JSON(strings.NewReader(erc20.ContractABI))
	return []*eth.EventConsumer{
		{
			Address: common.HexToAddress(ge.ContractAddr),
			Topic: crypto.Keccak256Hash(
				[]byte(parsed.Events["Transfer"].Sig),
			),
			ParseEvent: ge.Transfer,
		},
	}, nil
}

func (ge *Erc20Consumer) GetFilterQuery() []ethereum.FilterQuery {
	parsed, _ := abi.JSON(strings.NewReader(erc20.ContractABI))
	return []ethereum.FilterQuery{
		ethereum.FilterQuery{
			Addresses: []common.Address{common.HexToAddress(ge.ContractAddr)},
			Topics: [][]common.Hash{{
				crypto.Keccak256Hash(
					[]byte(parsed.Events["Transfer"].Sig),
				),
			}},
		},
	}

}

func (ge *Erc20Consumer) Transfer(l types.Log) error {
	sContractTransfer, error := ge.Contract.ParseTransfer(l)
	if error == nil {
		//如果有指定的钱
		xlog.CInfo("txHash>>", l.TxHash)
		xlog.CInfo("txHash>>", l.BlockHash)
		xlog.CInfo("sContractTransfer.From>>", sContractTransfer.From)
		xlog.CInfo("sContractTransfer.To>>", sContractTransfer.To)
		xlog.CInfo("sContractTransfer.Value>>", sContractTransfer.Value)
		//	fmt.Println("sContractTransfer.Raw>>", sContractTransfer.Raw)
		nValue, _ := ge.Contract.Decimals(nil)
		xlog.CInfo("Decimals>>", nValue)
	}
	return nil
}
