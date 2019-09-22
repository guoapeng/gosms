package gosms

type Credential struct {
	AccessKeyId  string
	AccessSecret string
}

type Profile struct {
	SignatureMethod string
	SignatureVersion string
	Credential Credential
}
