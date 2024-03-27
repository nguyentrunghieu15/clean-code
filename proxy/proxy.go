package main

import "fmt"

type Video struct{}

type ThisPartyYoutube interface {
	ListVideo() []Video
	GetVideoInfo(int) Video
	DownloadVideo(int)
}

type ThirdPartyYoutubeClass struct{}

// DownloadVideo implements ThisPartyYoutube.
func (s *ThirdPartyYoutubeClass) DownloadVideo(id int) {
	fmt.Println("Downloadding")
}

// GetVideoInfo implements ThisPartyYoutubde.
func (s *ThirdPartyYoutubeClass) GetVideoInfo(id int) Video {
	fmt.Println("Getting")
	return Video{}
}

// ListVideo implements ThisPartyYoutube.
func (s *ThirdPartyYoutubeClass) ListVideo() []Video {
	return make([]Video, 0)
}

// CachedYoutube là 1 lớp proxy của ThirdPartyYoutubeClass
// Chỉ lấy video khi video không có trong cache
// Giảm load khi có nhiều yêu cầu lấy danh sách video
type CachedYoutube struct {
	Service         ThirdPartyYoutubeClass
	ListVideoCached []Video
	VideoCached     Video
	NeedReset       bool
}

func (s *CachedYoutube) IsExistedVideo(id int) bool {
	// for range ListVideoCached to check exists
	return true
}

// DownloadVideo implements ThisPartyYoutube.
func (s *CachedYoutube) DownloadVideo(id int) {
	if s.NeedReset || !s.IsExistedVideo(id) {
		s.Service.DownloadVideo(id)
		return
	}
	fmt.Println("Dowload video from cache")
}

// GetVideoInfo implements ThisPartyYoutube.
func (s *CachedYoutube) GetVideoInfo(id int) Video {
	if s.NeedReset || !s.IsExistedVideo(id) {
		return s.Service.GetVideoInfo(id)
	}
	fmt.Println("Getting video from cache")
	return Video{}
}

// ListVideo implements ThisPartyYoutube.
func (s *CachedYoutube) ListVideo() []Video {
	if s.ListVideoCached == nil || s.NeedReset {
		s.ListVideoCached = s.Service.ListVideo()
	}
	return s.ListVideoCached
}
