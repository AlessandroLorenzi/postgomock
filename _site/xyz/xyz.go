package xyz

type XYZ struct {
	keyId     string
	secretKey string
}

func New(keyId string, secretKey string) *XYZ {
	return &XYZ{keyId: keyId, secretKey: secretKey}
}

type GetInfoOutput struct {
	InfoINeed string
}

func (x *XYZ) GetInfo(_ int) (*GetInfoOutput, error) {
	return &GetInfoOutput{
		InfoINeed: "hallo",
	}, nil
}
