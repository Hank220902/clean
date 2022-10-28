package getfilter

import model "clean/app/domain/model/todo"

func convertGetFilterListToOutput(in []*model.Get) []*GetFilterOutput {
	ans := make([]*GetFilterOutput, len(in))
	for i, v := range in {
		ans[i] = converTodo(v)
	}
	return ans

}

func converTodo(in *model.Get) *GetFilterOutput {
	if in == nil {
		return new(GetFilterOutput)
	}
	return &GetFilterOutput{
		ID:                in.ID,
		Matter:            in.Matter,
		EndTime:           in.EndTime.String(),
		FinishedCondition: in.FinishedCondition,
		Status:            in.Status,
		Email:             in.Email,
		Note:              in.Note,
		CreateTime:        in.CreateTime.String(),
		UpdateTime:        in.UpdateTime.String(),
	}
}
