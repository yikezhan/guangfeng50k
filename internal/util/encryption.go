package util

import "github.com/golang-module/dongle"

var MacMd5 string

var cipher = dongle.NewCipher()

func init() {
	MacMd5 = "ywq"

	cipher.SetMode(dongle.CBC)        // CBC、ECB、CFB、OFB、CTR、GCM
	cipher.SetPadding(dongle.PKCS7)   // No、Zero、PKCS5、PKCS7
	cipher.SetKey("1234567887654321") // key 长度必须是 16、24 或 32
	cipher.SetIV("1234567887654321")
}

func EncryptHmacMd5(s string) string {
	return dongle.Encrypt.FromString(s).ByHmacMd5(MacMd5).ToHexString()
}
func Md5(s string) string {
	return dongle.Encrypt.FromString(s).ByMd5().ToHexString()
}
func EncryptAES(s string) string {
	// 对字符串进行 aes 加密，输出经过 hex 编码的字符串
	return dongle.Encrypt.FromString(s).ByAes(cipher).ToHexString()
}
func DecryptAES(s string) string {
	// 对经过 hex 编码的字符串进行 aes 解密，输出字符串
	return dongle.Decrypt.FromHexString(s).ByAes(cipher).ToString()
}
