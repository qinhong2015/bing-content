package Service

type feedUrlImporterService struct {

}

var feedUrlImporterInstance *feedUrlImporterService

func GetFeedUrlImporterService() *feedUrlImporterService {
	if feedUrlImporterInstance == nil {
		feedUrlImporterInstance = new(feedUrlImporterService)
	}
	return feedUrlImporterInstance
}
