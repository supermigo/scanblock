package contractparse

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"scanchainlist/contract"
	"scanchainlist/scanblock/eth"
	"scanchainlist/utils"
	"scanchainlist/xlog"
	"strings"
)

type IdoConsumer struct {
	ContractAddr string
	Contract     *contract.Ido
}

func NewIdoConsumer(addr string, ethClientPtr *ethclient.Client) *IdoConsumer {
	result := new(IdoConsumer)
	result.ContractAddr = addr
	result.Contract, _ = contract.NewIdo(common.HexToAddress(addr), ethClientPtr)
	return result
}

func (ge *IdoConsumer) GetConsumer() ([]*eth.EventConsumer, error) {
	parsed, _ := abi.JSON(strings.NewReader(contract.IdoMetaData.ABI))
	return []*eth.EventConsumer{
		{
			Address: common.HexToAddress(ge.ContractAddr),
			Topic: crypto.Keccak256Hash(
				[]byte(parsed.Events["debug"].Sig),
			),
			ParseEvent: ge.Debug,
		},
	}, nil
}

func (ge *IdoConsumer) GetFilterQuery() []ethereum.FilterQuery {
	parsed, _ := abi.JSON(strings.NewReader(contract.IdoABI))
	return []ethereum.FilterQuery{
		ethereum.FilterQuery{
			Addresses: []common.Address{common.HexToAddress(ge.ContractAddr)},
			Topics: [][]common.Hash{{
				crypto.Keccak256Hash(
					[]byte(parsed.Events["debug"].Sig),
				),
			}},
		},
	}

}

type EventPool_send struct {
	Name   string `json:"name"`
	Msg    string `json:"msg"`
	Addr   string `json:"addr"`
	Exec   string `json:"exec"`
	Number string `json:"number"`
}

func (ge *IdoConsumer) Debug(l types.Log) error {
	sContractTransfer, error := ge.Contract.ParseDebug(l)
	sContractTransfer = sContractTransfer
	if error == nil {
		//xlog.CInfo("ido", l.BlockNumber, " >>> ", l.TxHash.Hex())
		//如果有指定的钱
		var handle_send EventPool_send
		handle_send.Name = sContractTransfer.Name
		address := common.HexToAddress(sContractTransfer.Msg.String()).String()
		handle_send.Msg = address
		address = common.HexToAddress(sContractTransfer.Addr.String()).String()
		handle_send.Addr = address
		handle_send.Exec = sContractTransfer.Exec.String()
		handle_send.Number = sContractTransfer.Number.String()

		str_tmp := utils.DataZipString(handle_send)
		str_tmp = str_tmp
		xlog.CInfo(l.BlockNumber, utils.Struct2JsonString(handle_send))
		//url_str = Http_server + "/process?id=pool&data=" + str_tmp
		//fmt.Println(url_str)
		//fmt.Println(lib1.Get_HTTP(url_str))
	} else {
		xlog.CError(error.Error())
	}
	return nil
}
