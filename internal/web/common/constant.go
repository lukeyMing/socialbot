package common

const (
	DefaultPage = 15
	GalleryPage = 30
	SortNew    = 0
	sortHot   = 1
	// account type
	AccountTypeEmail = "email"
	// jwt
	JwtAuthUserKey = "userId"
	SuperAdmin     = 3

	// storage
	StorageMediaDir      = "media"
	StorageAvatarDir      = "avatar"
	DefaultStorageSource = 0
	SourceTypeImage      = 0
	SourceTypeVideo      = 1

	// crawler
	CrawlerListPageNum   = 30
	CrawlerItemNew       = 0  // 新入库
	CrawlerItemPublished = 10 // 已发布
	CrawlerImgsLimit     = 10 // 发布图片数量最大长度限制

	// media
	MediaStatusPublished      = 10
	MediaTypePromotionProduct = 4
	MediaTypeArticle          = 3
	MediaTypeSocial           = 2
)
