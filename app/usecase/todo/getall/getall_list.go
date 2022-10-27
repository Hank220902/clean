package getall

import model "clean/app/domain/model/todo"

func convertGetAllListToOutput(in []*model.Get) []*GetAllOutput {
	ans := make([]*GetAllOutput, len(in))
	for i, v := range in {
		ans[i] = converTodo(v)
	}
	return ans

}

func converTodo(in *model.Get) *GetAllOutput {
	if in == nil {
		return new(GetAllOutput)
	}
	return &GetAllOutput{
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
