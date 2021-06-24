package service

import (
	"spam.com/enty"	
	"spam.com/repository"
)


type VideoService interface {
	Save(enty.Video) enty.Video
	FindAll() []enty.Video
}

type videoService struct {
	//videos []enty.Video
	videoRepository repository.VideoRepository
}

func New(repo  repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Save(video enty.Video) enty.Video {
	service.videoRepository.Save(video)
	return video
}
func (service *videoService) FindAll() []enty.Video {
	return service.videoRepository.FindAll()
}
