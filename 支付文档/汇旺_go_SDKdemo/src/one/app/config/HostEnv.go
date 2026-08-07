package config

type HostEnv string

const (
	Uat HostEnv = "UAT_HUI_ONE"
	Pro HostEnv = "PRO_TEST"
)

var Host = map[HostEnv]string{
	Uat: "https://sdk.oykqk.com/app/sdk-server/",
	Pro: "https://sdk.xone.la/app/sdk-server/",
}
