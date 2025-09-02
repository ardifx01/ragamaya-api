package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type Pattern string

const (
	BALI        Pattern = "Bali"
	BETAWI      Pattern = "Betawi"
	CELUP       Pattern = "Celup"
	CENDRAWASIH Pattern = "Cendrawasih"
	CEPLOK      Pattern = "Ceplok"
	CIAMIS      Pattern = "Ciamis"
	GARUTAN     Pattern = "Garutan"
	GENTONGAN   Pattern = "Gentongan"
	KAWUNG      Pattern = "Kawung"
	KERATON     Pattern = "Keraton"
	LASEM       Pattern = "Lasem"
	MEGAMENDUNG Pattern = "Megamendung"
	PARANG      Pattern = "Parang"
	PEKALONGAN  Pattern = "Pekalongan"
	PRIANGAN    Pattern = "Priangan"
	SEKAR       Pattern = "Sekar"
	SIDOLUHUR   Pattern = "Sidoluhur"
	SIDOMUKTI   Pattern = "Sidomukti"
	SOGAN       Pattern = "Sogan"
	TAMBAL      Pattern = "Tambal"
)

type MatchLevel string

const (
	High   MatchLevel = "high"
	Medium MatchLevel = "medium"
	Low    MatchLevel = "low"
)

type Alternative struct {
	Pattern Pattern    `json:"pattern"`
	Code    string     `json:"code"`
	Score   float32    `json:"score"`
	Match   MatchLevel `json:"match"`
}

type MLRes struct {
	Detected     Pattern       `json:"detected"`
	Alternatives []Alternative `json:"alternatives"`
	Error        *string       `json:"error,omitempty"`
}

type PredictRes struct {
	Pattern     string     `json:"pattern"`
	Origin      string     `json:"origin"`
	Description string     `json:"description"`
	History     string     `json:"history"`
	Score       float32    `json:"score"`
	Match       MatchLevel `json:"match"`

	Alternative []Alternative `json:"alternative"`
}
