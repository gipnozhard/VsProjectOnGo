package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(filename string) (map[string]struct{}, error) {
	lines := make(map[string]struct{})
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines[scanner.Text()] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

/*func readerFile(filename io.Reader) {
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(filename)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())
	fmt.Println(len(wr.Bytes()))

}*/

/*func errors(filename *os.File, err error) {
	if err != nil {
		panic(err)
	}
	defer func(filename *os.File) {
		err := filename.Close()
		if err != nil {

		}
	}(filename)
}*/

func main() {
	/*var wg sync.WaitGroup
	  wg.Add(3)*/

	/*file, _ := os.Open("/home/pavel/Рабочий стол/num.md")
	  file2, _ := os.Open("/home/pavel/Рабочий стол/num1.md")*/

	file := "files/num.md"
	file2 := "files/num1.md"

	/*defer file.Close()
	  defer file2.Close()

	  go func() {
	  	defer wg.Done()
	  	readerFile(file)
	  }()
	  go func() {
	  	defer wg.Done()
	  	readerFile(file2)
	  }()*/

	lines1, err := readLines(file)
	if err != nil {
		fmt.Println("Error reading file 1:", err)
		return
	}

	lines2, err := readLines(file2)
	if err != nil {
		fmt.Println("Error reading file 2:", err)
		return
	}

	fmt.Println("Всего элементв в первом файле: ", len(lines1))
	fmt.Println("Всего элементв во втором файле: ", len(lines2))

	if len(lines1) == 0 && len(lines2) == 0 {
		fmt.Println("Нет данных в обоих файлах")
		return
	} else if len(lines1) == 0 {
		fmt.Println("Нет данных в первом файле")
		return
	} else if len(lines2) == 0 {
		fmt.Println("Нет данных во втором файле")
		return
	}

	// Проверяем уникальные строки из первого файла
	fmt.Println("\nСтроки из первого файла, которых нет во втором:")
	foundDifferences := false
	for line := range lines1 {
		if _, found := lines2[line]; !found {
			fmt.Println(line)
			foundDifferences = true
		}
	}

	if !foundDifferences {
		fmt.Println("Нет таких!)")
	}

	// Проверяем уникальные строки из второго файла
	fmt.Println("\nСтроки из второго файла, которых нет в первом:")
	foundDifferences = false
	for line := range lines2 {
		if _, found := lines1[line]; !found {
			fmt.Println(line)
			foundDifferences = true
		}
	}

	if !foundDifferences {
		fmt.Println("Нет таких!)")
	}

	//wg.Wait()
}
