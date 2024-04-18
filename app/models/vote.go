package models

import "fmt"

// GetVote 查对应userID的所有投票
func GetVote(voteId int32) {
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
