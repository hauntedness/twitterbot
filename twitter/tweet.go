package twitter

type Tweet struct {
	NumResults      int64         `json:"num_results"`
	Users           []User        `json:"users"`
	Topics          []Topic       `json:"topics"`
	Events          []interface{} `json:"events"`
	Lists           []interface{} `json:"lists"`
	OrderedSections []string      `json:"ordered_sections"`
	Oneclick        []interface{} `json:"oneclick"`
	Hashtags        []interface{} `json:"hashtags"`
	CompletedIn     float64       `json:"completed_in"`
	Query           string        `json:"query"`
}

type Topic struct {
	Topic        string  `json:"topic"`
	RoundedScore int64   `json:"rounded_score"`
	Tokens       []Token `json:"tokens"`
	Inline       bool    `json:"inline"`
}

type Token struct {
	Token string `json:"token"`
}

type User struct {
	ID                   int64         `json:"id"`
	IDStr                string        `json:"id_str"`
	Verified             bool          `json:"verified"`
	IsDmAble             bool          `json:"is_dm_able"`
	IsBlocked            bool          `json:"is_blocked"`
	Name                 string        `json:"name"`
	ScreenName           string        `json:"screen_name"`
	ProfileImageURL      string        `json:"profile_image_url"`
	ProfileImageURLHTTPS string        `json:"profile_image_url_https"`
	Location             string        `json:"location"`
	IsProtected          bool          `json:"is_protected"`
	RoundedScore         int64         `json:"rounded_score"`
	SocialProof          int64         `json:"social_proof"`
	ConnectingUserCount  int64         `json:"connecting_user_count"`
	ConnectingUserIDS    []interface{} `json:"connecting_user_ids"`
	SocialProofsOrdered  []interface{} `json:"social_proofs_ordered"`
	SocialContext        SocialContext `json:"social_context"`
	Tokens               []Token       `json:"tokens"`
	Inline               bool          `json:"inline"`
	ResultContext        ResultContext `json:"result_context"`
}

type ResultContext struct {
	DisplayString *string       `json:"display_string,omitempty"`
	Types         []TypeElement `json:"types"`
}

type TypeElement struct {
	Type TypeEnum `json:"type"`
}

type SocialContext struct {
	Following  bool `json:"following"`
	FollowedBy bool `json:"followed_by"`
}

type TypeEnum string

const (
	Bio            TypeEnum = "bio"
	Location       TypeEnum = "location"
	NumOfFollowers TypeEnum = "num_of_followers"
)
