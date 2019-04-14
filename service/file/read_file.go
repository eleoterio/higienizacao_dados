package file

import (
	"bufio"
	"log"
	"os"
	"strings"
)

//ReadFile Leitura do arquivo e preparação para inserir na tabela ticket
func ReadFile(fileDir string) ([][]string, error) {

	file, err := os.Open(fileDir)
	if err != nil {
		log.Println("[ReadFile] Erro na abertura do arquivo: ", err.Error())
		return nil, err
	}
	defer file.Close()

	text := bufio.NewScanner(file)
	text.Scan()

	var list [][]string

	for text.Scan() {
		var columns = strings.Fields(text.Text())
		list = append(list, columns)
	}

	return list, text.Err()
}
