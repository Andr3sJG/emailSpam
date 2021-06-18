package service

import "spam.com/enty"



type VideoService interface {
	Save(enty.Video) enty.Video
	FindAll() [] enty.Video
}

type videoService struct{
	videos []enty.Video
}

func New() VideoService{
	return &videoService{}
}

func (service *videoService)Save(video enty.Video) enty.Video {
	service.videos = append(service.videos, video)
	return video
}
func (service *videoService)FindAll() []enty.Video{
	return service.videos
}