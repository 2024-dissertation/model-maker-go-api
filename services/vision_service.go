package services

type VisionService interface {
	AnalyseImage(imagePath string) (string, error)
}
