package model

import (
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

var (
	//插入时omit的字段
	BaseInsertOmitField = []string{"`id`", "`create_at`", "`update_at`", "`is_deleted`", "`delete_at`"}
	//更新时omit的字段
	BaseUpdateOmitField = []string{"`id`", "`uuid`", "`create_at`", "`update_at`", "`is_deleted`", "`delete_at`"}
)

type LuckerDetails struct {
	//抽奖者的信息
	Lucker Lucker
	//奖品池的信息
	PrizesPool PrizesPool
	//活动信息
	Activity Activity
	//奖品信息
	Prize Prize
	//奖品级别
	PrizeGrade PrizeGrade
}

type LuckerDetail struct {
	//幸运者信息
	LuckerId     int64     `db:"lucker_id"`        // 表主键
	LuckerUuid   string    `db:"lucker_uuid"`      // 唯一识别码
	UserId       int64     `db:"lucker_user_id"`   // 中奖用户id
	LuckerStatus int64     `db:"lucker_status"`    // 状态: -1放弃领奖，0未领奖，1已领奖
	LuckDate     time.Time `db:"lucker_luck_date"` // 抽奖时间

	//奖品池信息
	PrizesPoolId        int64  `db:"prizes_pool_id"`         // 表主键
	PrizesPoolUuid      string `db:"prizes_pool_uuid"`       // 唯一识别码
	PrizesPoolLuckValue int64  `db:"prizes_pool_luck_value"` // 幸运值
	PrizesPoolStatus    int64  `db:"prizes_pool_status"`     // 是否可用，0不可用，1可用
	PrizesPoolGradeId   int64  `db:"prizes_pool_grade_id"`   // 等级名称，如：一等奖
	PrizesPoolNumber    int64  `db:"prizes_pool_number"`     // 奖品总数量
	PrizesPoolRemainder int64  `db:"prizes_pool_remainder"`  // 奖品剩余数量

	//活动信息
	ActivityId        int64     `db:"activity_id"`         // 活动id
	ActivityUuid      string    `db:"activity_uuid"`       // 唯一识别码
	ActivityName      string    `db:"activity_name"`       // 活动名称
	ActivityProjectId int64     `db:"activity_project_id"` // 项目id
	ActivityStatus    int64     `db:"activity_status"`     // 活动状态，是否在举办，默认关闭，需手动开启
	ActivityPattern   int64     `db:"activity_pattern"`    // 方案: 0私密的，1公开的
	ActivityMaxCount  int64     `db:"activity_max_count"`  // 最大抽奖次数：-1表示无限制，0表示不可抽奖
	ActivityMaxPrizes int64     `db:"activity_max_prizes"` // 最大抽奖次数：-1表示无限制，0表示不可抽奖
	ActivityStartTime time.Time `db:"activity_start_time"` // 活动开始时间
	ActivityEndTime   time.Time `db:"activity_end_time"`   // 活动结束时间

	//奖品信息
	PrizeId     int64   `db:"prize_id"`     // 奖品id
	PrizeUuid   string  `db:"prize_uuid"`   // 唯一识别码
	PrizeName   string  `db:"prize_name"`   // 奖品名称
	PrizeImage  string  `db:"prize_image"`  // 奖品图片
	PrizePrice  float64 `db:"prize_price"`  // 奖品价格
	PrizeStatus int64   `db:"prize_status"` // 是否可用

	//中奖级别
	GradeId   int64  `db:"prize_grade_id"`     // 表主键
	GradeName string `db:"prize_grade_name"`   // 级别名称
	IsWin     int64  `db:"prize_grade_is_win"` // 是否中奖
}
