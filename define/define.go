package define

import "github.com/golang-jwt/jwt/v4"

var (
	MysqlDSN   = "northmeter:Bcdz001?@tcp(192.168.2.5:3309)"
	DBName     = "iot-platform"
	Jwtkey     = "iot-platform"
	EmqxAddr   = "http://192.168.2.5:18083/api/v5"
	EmqxKey    = "f433e2a39e1be5be"
	EmqxSecret = "hB9BlNtP7moCilWn1zbMU0bbw8tcpTDhBMPILpmO1CkL"
)

type UserClaim struct {
	Id       uint   `json:"id"`
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}
