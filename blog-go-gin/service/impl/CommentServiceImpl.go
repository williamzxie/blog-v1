package impl

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"blog-go-gin/models/model"
	"blog-go-gin/models/page"
	"strconv"
	"sync"
)

type CommentServiceImpl struct {
	wg sync.WaitGroup
}

func (c *CommentServiceImpl) GetComments(articleId int, ipage *page.IPage) (*pb.CommentInfo, error) {
	condition := "article_id = ? and parent_id = ? and is_delete = ?"
	//查询文章评论量
	commentCount, err := model.GetCommentsCountByCondition(condition, articleId, 0, false)
	if err != nil {
		return nil, err
	}
	if commentCount == 0 {
		return nil, nil
	}
	//分页查询评论集合
	comments, err := model.GetCommentsAndUserInfo(ipage, condition, articleId, 0, false)
	if err != nil {
		return nil, err
	}
	var commentsSlice []*pb.Comment
	for _, comment := range comments {
		commentsSlice = append(commentsSlice, &pb.Comment{
			Id:             int32(comment.ID),
			UserId:         int32(comment.UserID),
			Nickname:       comment.Nickname,
			Avatar:         comment.Avatar,
			WebSite:        comment.WebSite,
			CommentContent: comment.CommentContent,
			CreateTime:     comment.CreateTime,
		})
	}
	//查询评论点赞数据
	likeCountMap, err := common.RedisUtil.HashGetAll(common.CommentLikeCount)
	if err != nil {
		return nil, err
	}
	logging.Logger.Debug(likeCountMap)
	//封装评论点赞量
	var commentIdSlice []int64
	for _, comment := range commentsSlice {
		commentIdSlice = append(commentIdSlice, int64(comment.Id))
		logging.Logger.Debug(comment.Id)
		likeCount, _ := strconv.Atoi(likeCountMap[string(comment.Id)])
		comment.LikeCount = uint32(likeCount)
	}
	//根据评论id集合查询回复数据
	replies, err := model.GetReplies(commentIdSlice)
	if err != nil {
		return nil, err
	}
	var repliesSlice []*pb.Reply
	for _, reply := range replies {
		repliesSlice = append(repliesSlice, &pb.Reply{
			Id:             int32(reply.ID),
			ParentId:       int32(reply.ParentID),
			UserId:         int32(reply.UserID),
			Nickname:       reply.Nickname,
			Avatar:         reply.Avatar,
			WebSite:        reply.WebSite,
			ReplyId:        int32(reply.ReplyID),
			CommentContent: reply.CommentContent,
			CreateTime:     reply.CreateTime,
		})
	}
	//封装回复点赞量
	for _, reply := range repliesSlice {
		likeCount, _ := strconv.Atoi(likeCountMap[string(reply.Id)])
		reply.LikeCount = uint32(likeCount)
	}
	//根据评论id分组回复数据
	replyMap := make(map[int32][]*pb.Reply)
	for _, reply := range repliesSlice {
		repliesByGroup := append(replyMap[reply.ParentId], reply)
		replyMap[reply.ParentId] = repliesByGroup
	}
	//根据评论id查询回复量
	replyCountList, err := model.GetReplyCountByCommentId(commentIdSlice)
	replyCountMap := make(map[int32]int32)
	for _, reply := range replyCountList {
		replyCountMap[int32(reply.ParentID)] = int32(reply.ReplyCount)
	}
	//将分页回复数据和回复量封装进对应的评论
	for _, comment := range commentsSlice {
		comment.ReplyList = replyMap[comment.Id]
		comment.ReplyCount = uint32(replyCountMap[comment.Id])
	}
	logging.Logger.Debug(commentsSlice)

	return &pb.CommentInfo{
		CommentList: commentsSlice,
		Count:       uint32(commentCount),
	}, nil
}
