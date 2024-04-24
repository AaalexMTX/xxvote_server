package models

import (
	"fmt"
	"gorm.io/gorm"
	"sync"
	"time"
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

// GetVotes 查询所有vote选项
func GetVotes() []Vote {
	//用make 别直接建立 == nil
	//var votes []VoteOption
	votes := make([]Vote, 0)
	if err := MysqlConnect.Raw("select * from vote where status > 0").Scan(&votes).Error; err != nil {
		fmt.Printf("[models_GetVote]err:%s\n", err.Error())
	}
	return votes
}

// GetVoteWithOptions  将vote_id 对应的所有项目返回
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

// DoVote 谁 在什么项目  投了哪些选项
func DoVote(userID, voteID int, options []int) bool {
	// 还没用apiFox验证
	err := MysqlConnect.Transaction(func(tx *gorm.DB) error {
		// 1. 查用户是否存在（合法）
		var user User
		if err := tx.Raw("select * from user where id = ?", userID).Scan(&user).Error; err != nil {
			fmt.Printf("[DoVote] user not exist -- err : %s", err.Error())
			return err
		}
		// 2.1 检查投票选项是否合法
		var vote Vote
		if err := tx.Raw("select * from vote where id = ?", userID).Scan(&vote).Error; err != nil || vote.Status == 0 {
			fmt.Printf("[DoVote] vote not exist || timeOut -- err : %s", err.Error())
			return err
		}

		for _, optionId := range options {
			// 2.2 更新vote_option 对应的count
			if err := tx.Exec("update vote_option set count = count + 1 where  id = ? ", optionId).Error; err != nil {
				fmt.Printf("[DoVote] vote_opt_count++ error -- err : %s", err.Error())
				return err
			}
			// 3. 创建vote_opt_user的一条记录
			var voteOptUser = VoteOptionUser{
				UserId:       userID,
				VoteId:       voteID,
				VoteOptionId: optionId,
				CreateTime:   time.Now(),
				UpdateTime:   time.Now(),
			}
			if err := tx.Create(&voteOptUser).Error; err != nil {
				fmt.Printf("[DoVote] vote_opt_User create a Record error -- err : %s", err.Error())
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false
	}
	return true
}
