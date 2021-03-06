package commands

import (
	"github.com/coschain/contentos-go/cmd/wallet-cli/commands/utils/mock"
	"github.com/coschain/contentos-go/cmd/wallet-cli/wallet"
	"github.com/coschain/contentos-go/cmd/wallet-cli/wallet/mock"
	"github.com/coschain/contentos-go/rpc/mock_grpcpb"
	"github.com/coschain/contentos-go/rpc/pb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostWithoutBeneficiaries(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)
	mywallet := mock_wallet.NewMockWallet(ctrl)
	myassert := assert.New(t)
	passwordReader := mock_utils.NewMockPasswordReader(ctrl)
	cmd := PostCmd()
	cmd.SetContext("wallet", mywallet)
	cmd.SetContext("rpcclient", client)
	cmd.SetContext("preader", passwordReader)
	for _, child := range cmd.Commands() {
		child.Context = cmd.Context
	}
	cmd.SetArgs([]string{"initminer", "article,image", "Lorem Ipsum", "Lorem ipsum dolor sit amet"})
	priv_account := &wallet.PrivAccount{
		Account: wallet.Account{
			Name:   "initminer",
			PubKey: "COS5JVLLcTPhq4Unr194JzWPDNSYGoMcam8yxnsjgRVo3Nb7ioyFW",
		},
		PrivKey: "4DjYx2KAGh1NP3dai7MZTLUBMMhMBPmwouKE8jhVSESywccpVZ",
	}
	mywallet.EXPECT().GetUnlockedAccount("initminer").Return(priv_account, true)
	resp := &grpcpb.BroadcastTrxResponse{Status: 1, Msg: "success"}
	client.EXPECT().BroadcastTrx(gomock.Any(), gomock.Any()).Return(resp, nil).Do(func(context interface{}, req *grpcpb.BroadcastTrxRequest) {
		op := req.Transaction.Trx.Operations[0]
		post_op := op.GetOp6()
		myassert.Equal(post_op.Title, "Lorem Ipsum")
		myassert.Equal(post_op.Content, "Lorem ipsum dolor sit amet")
	})
	_, err := cmd.ExecuteC()
	if err != nil {
		t.Error(err)
	}
}

func TestPostWithBeneficiaries(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)
	mywallet := mock_wallet.NewMockWallet(ctrl)
	myassert := assert.New(t)
	passwordReader := mock_utils.NewMockPasswordReader(ctrl)
	cmd := PostCmd()
	cmd.SetContext("wallet", mywallet)
	cmd.SetContext("rpcclient", client)
	cmd.SetContext("preader", passwordReader)
	for _, child := range cmd.Commands() {
		child.Context = cmd.Context
	}
	cmd.SetArgs([]string{"initminer", "article,image", "Lorem Ipsum", "Lorem ipsum dolor sit amet",
		"-b", "Alice=5,Bob=5"})
	priv_account := &wallet.PrivAccount{
		Account: wallet.Account{
			Name:   "initminer",
			PubKey: "COS5JVLLcTPhq4Unr194JzWPDNSYGoMcam8yxnsjgRVo3Nb7ioyFW",
		},
		PrivKey: "4DjYx2KAGh1NP3dai7MZTLUBMMhMBPmwouKE8jhVSESywccpVZ",
	}
	mywallet.EXPECT().GetUnlockedAccount("initminer").Return(priv_account, true)
	resp := &grpcpb.BroadcastTrxResponse{Status: 1, Msg: "success"}
	client.EXPECT().BroadcastTrx(gomock.Any(), gomock.Any()).Return(resp, nil).Do(func(context interface{}, req *grpcpb.BroadcastTrxRequest) {
		op := req.Transaction.Trx.Operations[0]
		post_op := op.GetOp6()
		if post_op.Beneficiaries[0].Name.Value == "Alice" {
			myassert.Equal(post_op.Beneficiaries[1].Name.Value, "Bob")
			myassert.Equal(post_op.Beneficiaries[1].Weight, uint32(5))
		} else {
			myassert.Equal(post_op.Beneficiaries[1].Name.Value, "Alice")
			myassert.Equal(post_op.Beneficiaries[1].Weight, uint32(5))
		}
	})
	_, err := cmd.ExecuteC()
	if err != nil {
		t.Error(err)
	}
}
