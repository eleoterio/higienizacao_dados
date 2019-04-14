package model

import (
	"strconv"
	"strings"
)

//ValidaCPF verifica se o CPF é valido ou nao
func ValidaCPF(CPF string) bool {

	if strings.Count(CPF, string(CPF[0])) == 11 {
		return false
	}

	var intCPF, intSoma, intMod int

	for i := 0; i < 9; i++ {
		intCPF, _ = strconv.Atoi(string(CPF[i]))
		intSoma = intSoma + intCPF*(10-i)
	}

	intMod = (intSoma * 10) % 11
	if intMod == 10 || intMod == 11 {
		intMod = 0
	}
	if intCPF, _ = strconv.Atoi(string(CPF[9])); intCPF != intMod {
		return false
	}

	intSoma = 0
	for i := 0; i < 10; i++ {
		intCPF, _ = strconv.Atoi(string(CPF[i]))
		intSoma = intSoma + intCPF*(11-i)
	}
	intMod = (intSoma * 10) % 11
	if intMod == 10 || intMod == 11 {
		intMod = 0
	}

	if intCPF, _ = strconv.Atoi(string(CPF[10])); intCPF != intMod {
		return false
	}

	return true
}

//ValidaCNPJ verifica se o CNPJ é valido ou nao
func ValidaCNPJ(CNPJ string) bool {

	if strings.Count(CNPJ, string(CNPJ[0])) == 14 {
		return false
	}

	var intCNPJ, intSoma, intMod int
	for i := 0; i < 4; i++ {
		intCNPJ, _ = strconv.Atoi(string(CNPJ[i]))
		intSoma = intSoma + intCNPJ*(5-i)
	}

	for i := 4; i < 12; i++ {
		intCNPJ, _ = strconv.Atoi(string(CNPJ[i]))
		intSoma = intSoma + intCNPJ*(13-i)
	}
	intMod = intSoma % 11
	if intMod < 2 {
		intMod = 0
	} else {
		intMod = 11 - intMod
	}

	if intCNPJ, _ = strconv.Atoi(string(CNPJ[12])); intCNPJ != intMod {
		return false
	}
	intSoma = 0

	for i := 0; i < 5; i++ {
		intCNPJ, _ = strconv.Atoi(string(CNPJ[i]))
		intSoma = intSoma + intCNPJ*(6-i)
	}

	for i := 5; i < 13; i++ {
		intCNPJ, _ = strconv.Atoi(string(CNPJ[i]))
		intSoma = intSoma + intCNPJ*(14-i)
	}
	intMod = intSoma % 11
	if intMod < 2 {
		intMod = 0
	} else {
		intMod = 11 - intMod
	}

	if intCNPJ, _ = strconv.Atoi(string(CNPJ[13])); intCNPJ != intMod {
		return false
	}

	return true
}

//ValidaNull Verifica so o campo é NULL
func ValidaNull(strDate string) bool {
	if strDate == "NULL" {
		return false
	}
	return true
}
