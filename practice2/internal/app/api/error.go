package api

const (
	ErrorInvalidParams = "4000"
	ErrorNotFoundUrl   = "4040"
	ErrorNoExist       = "4041"
	ErrorNoMethod      = "4050"

	ErrorDb = "5001"

	Error = "9999"
)

var Messages = map[string]string{
	Error: "오류가 발생하였습니다",

	ErrorInvalidParams: "유효하지 않은 파라미터입니다",
	ErrorNoExist:       "데이터가 없습니다",
	ErrorNotFoundUrl:   "유효하지 않은 호출(URL)입니다",
	ErrorNoMethod:      "유효하지 않은 호출(method)입니다",

	ErrorDb: "서버(DB)에서 오류가 발생하였습니다. 잠시 후 다시 시도하여 주세요",
}
