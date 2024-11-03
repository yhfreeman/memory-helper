package main

// type Event struct {
// 	Type string
// 	Title string
// 	Description string
// 	Addtime int64
// 	Tiptime int64
// 	Owner string
// }

type StudyItem struct {
    ID          string `yaml:"id"`
    Type        string `yaml:"type"`
    Owner       string `yaml:"owner"`
    Title       string `yaml:"title"`
    Description string `yaml:"description"`
    Addtime     int64  `yaml:"addtime"`
}

type Event struct {
    StudyItemID string `yaml:"study_item_id"`
    Tiptime     int64  `yaml:"tiptime"`
    IsCompleted bool   `yaml:"is_completed"`
}

type ServerConfigs struct {
	addr string
	static string
	dbFile string
}
