package main

import "fmt"

type VideoDownload interface {
	videoDownload(string) string
}

type RealVideoDownloader struct{}

func (v *RealVideoDownloader) videoDownload(videoURL string) string {
	fmt.Printf("%s video download started ...\n", videoURL)
	return fmt.Sprintf("Video content from %s: ", videoURL)
}

type CacheVideoDownloader struct {
	realVideo *RealVideoDownloader
	cache     map[string]string
}

func (c *CacheVideoDownloader) videoDownload(videoURL string) string {
	// if videoURL in cache, return the cache value
	if val, ok := c.cache[videoURL]; ok {
		fmt.Println("Getting the video from cache")
		return val
	}
	fmt.Println("Getting the actual video")
	realVid := c.realVideo.videoDownload(videoURL)
	c.cache[videoURL] = realVid
	return realVid
}

func main() {
	cacheVideo := &CacheVideoDownloader{
		realVideo: &RealVideoDownloader{},
		cache:     make(map[string]string),
	}
	cacheVideo.videoDownload("lld/proxy-pattern")
	cacheVideo.videoDownload("lld/proxy-pattern")
	cacheVideo.videoDownload("lld/composite-pattern")

	cacheVideo2 := &CacheVideoDownloader{
		realVideo: &RealVideoDownloader{},
		cache:     make(map[string]string),
	}
	cacheVideo2.videoDownload("lld/proxy-pattern")
}