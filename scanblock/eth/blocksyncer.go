package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/panjf2000/ants/v2"
	"math/big"
	"scanchainlist/daemon"
	"scanchainlist/db"
	"scanchainlist/xlog"
	"sync"
	"time"
)

type EthListener struct {
	EthInfoDao               *db.SyncStatus
	Log                      chan types.Log //每个eth时间
	EthClient                *ethclient.Client
	EventConsumerMap         map[string]*EventConsumer
	TxMonitors               map[common.Address]ITxMonitor
	errC                     chan error
	blockTime                uint64
	blockOffset              int64
	chainID                  int32
	syncMutext               sync.RWMutex
	ChannelStringForCustomer chan IEventConsumer
}

func NewEthListener(
	ethInfo *db.SyncStatus,
	ethClient *ethclient.Client,
	blockTime uint64,
	blockOffset int64,
	ChainID int32,
) *EthListener {
	st := &EthListener{
		EthInfoDao:               ethInfo,
		EthClient:                ethClient,
		EventConsumerMap:         make(map[string]*EventConsumer),
		TxMonitors:               make(map[common.Address]ITxMonitor),
		Log:                      make(chan types.Log),
		errC:                     make(chan error),
		blockTime:                blockTime,
		blockOffset:              blockOffset,
		chainID:                  ChainID,
		ChannelStringForCustomer: make(chan IEventConsumer, 0),
	}
	return st
}
func (st *EthListener) ReadConsumer(ctx context.Context) {
	for {
		select {
		case address := <-st.ChannelStringForCustomer:
			st.RegisterConsumer(address)
		case <-ctx.Done():
			return
		}
	}
}

func KeyFromBEConsumer(address string, topic string) string {
	return fmt.Sprintf("%s:%s", address, topic)
}

func (s *EthListener) RegisterConsumer(consumer IEventConsumer) error {
	s.syncMutext.Lock()
	defer s.syncMutext.Unlock()
	consumerHandler, err := consumer.GetConsumer()
	if err != nil {
		xlog.CError("[eth listener] Unable to get consumer", err.Error())
		return err
	}
	for i := 0; i < len(consumerHandler); i++ {
		xlog.CDebugf("Key for comsumer: %+v", KeyFromBEConsumer(consumerHandler[i].Address.Hex(), consumerHandler[i].Topic.Hex()))
		if _, ok := s.EventConsumerMap[KeyFromBEConsumer(consumerHandler[i].Address.Hex(), consumerHandler[i].Topic.Hex())]; !ok {
			s.EventConsumerMap[KeyFromBEConsumer(consumerHandler[i].Address.Hex(), consumerHandler[i].Topic.Hex())] = consumerHandler[i]
		}
	}
	return nil
}

func (s *EthListener) RegisterTxMonitor(monitor ITxMonitor) error {
	if monitor == nil {
		err := fmt.Errorf("Null monitor")
		xlog.CError("[eth listener] Register nil monitor", err.Error())
		return err
	}
	address := monitor.MonitoredAddress()
	if _, ok := s.TxMonitors[address]; !ok {
		s.TxMonitors[address] = monitor
		xlog.CDebug("Monitor for " + fmt.Sprintf("0x%x", address) + " Added")
		return nil
	}
	return fmt.Errorf("Monitor for " + fmt.Sprintf("0x%x", address) + " already existed")
}

func (s *EthListener) Start(ctx context.Context) {
	daemon.BootstrapDaemons(ctx, s.Scan)
}
func (s *EthListener) Scan(parentContext context.Context) (daemonfun daemon.Daemon, err error) {
	daemonfun = func() {
		for {
			select {
			case <-parentContext.Done():
				xlog.CDebug("[eth listener] Stop listening")
				return
			default:
				sysInfo, err := db.DB.GetSyncStatus(s.chainID)
				if err != nil {
					xlog.CError("[eth_listener] can't get system info", err.Error())
					continue
				}

				header, err := s.EthClient.HeaderByNumber(parentContext, nil)
				if err != nil {
					xlog.CError("[eth_listener] can't get head by number, possibly due to rpc node failure", err.Error())
					continue
				}
				fmt.Println("header.Number", header.Number)
				currBlock := header.Number
				scannedBlock := sysInfo.LastEthBlockBigInt()
				diff := big.NewInt(0).Sub(currBlock, scannedBlock)
				fmt.Println(diff)
				for begin := scannedBlock; currBlock.Cmp(begin) > 0; begin = begin.Add(begin, big.NewInt(1)) {
					block, err := s.EthClient.BlockByNumber(context.Background(), begin)
					if err != nil {
						xlog.CError(err.Error())
						begin = big.NewInt(0).Sub(begin, big.NewInt(1))
						continue
					}
					var wg sync.WaitGroup

					fmt.Println(block.Number(), len(block.Transactions()))
					p, _ := ants.NewPoolWithFunc(20, func(input interface{}) {
						tx := input.(*types.Transaction)
						receipt, err := s.EthClient.TransactionReceipt(context.Background(), tx.Hash())
						if err != nil {
							xlog.CError(err.Error())
						} else {
							for _, vLog := range receipt.Logs {
								consumer, isExisted := s.matchEvent(*vLog)
								if isExisted {
									err := consumer.ParseEvent(*vLog)
									if err != nil {
										xlog.CErrorf("[eth_client] Consume event error")
									}
								}
							}
						}
						wg.Done()
					})
					for key := 0; key < len(block.Transactions()); key++ {
						wg.Add(1)
						tx := block.Transactions()[key]
						_ = p.Invoke(tx)
					}
					wg.Wait()
					p.Release()
					sysInfo.LastEthBlockRecorded = begin.Uint64()
					db.DB.UpdateSyncStatusWithBlockNumber(s.chainID, sysInfo.LastEthBlockRecorded)
				}
				header2, _ := s.EthClient.HeaderByNumber(parentContext, nil)
				if header2.Number == currBlock {
					daemon.SleepContext(parentContext, time.Second*time.Duration(s.blockTime))
				}
			}

		}
	}
	return daemonfun, nil
}

func (s *EthListener) matchEvent(vLog types.Log) (*EventConsumer, bool) {
	if len(vLog.Topics) == 0 {
		return nil, false
	}
	key := KeyFromBEConsumer(vLog.Address.Hex(), vLog.Topics[0].Hex())
	s.syncMutext.RLock()
	defer s.syncMutext.RUnlock()
	consumer, isExisted := s.EventConsumerMap[key]
	return consumer, isExisted
}

func (s *EthListener) consumeEvent(vLog types.Log) {
	consumer, isExisted := s.matchEvent(vLog)
	if isExisted {
		err := consumer.ParseEvent(vLog)
		if err != nil {
			xlog.CErrorf("[eth_client] Consume event error")
		}
	}
}
