package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"scanchainlist/Service"
	"scanchainlist/contractparse"
	"scanchainlist/db"
	"scanchainlist/scanblock/eth"
	"scanchainlist/xlog"
	"sync"
	"syscall"
	"time"
)

const idofactoryAddress = ""

func main() {
	deption := xlog.DefaultOption()
	deption.MaxSize = 1
	deption.Level = zap.InfoLevel.String()

	xlog.InitLevel(deption)
	db.Init()

	db.DB.UpdateSyncStatusWithBlockNumber(5, 8125387)
	db.DB.ResetUpdateSyncStatusWithBlockNumber(5, 8125387)
	db.DB.UpdateSyncStatusWithBatchNumber(5, 0)

	db.DB.UpdateSyncStatusWithBlockNumber(97, 27600498)
	db.DB.ResetUpdateSyncStatusWithBlockNumber(97, 27600498)
	db.DB.UpdateSyncStatusWithBatchNumber(97, 0)

	tList := Service.GetSystemHaveIdo()

	ethclientPtr := openEthListen()
	defer ethclientPtr.Close()
	bscClientPtr := openBscListen("0xbecd39a8f9f52869e8d574a84a8060ce4e7cf43f", tList)
	defer bscClientPtr.Close()

	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
EXIT:
	for {
		sig := <-sc
		xlog.CDebugf("接收到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	time.Sleep(time.Second)
	os.Exit(state)
}
func openBscListen(iodfactory string, liststring []string) *ethclient.Client {
	ethCli, err := ethclient.Dial("https://data-seed-prebsc-2-s2.binance.org:8545")
	if err != nil {
		xlog.CError("Unable to connect to ethereum RPC server")
		return nil
	}
	ethListen := eth.NewEthListener(new(db.SyncStatus), ethCli, 5, 10, 97)
	idoFactoryCustomer := contractparse.NewIdoPoolConsumer(iodfactory, ethCli)
	idoFactoryCustomer.DynimacCustom = ethListen.ChannelStringForCustomer
	ethListen.RegisterConsumer(idoFactoryCustomer)
	for _, value := range liststring {
		idoCustomer := contractparse.NewIdoConsumer(value, ethCli)
		ethListen.RegisterConsumer(idoCustomer)
	}
	ctx := context.Background()
	go func() {
		ethListen.ReadConsumer(ctx)
	}()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		ethListen.Start(context.Background())
		wg.Done()
	}()
	return ethCli
}

func openEthListen() *ethclient.Client {
	ethCli, err := ethclient.Dial("https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	if err != nil {
		xlog.CDebug("Unable to connect to ethereum RPC server")
		return nil
	}
	ethListen := eth.NewEthListener(new(db.SyncStatus), ethCli, 2, 10, 5)
	erc20Customer := contractparse.NewErc20Consumer("0x1d3a434eeac22d9935c14557d04b6b2d466696f7", ethCli)
	ethListen.RegisterConsumer(erc20Customer)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		ethListen.Start(context.Background())
		wg.Done()
	}()

	return ethCli

}
