package main

import (
	"fmt"
	"os"
	"text/template"
)

type ChangedContent struct {
	View string
}

func Create() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}
	path := cwd
	path += "/Bubble"

	PathError := os.Mkdir(path, os.ModePerm)
	if PathError != nil {
		fmt.Println("Err creating the directory :", PathError.Error())
	} else {
		go_main_Content :=
			`
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	board    []string
	cursor   int
	Selected string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		
		case tea.KeyMsg:
	
			switch msg.String() {
	
			case "ctrl+c", "q":
				return m, tea.Quit
	
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
	
			case "down", "j":
				if m.cursor < len(m.board)-1 {
					m.cursor++
				}
	
			case "enter", " ":
				m.Selected = m.board[m.cursor]
				}
			}
			return m, nil
		}
		
func (m model) View() string {
	s := "\n"
			// Iterate over our choices
			for i := range m.board {
		
				cursor := " "
				if m.cursor == i {
					cursor = ">"
				}
				s += fmt.Sprintf("%s %s", cursor, m.board[i])
				s += "\n"
			}
			if m.Selected != "" {
				s += fmt.Sprintf("\nYou Selected %s \n", m.Selected)
			}
			// The footer
			s += "\nPress q to quit.\n"
		
			// Send the UI for rendering
			return s
		}
		
		func initialModel() model {
			return model{
				board:  []string{"A", "B", "C"},
				cursor: 0,
			}
		}
		
		func main() {
			p := tea.NewProgram(initialModel())
			if _, err := p.Run(); err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}
			fmt.Print("{.view}")
		}		
		
`
		go_module_content :=
			`
module Ryaaad

go 1.22.3

require github.com/charmbracelet/bubbletea v0.26.4

require (
	github.com/charmbracelet/x/ansi v0.1.2 // indirect
	github.com/charmbracelet/x/input v0.1.0 // indirect
	github.com/charmbracelet/x/term v0.1.1 // indirect
	github.com/charmbracelet/x/windows v0.1.0 // indirect
	github.com/erikgeiser/coninput v0.0.0-20211004153227-1c3628e74d0f // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.3.8 // indirect
)

`
		go_sum_content :=
			`
github.com/charmbracelet/bubbletea v0.26.4 h1:2gDkkzLZaTjMl/dQBpNVtnvcCxsh/FCkimep7FC9c40=
github.com/charmbracelet/bubbletea v0.26.4/go.mod h1:P+r+RRA5qtI1DOHNFn0otoNwB4rn+zNAzSj/EXz6xU0=
github.com/charmbracelet/x/ansi v0.1.2 h1:6+LR39uG8DE6zAmbu023YlqjJHkYXDF1z36ZwzO4xZY=
github.com/charmbracelet/x/ansi v0.1.2/go.mod h1:dk73KoMTT5AX5BsX0KrqhsTqAnhZZoCBjs7dGWp4Ktw=
github.com/charmbracelet/x/input v0.1.0 h1:TEsGSfZYQyOtp+STIjyBq6tpRaorH0qpwZUj8DavAhQ=
github.com/charmbracelet/x/input v0.1.0/go.mod h1:ZZwaBxPF7IG8gWWzPUVqHEtWhc1+HXJPNuerJGRGZ28=
github.com/charmbracelet/x/term v0.1.1 h1:3cosVAiPOig+EV4X9U+3LDgtwwAoEzJjNdwbXDjF6yI=
github.com/charmbracelet/x/term v0.1.1/go.mod h1:wB1fHt5ECsu3mXYusyzcngVWWlu1KKUmmLhfgr/Flxw=
github.com/charmbracelet/x/windows v0.1.0 h1:gTaxdvzDM5oMa/I2ZNF7wN78X/atWemG9Wph7Ika2k4=
github.com/charmbracelet/x/windows v0.1.0/go.mod h1:GLEO/l+lizvFDBPLIOk+49gdX49L9YWMB5t+DZd0jkQ=
github.com/erikgeiser/coninput v0.0.0-20211004153227-1c3628e74d0f h1:Y/CXytFA4m6baUTXGLOoWe4PQhGxaX0KpnayAqC48p4=
github.com/erikgeiser/coninput v0.0.0-20211004153227-1c3628e74d0f/go.mod h1:vw97MGsxSvLiUE2X8qFplwetxpGLQrlU1Q9AUEIzCaM=
github.com/mattn/go-localereader v0.0.1 h1:ygSAOl7ZXTx4RdPYinUpg6W99U8jWvWi9Ye2JC/oIi4=
github.com/mattn/go-localereader v0.0.1/go.mod h1:8fBrzywKY7BI3czFoHkuzRoWE9C+EiG4R1k4Cjx5p88=
github.com/mattn/go-runewidth v0.0.15 h1:UNAjwbU9l54TA3KzvqLGxwWjHmMgBUVhBiTjelZgg3U=
github.com/mattn/go-runewidth v0.0.15/go.mod h1:Jdepj2loyihRzMpdS35Xk/zdY8IAYHsh153qUoGf23w=
github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 h1:ZK8zHtRHOkbHy6Mmr5D264iyp3TiX5OmNcI5cIARiQI=
github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6/go.mod h1:CJlz5H+gyd6CUWT45Oy4q24RdLyn7Md9Vj2/ldJBSIo=
github.com/muesli/cancelreader v0.2.2 h1:3I4Kt4BQjOR54NavqnDogx/MIoWBFa0StPA8ELUXHmA=
github.com/muesli/cancelreader v0.2.2/go.mod h1:3XuTXfFS2VjM+HTLZY9Ak0l6eUKfijIfMUZ4EgX0QYo=
github.com/rivo/uniseg v0.2.0/go.mod h1:J6wj4VEh+S6ZtnVlnTBMWIodfgj8LQOQFoIToxlJtxc=
github.com/rivo/uniseg v0.4.7 h1:WUdvkW8uEhrYfLC4ZzdpI2ztxP1I582+49Oc5Mq64VQ=
github.com/rivo/uniseg v0.4.7/go.mod h1:FN3SvrM+Zdj16jyLfmOkMNblXMcoc8DfTHruCPUcx88=
github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e h1:JVG44RsyaB9T2KIHavMF/ppJZNG9ZpyihvCd0w101no=
github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e/go.mod h1:RbqR21r5mrJuqunuUZ/Dhy/avygyECGrLceyNeo4LiM=
golang.org/x/sync v0.7.0 h1:YsImfSBoP9QPYL0xyKJPq0gcaJdG3rInoqxTWbfQu9M=
golang.org/x/sync v0.7.0/go.mod h1:Czt+wKu1gCyEFDUtn0jG5QVvpJ6rzVqr5aXyt9drQfk=
golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.20.0 h1:Od9JTbYCk261bKm4M/mw7AklTlFYIa0bIp9BgSm1S8Y=
golang.org/x/sys v0.20.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
golang.org/x/text v0.3.8 h1:nAL+RVCQ9uMn3vJZbV+MRnydTJFPf8qqY42YiA6MrqY=
golang.org/x/text v0.3.8/go.mod h1:E6s5w1FMmriuDzIBO73fBruAKo1PCIq6d2Q6DHfQ8WQ=

`
		Nostyle := ChangedContent{"Style"}
		mainGOtmplt, err := template.New("main.go").Parse(go_main_Content)
		if err != nil {
			panic(err)
		}
		goModtmplt, err := template.New("go.mod").Parse(go_module_content)
		if err != nil {
			panic(err)
		}
		goSumtmplt, err := template.New("go.sum").Parse(go_sum_content)
		if err != nil {
			panic(err)
		}
		maingoPath := path + "/main.go"
		mainfile, Error := os.Create(maingoPath)
		if Error != nil {
			fmt.Print("Err creating the file : ", Error)
		} else {
			err = mainGOtmplt.Execute(mainfile, Nostyle)
			if err != nil {
				panic(err)
			}
		}
		modgoPath := path + "/go.mod"
		modfile, Error := os.Create(modgoPath)
		if Error != nil {
			fmt.Print("Err creating the file : ", Error)
		} else {
			err = goModtmplt.Execute(modfile, nil)
			if err != nil {
				panic(err)
			}
		}
		gosumPath := path + "/go.sum"
		sumfile, Error := os.Create(gosumPath)
		if Error != nil {
			fmt.Print("Err creating the file : ", Error)
		} else {
			err = goSumtmplt.Execute(sumfile, nil)
			if err != nil {
				panic(err)
			}
		}
	}
}
func main() {
	Create()
}
