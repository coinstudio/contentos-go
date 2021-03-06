package request

import (
	"context"
	"fmt"
	"github.com/coschain/contentos-go/cmd/wallet-cli/commands/utils"
	"github.com/coschain/contentos-go/cmd/wallet-cli/wallet"
	"github.com/coschain/contentos-go/prototype"
	"github.com/coschain/contentos-go/rpc/pb"
	"math/rand"
	"strings"
	"time"
	"errors"
)

var nameLib = "abcdefghijklmnopqrstuvwxyz01234567890"

func stake(rpcClient grpcpb.ApiServiceClient, act *wallet.PrivAccount, amount uint64) {
	stkop := &prototype.StakeOperation{
		Account:        &prototype.AccountName{Value: act.Name},
		Amount:            amount,
	}

	signTx, err := utils.GenerateSignedTxAndValidate2(rpcClient, []interface{}{stkop}, act)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("Request command: ", fmt.Sprintf("create %s %s", creatorAccount.Name, newAccountName) )

	req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
	resp, err := rpcClient.BroadcastTrx(context.Background(), req)
	if err != nil {
		fmt.Println("stake error:",err)
	} else {
		fmt.Println("Request command: ",
			fmt.Sprintf("stake %s ", act.Name),
			" ",
			fmt.Sprintf("Result: %v", resp))
	}
}

func createAccount(mywallet *wallet.BaseWallet, rpcClient grpcpb.ApiServiceClient, creatorAccount *wallet.PrivAccount, newAccountName string) {

	if creatorAccount == nil {
		GlobalAccountLIst.RLock()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn( len(GlobalAccountLIst.arr) )
		creatorAccount = GlobalAccountLIst.arr[idx]
		GlobalAccountLIst.RUnlock()
	}

	// if not specify a name, generate a random one
	if newAccountName == "" {
		for {
			for i:=0;i<15;i++{
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				idx := r.Intn(len(nameLib))
				newAccountName += string(nameLib[idx])
			}
			if creatorAccount.Name != newAccountName {
				break
			}
		}
	}

	pubKeyStr, privKeyStr, err := mywallet.GenerateNewKey()
	if err != nil {
		fmt.Println(err)
		return
	}
	pubkey, _ := prototype.PublicKeyFromWIF(pubKeyStr)

	acop := &prototype.AccountCreateOperation{
		Fee:            prototype.NewCoin(1),
		Creator:        &prototype.AccountName{Value: creatorAccount.Name},
		NewAccountName: &prototype.AccountName{Value: newAccountName},
		Owner:          pubkey,
	}
	signTx, err := utils.GenerateSignedTxAndValidate2(rpcClient, []interface{}{acop}, creatorAccount)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("Request command: ", fmt.Sprintf("create %s %s", creatorAccount.Name, newAccountName) )

	req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
	resp, err := rpcClient.BroadcastTrx(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		if resp.Invoice.Status == 200 {
			GlobalAccountLIst.Lock()
			obj := &wallet.PrivAccount{
				Account: wallet.Account{Name: newAccountName, PubKey: pubKeyStr},
				PrivKey: privKeyStr,
			}
			GlobalAccountLIst.arr = append(GlobalAccountLIst.arr, obj)
			GlobalAccountLIst.Unlock()
		}

		if strings.Contains(resp.Invoice.ErrorInfo, "Insufficient") {
			err := transfer(rpcClient, GlobalAccountLIst.arr[0], creatorAccount, 5)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = vest(rpcClient, GlobalAccountLIst.arr[0], creatorAccount, 5)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(fmt.Sprintf("====== createaccount from:%v to:%v amount:%v",GlobalAccountLIst.arr[0].Name,creatorAccount.Name,5))
			createAccount(mywallet, rpcClient, creatorAccount, newAccountName)
			return
		}
		if strings.Contains(resp.Invoice.ErrorInfo,"net resource not enough") {
			stake(rpcClient,creatorAccount,1)
		}
		fmt.Println("Request command: ",
			fmt.Sprintf("create %s %s", creatorAccount.Name, newAccountName),
			" ",
			fmt.Sprintf("Result: %v", resp))
	}
	// give new account 1 coin and let him stake
	toAccount := &wallet.PrivAccount{}
	toAccount.Name = newAccountName
	transfer(rpcClient,creatorAccount,toAccount,1)
	stake(rpcClient,toAccount,1)
}

func transfer(rpcClient grpcpb.ApiServiceClient, fromAccount, toAccount  *wallet.PrivAccount, amount int) error {
	if fromAccount == nil {
		GlobalAccountLIst.RLock()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn( len(GlobalAccountLIst.arr) )
		fromAccount = GlobalAccountLIst.arr[idx]
		GlobalAccountLIst.RUnlock()
	}

	if toAccount == nil {
		for {
			GlobalAccountLIst.RLock()
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			idx := r.Intn( len(GlobalAccountLIst.arr) )
			toAccount = GlobalAccountLIst.arr[idx]
			GlobalAccountLIst.RUnlock()
			if fromAccount != toAccount {
				break
			}
		}
	}

	if amount == 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		amount = 1 + r.Intn(5)
	}

	transfer_op := &prototype.TransferOperation{
		From:   &prototype.AccountName{Value: fromAccount.Name},
		To:     &prototype.AccountName{Value: toAccount.Name},
		Amount: prototype.NewCoin(uint64(amount)),
		Memo:   "",
	}
	signTx, err := utils.GenerateSignedTxAndValidate2(rpcClient, []interface{}{transfer_op}, fromAccount)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//fmt.Println("Request command: ", fmt.Sprintf("transfer %s %s %d", fromAccount.Name, toAccount.Name, amount) )

	req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
	resp, err := rpcClient.BroadcastTrx(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		if strings.Contains(resp.Invoice.ErrorInfo, "Insufficient") {
			if fromAccount == GlobalAccountLIst.arr[0] {
				return errors.New("initminer has no money left")
			}
			err := transfer(rpcClient, GlobalAccountLIst.arr[0], fromAccount, 5)
			if err != nil {
				fmt.Println(err)
				return err
			}
			err = vest(rpcClient, GlobalAccountLIst.arr[0], fromAccount, 5)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println(fmt.Sprintf("====== transfer from:%v to:%v amount:%v",GlobalAccountLIst.arr[0].Name,fromAccount.Name,5))
			transfer(rpcClient, fromAccount, toAccount, amount)
			return nil
		}

		if strings.Contains(resp.Invoice.ErrorInfo,"net resource not enough") {
			stake(rpcClient,fromAccount,1)
		}

		fmt.Println("Request command: ",
			fmt.Sprintf("transfer %s %s %d", fromAccount.Name, toAccount.Name, amount),
			" ",
			fmt.Sprintf("Result: %v", resp))
	}
	return nil
}

func vest(rpcClient grpcpb.ApiServiceClient, fromAccount, toAccount  *wallet.PrivAccount, amount int) error {
	if fromAccount == nil {
		GlobalAccountLIst.RLock()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn( len(GlobalAccountLIst.arr) )
		fromAccount = GlobalAccountLIst.arr[idx]
		GlobalAccountLIst.RUnlock()
	}

	if toAccount == nil {
		for {
			GlobalAccountLIst.RLock()
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			idx := r.Intn( len(GlobalAccountLIst.arr) )
			toAccount = GlobalAccountLIst.arr[idx]
			GlobalAccountLIst.RUnlock()
			if fromAccount != toAccount {
				break
			}
		}
	}

	if amount == 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		amount = 1 + r.Intn(10)
	}

	transfer_op := &prototype.TransferToVestingOperation{
		From:   &prototype.AccountName{Value: fromAccount.Name},
		To:     &prototype.AccountName{Value: toAccount.Name},
		Amount: prototype.NewCoin(uint64(amount)),
	}
	signTx, err := utils.GenerateSignedTxAndValidate2(rpcClient, []interface{}{transfer_op}, fromAccount)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//fmt.Println("Request command: ", fmt.Sprintf("transfer vest %s %s %d", fromAccount.Name, toAccount.Name, amount) )

	req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
	resp, err := rpcClient.BroadcastTrx(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		if strings.Contains(resp.Invoice.ErrorInfo, "Insufficient") {
			if fromAccount == GlobalAccountLIst.arr[0] {
				return errors.New("initminer has no money left")
			}
			err := vest(rpcClient, GlobalAccountLIst.arr[0], fromAccount, 10)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println(fmt.Sprintf("====== vest from:%v to:%v amount:%v",GlobalAccountLIst.arr[0].Name,fromAccount.Name,5))
			vest(rpcClient, fromAccount, toAccount, amount)
			return nil
		}

		if strings.Contains(resp.Invoice.ErrorInfo,"net resource not enough") {
			stake(rpcClient,fromAccount,1)
		}

		fmt.Println("Request command: ",
			fmt.Sprintf("transfer vest %s %s %d", fromAccount.Name, toAccount.Name, amount),
			" ",
			fmt.Sprintf("Result: %v", resp))
	}
	return nil
}


func postArticle(rpcClient grpcpb.ApiServiceClient, authorAccount *wallet.PrivAccount) {

	if authorAccount == nil {
		GlobalAccountLIst.RLock()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn( len(GlobalAccountLIst.arr) )
		authorAccount = GlobalAccountLIst.arr[idx]
		GlobalAccountLIst.RUnlock()
	}

	var tag = ""
	var title = ""
	var content = ""
	beneficiaries := []*prototype.BeneficiaryRouteType{}
	for i:=0;i<10;i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn(len(nameLib))
		tag += string(nameLib[idx])
	}
	for i:=0;i<20;i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn(len(nameLib))
		title += string(nameLib[idx])
	}
	for i:=0;i<1024;i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn(len(nameLib))
		content += string(nameLib[idx])
	}

	uuid := utils.GenerateUUID( authorAccount.Name + title )
	post_op := &prototype.PostOperation{
		Uuid:          uuid,
		Owner:         &prototype.AccountName{Value: authorAccount.Name},
		Title:         title,
		Content:       content,
		Tags:          []string{tag},
		Beneficiaries: beneficiaries,
	}
	signTx, err := utils.GenerateSignedTxAndValidate2(rpcClient, []interface{}{post_op}, authorAccount)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("Request command: ", fmt.Sprintf("%s post an article", authorAccount.Name) )

	req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
	resp, err := rpcClient.BroadcastTrx(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		if resp.Invoice.Status == 200 {
			PostIdList.Lock()
			PostIdList.arr = append(PostIdList.arr, uuid)
			PostIdList.Unlock()
		}

		if strings.Contains(resp.Invoice.ErrorInfo, "Insufficient") {
			err := transfer(rpcClient, GlobalAccountLIst.arr[0], authorAccount, 5)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = vest(rpcClient, GlobalAccountLIst.arr[0], authorAccount, 5)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(fmt.Sprintf("====== post transfer from:%v to:%v amount:%v",GlobalAccountLIst.arr[0].Name,authorAccount.Name,5))
			postArticle(rpcClient, authorAccount)
			return
		}
		if strings.Contains(resp.Invoice.ErrorInfo,"net resource not enough") {
			stake(rpcClient,authorAccount,1)
		}

		fmt.Println("Request command: ",
			fmt.Sprintf("%s post an article", authorAccount.Name),
			" ",
			fmt.Sprintf("Result: %v", resp))
	}
}

func follow(rpcClient grpcpb.ApiServiceClient, followerAccount, followingAccount *wallet.PrivAccount) {
	if followerAccount == nil {
		GlobalAccountLIst.RLock()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn( len(GlobalAccountLIst.arr) )
		followerAccount = GlobalAccountLIst.arr[idx]
		GlobalAccountLIst.RUnlock()
	}

	if followingAccount == nil {
		for {
			GlobalAccountLIst.RLock()
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			idx := r.Intn( len(GlobalAccountLIst.arr) )
			followingAccount = GlobalAccountLIst.arr[idx]
			GlobalAccountLIst.RUnlock()
			if followerAccount != followingAccount {
				break
			}
		}
	}

	follow_op := &prototype.FollowOperation{
		Account:  &prototype.AccountName{Value: followerAccount.Name},
		FAccount: &prototype.AccountName{Value: followingAccount.Name},
		Cancel:   false,
	}

	signTx, err := utils.GenerateSignedTxAndValidate2(rpcClient, []interface{}{follow_op}, followerAccount)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("Request command: ", fmt.Sprintf("follow %s %s", followerAccount.Name, followingAccount.Name) )

	req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
	resp, err := rpcClient.BroadcastTrx(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		if strings.Contains(resp.Invoice.ErrorInfo,"net resource not enough") {
			stake(rpcClient,followerAccount,1)
		}
		fmt.Println("Request command: ",
			fmt.Sprintf("follow %s %s", followerAccount.Name, followingAccount.Name),
			" ",
			fmt.Sprintf("Result: %v", resp))
	}
}

func voteArticle(rpcClient grpcpb.ApiServiceClient, voterAccount *wallet.PrivAccount, postId uint64) {
	if voterAccount == nil {
		GlobalAccountLIst.RLock()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn( len(GlobalAccountLIst.arr) )
		voterAccount = GlobalAccountLIst.arr[idx]
		GlobalAccountLIst.RUnlock()
	}

	if postId == 0 {
		PostIdList.RLock()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn( len(PostIdList.arr) )
		postId = PostIdList.arr[idx]
		PostIdList.RUnlock()
	}

	vote_op := &prototype.VoteOperation{
		Voter: &prototype.AccountName{Value: voterAccount.Name},
		Idx:   postId,
	}

	signTx, err := utils.GenerateSignedTxAndValidate2(rpcClient, []interface{}{vote_op}, voterAccount)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("Request command: ", fmt.Sprintf("vote %s %d", voterAccount.Name, postId) )

	req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
	resp, err := rpcClient.BroadcastTrx(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		if strings.Contains(resp.Invoice.ErrorInfo, "Insufficient") {
			err := transfer(rpcClient, GlobalAccountLIst.arr[0], voterAccount, 5)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = vest(rpcClient, GlobalAccountLIst.arr[0], voterAccount, 5)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(fmt.Sprintf("====== vote from:%v to:%v amount:%v",GlobalAccountLIst.arr[0].Name,voterAccount.Name,5))
			voteArticle(rpcClient, voterAccount, postId)
			return
		}
		if strings.Contains(resp.Invoice.ErrorInfo,"net resource not enough") {
			stake(rpcClient,voterAccount,1)
		}

		fmt.Println("Request command: ",
			fmt.Sprintf("vote %s %d", voterAccount.Name, postId),
			" ",
			fmt.Sprintf("Result: %v", resp))
	}
}

func replyArticle(rpcClient grpcpb.ApiServiceClient, fromAccount *wallet.PrivAccount, postId uint64) {
	if fromAccount == nil {
		GlobalAccountLIst.RLock()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn( len(GlobalAccountLIst.arr) )
		fromAccount = GlobalAccountLIst.arr[idx]
		GlobalAccountLIst.RUnlock()
	}

	if postId == 0 {
		PostIdList.RLock()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn( len(PostIdList.arr) )
		postId = PostIdList.arr[idx]
		PostIdList.RUnlock()
	}

	var content = ""
	for i:=0;i<128;i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn(len(nameLib))
		content += string(nameLib[idx])
	}

	uuid := utils.GenerateUUID(fromAccount.Name)
	beneficiaries := []*prototype.BeneficiaryRouteType{}

	reply_op := &prototype.ReplyOperation{
		Uuid:          uuid,
		Owner:         &prototype.AccountName{Value: fromAccount.Name},
		Content:       content,
		ParentUuid:    postId,
		Beneficiaries: beneficiaries,
	}

	signTx, err := utils.GenerateSignedTxAndValidate2(rpcClient, []interface{}{reply_op}, fromAccount)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("Request command: ", fmt.Sprintf("reply %s %d", fromAccount.Name, postId) )

	req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
	resp, err := rpcClient.BroadcastTrx(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		if strings.Contains(resp.Invoice.ErrorInfo,"net resource not enough") {
			stake(rpcClient,fromAccount,1)
		}
		fmt.Println("Request command: ",
			fmt.Sprintf("reply %s %d", fromAccount.Name, postId),
			" ",
			fmt.Sprintf("Result: %v", resp))
	}
}

func callContract(rpcClient grpcpb.ApiServiceClient, fromAccount  *wallet.PrivAccount) error {
	if fromAccount == nil {
		GlobalAccountLIst.RLock()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		idx := r.Intn( len(GlobalAccountLIst.arr) )
		fromAccount = GlobalAccountLIst.arr[idx]
		GlobalAccountLIst.RUnlock()
	}

	param := fmt.Sprintf(" [\"%v\"] ", fromAccount.Name)

	applyOp := &prototype.ContractApplyOperation{
		Caller:   &prototype.AccountName{Value: fromAccount.Name},
		Owner:    &prototype.AccountName{Value: "initminer"},
		Amount:   &prototype.Coin{Value: 0},
		Contract: "registercount",
		Params:   param,
		Method:   "checkincount",
	}

	signTx, err := utils.GenerateSignedTxAndValidate2(rpcClient, []interface{}{applyOp}, fromAccount)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
	resp, err := rpcClient.BroadcastTrx(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		if strings.Contains(resp.Invoice.ErrorInfo,"net resource not enough") {
			stake(rpcClient,fromAccount,1)
		}
		fmt.Println("Request command: ",
			fmt.Sprintf("callContract %s %s %d", fromAccount.Name),
			" ",
			fmt.Sprintf("Result: %v", resp))
	}
	return nil
}