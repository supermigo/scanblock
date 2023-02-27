package contractparse

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

type IdoPoolConsumer struct {
	ContractAddr  string
	Contract      *contract.Idfactory
	ethCli        *ethclient.Client
	DynimacCustom chan eth.IEventConsumer
}

func NewIdoPoolConsumer(addr string, ethClientPtr *ethclient.Client) *IdoPoolConsumer {
	result := new(IdoPoolConsumer)
	result.ContractAddr = addr
	result.ethCli = ethClientPtr
	result.Contract, _ = contract.NewIdfactory(common.HexToAddress(addr), ethClientPtr)
	return result
}

func (ge *IdoPoolConsumer) GetConsumer() ([]*eth.EventConsumer, error) {
	parsed, _ := abi.JSON(strings.NewReader(contract.IdfactoryABI))
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

func (ge *IdoPoolConsumer) GetFilterQuery() []ethereum.FilterQuery {
	parsed, _ := abi.JSON(strings.NewReader(contract.IdfactoryABI))
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

// 用于 HTTP 传输的数据结构
type EventLog_send struct {
	Name  string              `json:"name"`
	Msg   string              `json:"msg"`
	Owner string              `json:"owner"`
	Admin string              `json:"admin"`
	Text  string              `json:"text"`
	Group string              `json:"group"`
	Mode  string              `json:"mode"`
	Data  IDOprojectInfo_send `json:"data"`
}

// 一次性结构化传输
type IDOprojectInfo_send struct {
	TokenA           string `json:"tokenA"`
	TokenB           string `json:"tokenB"`
	ProjectText      string `json:"projectText"`
	ProjectType      string `json:"projectType"`
	GroupID          string `json:"groupID"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	InTokenCapacity  string `json:"inTokenCapacity"`
	InTokenAmount    string `json:"inTokenAmount"`
	OutTokenCapacity string `json:"outTokenCapacity"`
	MaxExchange      string `json:"maxExchange"`
	Exchange         string `json:"exchange"`
	DecimalA         string `json:"decimalA"`
	DecimalB         string `json:"decimalB"`
	LockNum          string `json:"lockNum"`
	TimeList         []uint `json:"timeList"`
	TokenNameA       string `json:"TokenNameA"`
	SymbolA          string `json:"SymbolA"`
	DecimalsA        string `json:"DecimalsA"`
	TotalSupplyA     string `json:"TotalSupplyA"`
	TokenNameB       string `json:"TokenNameB"`
	SymbolB          string `json:"SymbolB"`
	DecimalsB        string `json:"DecimalsB"`
	TotalSupplyB     string `json:"TotalSupplyB"`
}

func (ge *IdoPoolConsumer) Debug(l types.Log) error {
	var handle_send EventLog_send
	sContractTransfer, error := ge.Contract.ParseDebug(l)
	sContractTransfer = sContractTransfer
	//xlog.CInfo(l.TxHash, l.BlockNumber)
	if error == nil {
		xlog.CInfo("idofactor", l.BlockNumber)
		//如果有指定的钱
		//如果发现是新的合约，1 需要提交给服务端， 并且让listen 增加一个新的ido的监听合约。
		//逐个解析数据并且填充 HTTP 发送结构
		address := common.HexToAddress(sContractTransfer.Msg.String()).String()
		handle_send.Msg = address
		address = common.HexToAddress(sContractTransfer.Owner.String()).String()
		handle_send.Owner = address
		address = common.HexToAddress(sContractTransfer.Admin.String()).String()
		handle_send.Admin = address

		handle_send.Name = sContractTransfer.Name
		handle_send.Text = sContractTransfer.Text
		handle_send.Group = sContractTransfer.Group

		value := sContractTransfer.Mode.String()
		handle_send.Mode = value

		//查询合约中的对应消息结构
		handle_send.Data = ge.contract_idopool(handle_send.Msg)
		xlog.CInfo(utils.Struct2JsonString(handle_send))
		str_tmp := utils.DataZipString(handle_send)
		str_tmp = str_tmp
		strdata := NewIdoConsumer(strings.ToLower(sContractTransfer.Msg.String()), ge.ethCli)
		ge.DynimacCustom <- strdata

	} else {
		xlog.CError("now ：", l.TxHash, l.BlockNumber, error.Error())
	}
	return nil
}
func (s *IdoPoolConsumer) contract_idopool(address string) (handle_send_project IDOprojectInfo_send) {
	contractAddr := common.HexToAddress(address)
	contractAddrPtr, _ := contract.NewIdo(contractAddr, s.ethCli)
	handle_project, _ := contractAddrPtr.QueryProjectInfo(&bind.CallOpts{Pending: true})

	address = common.HexToAddress(handle_project.TokenA.String()).String()
	handle_send_project.TokenA = address
	address = common.HexToAddress(handle_project.TokenB.String()).String()
	handle_send_project.TokenB = address
	handle_send_project.ProjectText = handle_project.ProjectText
	handle_send_project.ProjectType = handle_project.ProjectType.String()
	handle_send_project.GroupID = handle_project.GroupID

	handle_send_project.StartTime = handle_project.StartTime.String()
	handle_send_project.EndTime = handle_project.EndTime.String()

	handle_send_project.InTokenCapacity = handle_project.InTokenCapacity.String()

	handle_send_project.InTokenAmount = handle_project.InTokenAmount.String()
	handle_send_project.OutTokenCapacity = handle_project.OutTokenCapacity.String()
	handle_send_project.MaxExchange = handle_project.MaxExchange.String()
	handle_send_project.Exchange = handle_project.Exchange.String()
	handle_send_project.DecimalA = handle_project.DecimalA.String()
	handle_send_project.DecimalB = handle_project.DecimalB.String()
	handle_send_project.LockNum = handle_project.LockNum.String()
	for _, value := range handle_project.TimeList {
		handle_send_project.TimeList = append(handle_send_project.TimeList, uint(value.Uint64()))
	}
	return handle_send_project
}
