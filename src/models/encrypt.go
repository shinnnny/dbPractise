package models

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

//填充函数，填充明文达到64位
func Fill(cle []byte, blocksize int)[]byte{
	fillsize := blocksize - len(cle)%blocksize
	repeat := bytes.Repeat([]byte{0}, fillsize)
	return append(cle,  repeat...)
}

func Out(cle []byte)[]byte{
	return bytes.TrimRightFunc(cle, func(r rune) bool {
		return r ==  rune(0)
	})
}

//加密函数
func DesEncrypt(cle, key []byte)[]byte{
	block,  e :=  des.NewCipher(key)
	if e != nil{
		fmt.Print(e)
	}

	fillcle := Fill(cle, block.BlockSize())

	//制作用于接收的密文切片
	cip :=make([]byte, len(fillcle))

	//使用CBC模式加密
	encrypter := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	//使用CBC模式加密，赋值给cip
	encrypter.CryptBlocks(cip, fillcle)
	return  cip
}

//解密函数
func DesDecrypt(cip, key []byte)[]byte{
	block, e := des.NewCipher(key)
	if e !=  nil{
		fmt.Print(e)
	}

	decrypter := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	cle := make([]byte, len(cip))
	decrypter.CryptBlocks(cle, cip)
	out := Out(cle)
	return out
}