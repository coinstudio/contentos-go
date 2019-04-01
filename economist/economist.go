package economist

import (
	"github.com/coschain/contentos-go/app/table"
	"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/iservices"
	"github.com/coschain/contentos-go/prototype"
	"github.com/pkg/errors"
	"math"
)

func Min(x, y uint64) uint64 {
	if x < y {
		return x
	} else {
		return y
	}
}

type Economist struct {
	db       iservices.IDatabaseService
	singleId *int32
}

func New(db iservices.IDatabaseService, singleId *int32) *Economist {
	return &Economist{db: db, singleId: singleId}
}

func (e *Economist) GetProps() (*prototype.DynamicProperties, error) {
	dgpWrap := table.NewSoGlobalWrap(e.db, e.singleId)
	if !dgpWrap.CheckExist() {
		return nil, errors.New("the mainkey is already exist")
	}
	return dgpWrap.GetProps(), nil
}

func (e *Economist) GetRewardsKeeper() (*prototype.InternalRewardsKeeper, error) {
	keeperWrap := table.NewSoRewardsKeeperWrap(e.db, e.singleId)
	if !keeperWrap.CheckExist() {
		return nil, errors.New("Economist access rewards keeper error")
	}
	return keeperWrap.GetKeeper(), nil
}

func (e *Economist) updateRewardsKeeper(rewardKeeper *prototype.InternalRewardsKeeper) {
	keeper := table.NewSoRewardsKeeperWrap(e.db, e.singleId)
	success := keeper.MdKeeper(rewardKeeper)
	if !success {
		panic("flush rewards into db error")
	}
}

func (e *Economist) modifyGlobalDynamicData(f func(props *prototype.DynamicProperties)) {
	dgpWrap := table.NewSoGlobalWrap(e.db, e.singleId)
	props := dgpWrap.GetProps()

	f(props)

	success := dgpWrap.MdProps(props)
	if !success {
		panic("flush globalDynamic into db error")
	}
}

func (e *Economist) BaseBudget(ith uint32) uint64 {
	if ith > 12 {
		return 0
	}
	return uint64(ith) * uint64(constants.TotalRelease) * uint64(448) / 1000 / 100 * constants.BaseRate

}

// whitepaper p16
func (e *Economist) InitialBonus(ith uint32) uint64 {
	if ith > 5 {
		return 0
	}
	switch ith {
	case 0:
		return 0
	case 1:
		return 169580103 * constants.BaseRate
	case 2:
		return 106993132 * constants.BaseRate
	case 3:
		return 84790051 * constants.BaseRate
	case 4:
		return 73034175 * constants.BaseRate
	case 5:
		return 65602539 * constants.BaseRate
	}
	return 0
}

func (e *Economist) CalculateBudget(ith uint32) uint64 {
	return e.BaseBudget(ith) + e.InitialBonus(ith)
}

func (e *Economist) CalculatePerBlockBudget(annalBudget uint64) uint64 {
	return annalBudget / (86400 / 3 * 365)
}

func (e *Economist) Mint() {
	//blockCurrent := constants.PerBlockCurrent
	globalProps, err := e.GetProps()
	ith := globalProps.GetIthYear()
	annualBudget := e.CalculateBudget(ith)
	// new year arrived
	if globalProps.GetAnnualBudget().Value != annualBudget {
		e.modifyGlobalDynamicData(func(props *prototype.DynamicProperties) {
			props.AnnualBudget.Value = annualBudget
			props.AnnualMinted.Value = 0
		})
	}
	blockCurrent := e.CalculatePerBlockBudget(annualBudget)
	// prevent deficit
	if globalProps.GetAnnualBudget().Value < (globalProps.GetAnnualMinted().Value + blockCurrent) {
		blockCurrent = globalProps.GetAnnualBudget().Value - globalProps.GetAnnualMinted().Value
		// it is impossible
		if blockCurrent < 0 {
			blockCurrent = 0
		}
		// time to update year
		e.modifyGlobalDynamicData(func(props *prototype.DynamicProperties) {
			props.IthYear = props.IthYear + 1
		})
	}

	authorReward := blockCurrent * constants.RewardRateAuthor / constants.PERCENT
	replyReward := blockCurrent * constants.RewardRateReply / constants.PERCENT
	bpReward := blockCurrent * constants.RewardRateBP / constants.PERCENT
	reportReward := blockCurrent - authorReward - replyReward - bpReward

	if err != nil {
		panic("Mint failed when getprops")
	}
	keeper, err := e.GetRewardsKeeper()

	_ = bpReward
	currentBp := globalProps.GetCurrentWitness().Value
	rewards := keeper.GetRewards()
	if vest, ok := rewards[currentBp]; !ok {
		rewards[currentBp] = &prototype.Vest{Value: uint64(bpReward)}
	} else {
		vest.Value += uint64(bpReward)
	}
	keeper.Rewards = rewards
	e.updateRewardsKeeper(keeper)

	e.modifyGlobalDynamicData(func(props *prototype.DynamicProperties) {
		props.PostRewards.Value += uint64(authorReward)
		props.ReplyRewards.Value += uint64(replyReward)
		props.ReportRewards.Value += uint64(reportReward)
		props.AnnualMinted.Value += blockCurrent
	})

}

// Should be claiming or direct modify the balance?
func (e *Economist) Do() {
	e.decayGlobalVotePower()
	globalProps, err := e.GetProps()
	if err != nil {
		panic("economist do failed when get props")
	}
	if globalProps.HeadBlockNumber%constants.ReportCashout == 0 {
		reportRewards := globalProps.ReportRewards.Value
		postRewards := reportRewards / 2
		replyRewards := reportRewards - postRewards
		e.modifyGlobalDynamicData(func(props *prototype.DynamicProperties) {
			props.ReportRewards.Value = 0
			props.PostRewards.Value += postRewards
			props.ReplyRewards.Value += replyRewards
		})
	}
	timestamp := globalProps.Time.UtcSeconds
	iterator := table.NewPostCashoutTimeWrap(e.db)
	var pids []*uint64
	err = iterator.ForEachByOrder(nil, &prototype.TimePointSec{UtcSeconds: timestamp}, nil, nil, func(mVal *uint64, sVal *prototype.TimePointSec, idx uint32) bool {
		pids = append(pids, mVal)
		return true
	})
	if err != nil {
		panic("economist do failed when iterator")
	}
	var posts []*table.SoPostWrap
	var replies []*table.SoPostWrap

	var vpAccumulator uint64 = 0
	for _, pid := range pids {
		post := table.NewSoPostWrap(e.db, pid)
		if post.GetParentId() == 0 {
			posts = append(posts, post)
			vpAccumulator += post.GetWeightedVp()
		} else {
			replies = append(replies, post)
			vpAccumulator += uint64(math.Ceil(math.Sqrt(float64(post.GetWeightedVp()))))
		}
	}
	e.modifyGlobalDynamicData(func(props *prototype.DynamicProperties) {
		props.WeightedVps += vpAccumulator
	})
	keeper, err := e.GetRewardsKeeper()
	if len(posts) > 0 {
		e.postCashout(keeper, posts)
	}

	if err != nil {
		panic("economist do failed when get reward keeper")
	}

	if len(replies) > 0 {
		e.replyCashout(keeper, replies)
	}
	e.updateRewardsKeeper(keeper)
}

func (e *Economist) decayGlobalVotePower() {
	e.modifyGlobalDynamicData(func(props *prototype.DynamicProperties) {
		props.WeightedVps -= props.WeightedVps * constants.BlockInterval / constants.VpDecayTime
	})
}

func (e *Economist) postCashout(rewardKeeper *prototype.InternalRewardsKeeper, posts []*table.SoPostWrap) {
	globalProps, err := e.GetProps()
	if err != nil {
		panic("post cashout get props failed")
	}

	var vpAccumulator uint64 = 0
	for _, post := range posts {
		vpAccumulator += post.GetWeightedVp()
	}
	var blockReward uint64 = 0
	if globalProps.WeightedVps > 0 {
		blockReward = vpAccumulator * globalProps.PostRewards.Value / globalProps.WeightedVps
	}
	innerRewards := rewardKeeper.Rewards
	for _, post := range posts {
		author := post.GetAuthor().Value
		var reward uint64 = 0
		// divide zero exception
		if vpAccumulator > 0 {
			reward = post.GetWeightedVp() * blockReward / vpAccumulator
		}
		if vest, ok := innerRewards[author]; !ok {
			innerRewards[author] = &prototype.Vest{Value: reward}
		} else {
			vest.Value += reward
		}
		post.MdCashoutTime(&prototype.TimePointSec{UtcSeconds: math.MaxUint32})
	}
}

// use same algorithm to simplify
func (e *Economist) replyCashout(rewardKeeper *prototype.InternalRewardsKeeper, replies []*table.SoPostWrap) {
	globalProps, err := e.GetProps()
	if err != nil {
		panic("reply cashout get props failed")
	}
	var vpAccumulator uint64 = 0
	for _, reply := range replies {
		vpAccumulator += uint64(math.Ceil(math.Sqrt(float64(reply.GetWeightedVp()))))
	}
	e.modifyGlobalDynamicData(func(props *prototype.DynamicProperties) {
		props.WeightedVps += vpAccumulator
	})
	var blockReward uint64 = 0
	if globalProps.WeightedVps > 0 {
		blockReward = vpAccumulator * globalProps.PostRewards.Value / globalProps.WeightedVps
	}
	innerRewards := rewardKeeper.Rewards
	for _, reply := range replies {
		author := reply.GetAuthor().Value
		var reward uint64 = 0
		// divide zero exception
		if vpAccumulator > 0 {
			reward = reply.GetWeightedVp() * blockReward / vpAccumulator
		}
		if vest, ok := rewardKeeper.Rewards[author]; !ok {
			innerRewards[author] = &prototype.Vest{Value: reward}
		} else {
			vest.Value += reward
		}
		reply.MdCashoutTime(&prototype.TimePointSec{UtcSeconds: math.MaxUint32})
	}
}

func (e *Economist) PowerDown() {
	globalProps, err := e.GetProps()
	if err != nil {
		panic("economist do failed when get props")
	}
	timestamp := globalProps.Time.UtcSeconds
	iterator := table.NewAccountNextPowerdownTimeWrap(e.db)
	var accountNames []*prototype.AccountName
	err = iterator.ForEachByOrder(nil, &prototype.TimePointSec{UtcSeconds: timestamp}, nil, nil, func(mVal *prototype.AccountName, sVal *prototype.TimePointSec, idx uint32) bool {
		accountNames = append(accountNames, mVal)
		return true
	})
	var powerdown_quota uint64 = 0
	for _, accountName := range accountNames {
		accountWrap := table.NewSoAccountWrap(e.db, accountName)
		if accountWrap.GetToPowerdown().Value-accountWrap.GetHasPowerdown().Value < accountWrap.GetEachPowerdownRate().Value {
			powerdown_quota = Min(accountWrap.GetVestingShares().Value, accountWrap.GetToPowerdown().Value-accountWrap.GetHasPowerdown().Value)
		} else {
			powerdown_quota = Min(accountWrap.GetVestingShares().Value, accountWrap.GetEachPowerdownRate().Value)
		}
		vesting := accountWrap.GetVestingShares().Value - powerdown_quota
		balance := accountWrap.GetBalance().Value + powerdown_quota
		hasPowerDown := accountWrap.GetHasPowerdown().Value + powerdown_quota

		accountWrap.MdVestingShares(&prototype.Vest{Value: vesting})
		accountWrap.MdBalance(&prototype.Coin{Value: balance})
		accountWrap.MdHasPowerdown(&prototype.Vest{Value: hasPowerDown})
		if accountWrap.GetHasPowerdown().Value >= accountWrap.GetToPowerdown().Value || accountWrap.GetVestingShares().Value == 0 {
			accountWrap.MdEachPowerdownRate(&prototype.Vest{Value: 0})
			accountWrap.MdNextPowerdownTime(&prototype.TimePointSec{UtcSeconds: math.MaxUint32})
		} else {
			accountWrap.MdNextPowerdownTime(&prototype.TimePointSec{UtcSeconds: timestamp + constants.POWER_DOWN_INTERVAL})
		}
	}
}
