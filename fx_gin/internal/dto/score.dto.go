package dto

import "small-run/fx_gin/internal/models/entities"

type ScoreReq struct {
	StudentID int `json:"student_id"`
	SubjectID int `json:"subject_id"`
}

type ScoreRes struct {
	StudentName string `json:"student_name"`
	SubjectName string `json:"subject_name"`
	Score int `json:"score"`
}

func BeRes(stu *entities.StudentEntity, sub *entities.SubjectEntity, score *entities.ScoreEntity) ScoreRes {
	return ScoreRes{StudentName: stu.Name, SubjectName: sub.Name, Score: score.Score}
}
