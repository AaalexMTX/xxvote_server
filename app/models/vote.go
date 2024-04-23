package models

import (
	"fmt"
	"sync"
)

// GetVote 查对应userID的所有投票
func GetVote(voteId int) {
	var vote []Vote
	if err := MysqlConnect.Raw("select * from vote where id = ?", voteId).Scan(&vote).Error; err != nil {
		fmt.Printf("[models_GetVote]err:%s\n", err.Error())
	}
	var opt []VoteOption
	if err := MysqlConnect.Raw("select * from vote_option where vote_id = ?", voteId).Scan(&opt).Error; err != nil {
		fmt.Printf("[models_GetVote]err:%s\n", err.Error())
	}
}

func GetVotes() []Vote {
	//用make 别直接建立 == nil
	//var votes []VoteOption
	votes := make([]Vote, 0)
	if err := MysqlConnect.Raw("select * from vote where status > 0").Scan(&votes).Error; err != nil {
		fmt.Printf("[models_GetVote]err:%s\n", err.Error())
	}
	return votes
}

func GetVoteWithOptions(VoteId int) VoteWithOptions {
	//vote 和 option的联合查询
	wg := sync.WaitGroup{}
	var vote Vote
	opt := make([]VoteOption, 0)
	//查vote - by id
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := MysqlConnect.Raw("select * from vote where id = ? limit 1", VoteId).Scan(&vote).Error; err != nil {
			fmt.Printf("[GetVoteWithOptions] vote query failure -- err:%s", err.Error())
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := MysqlConnect.Raw("select * from vote_option where vote_id = ?", VoteId).Scan(&opt).Error; err != nil {
			fmt.Printf("[GetVoteWithOptions] vote query failure -- err:%s", err.Error())
			return
		}
	}()

	wg.Wait()
	return VoteWithOptions{
		Vote:   vote,
		Option: opt,
	}
}
