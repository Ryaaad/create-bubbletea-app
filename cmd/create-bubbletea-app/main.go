package main

import (
	"fmt"
	"os"
	"text/template"
)

type Content struct {
	Imports string
	View    string
}

func createFile(path string, tmpl *template.Template, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating the file %s: %w", path, err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing template for %s: %w", path, err)
	}

	return nil
}

func validateParams() ([]string, string) {
	Args := os.Args[1:]
	params := []string{"Bibble_tea_app", ""}
	ArgsErr := ""
	if len(Args) >= 1 {
		if Args[0] != " " {
			params[0] = Args[0]
		}
	}
	if len(Args) == 2 {
		if Args[1] != "--lipgloss" {
			ArgsErr = fmt.Sprint("unrecognized option '", Args[1], "'")
			panic(ArgsErr)
		} else {
			params[1] = "--lipgloss"
		}
	}
	return params, ArgsErr
}
func CreateProjet() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	params, ArgsErr := validateParams()
	if ArgsErr == "" {
		path := cwd
		path += "/" + params[0]
		pathError := os.Mkdir(path, os.ModePerm)
		if pathError != nil {
			return fmt.Errorf("err creating the directory : %w", pathError)
		} else {
			go_main_Content :=
				`package main

{{.Imports}}

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
		
{{.View}}
		
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
}`
			go_module_content :=
				`module bubbleteaapp

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
				`github.com/charmbracelet/bubbletea v0.26.4 h1:2gDkkzLZaTjMl/dQBpNVtnvcCxsh/FCkimep7FC9c40=
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
			Nostyle := Content{
				`import (
				"fmt"
				"os"
			
				tea "github.com/charmbracelet/bubbletea"
			)`,
				`func (m model) View() string {
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
			`}

			if params[1] == "--lipgloss" {
				Nostyle = Content{
					`import (
					"fmt"
					"os"
					tea "github.com/charmbracelet/bubbletea"
					"github.com/charmbracelet/lipgloss"
				)`,
					`func (m model) View() string {
						// Style
						var Contentstyle = lipgloss.NewStyle().
							Border(lipgloss.RoundedBorder()).
							Width(15)
					
						var Exitstyle = lipgloss.NewStyle().
							Bold(true).
							Foreground(lipgloss.Color("9"))
						var SelectedColor = lipgloss.NewStyle().
							Foreground(lipgloss.Color("10"))
						s := "\n"
					
						// Iterate over our choices
						for i := range m.board {
					
							cursorleft := " "
							cursorright := " "
							if m.cursor == i {
								cursorleft = ">"
								cursorright = "<"
							}
							s += fmt.Sprintf("\t%s %s %s", SelectedColor.Render(cursorleft), m.board[i], SelectedColor.Render(cursorright))
							s += "\n"
						}
						s = Contentstyle.Render(s)
						if m.Selected != "" {
							s += fmt.Sprintf("\nYou Selected %s \n", SelectedColor.Render(m.Selected))
						}
						// The footer
						s += Exitstyle.Render("\nPress q to quit.\n")
					
						// Send the UI for rendering
						return s
					}
				`}
				go_module_content = `module bubbleteaapp

go 1.22.3

require github.com/charmbracelet/bubbletea v0.26.4

require (
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
)

require (
	github.com/charmbracelet/lipgloss v0.11.0
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
)	`
				go_sum_content = `github.com/aymanbagabas/go-osc52/v2 v2.0.1 h1:HwpRHbFMcZLEVr42D4p7XBqjyuxQH5SMiErDT4WkJ2k=
github.com/aymanbagabas/go-osc52/v2 v2.0.1/go.mod h1:uYgXzlJ7ZpABp8OJ+exZzJJhRNQ2ASbcXHWsFqH8hp8=
github.com/charmbracelet/bubbletea v0.26.4 h1:2gDkkzLZaTjMl/dQBpNVtnvcCxsh/FCkimep7FC9c40=
github.com/charmbracelet/bubbletea v0.26.4/go.mod h1:P+r+RRA5qtI1DOHNFn0otoNwB4rn+zNAzSj/EXz6xU0=
github.com/charmbracelet/lipgloss v0.11.0 h1:UoAcbQ6Qml8hDwSWs0Y1cB5TEQuZkDPH/ZqwWWYTG4g=
github.com/charmbracelet/lipgloss v0.11.0/go.mod h1:1UdRTH9gYgpcdNN5oBtjbu/IzNKtzVtb7sqN1t9LNn8=
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
github.com/lucasb-eyer/go-colorful v1.2.0 h1:1nnpGOrhyZZuNyfu1QjKiUICQ74+3FNCN69Aj6K7nkY=
github.com/lucasb-eyer/go-colorful v1.2.0/go.mod h1:R4dSotOR9KMtayYi1e77YzuveK+i7ruzyGqttikkLy0=
github.com/mattn/go-isatty v0.0.20 h1:xfD0iDuEKnDkl03q4limB+vH+GxLEtL/jb4xVJSWWEY=
github.com/mattn/go-isatty v0.0.20/go.mod h1:W+V8PltTTMOvKvAeJH7IuucS94S2C6jfK/D7dTCTo3Y=
github.com/mattn/go-localereader v0.0.1 h1:ygSAOl7ZXTx4RdPYinUpg6W99U8jWvWi9Ye2JC/oIi4=
github.com/mattn/go-localereader v0.0.1/go.mod h1:8fBrzywKY7BI3czFoHkuzRoWE9C+EiG4R1k4Cjx5p88=
github.com/mattn/go-runewidth v0.0.15 h1:UNAjwbU9l54TA3KzvqLGxwWjHmMgBUVhBiTjelZgg3U=
github.com/mattn/go-runewidth v0.0.15/go.mod h1:Jdepj2loyihRzMpdS35Xk/zdY8IAYHsh153qUoGf23w=
github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 h1:ZK8zHtRHOkbHy6Mmr5D264iyp3TiX5OmNcI5cIARiQI=
github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6/go.mod h1:CJlz5H+gyd6CUWT45Oy4q24RdLyn7Md9Vj2/ldJBSIo=
github.com/muesli/cancelreader v0.2.2 h1:3I4Kt4BQjOR54NavqnDogx/MIoWBFa0StPA8ELUXHmA=
github.com/muesli/cancelreader v0.2.2/go.mod h1:3XuTXfFS2VjM+HTLZY9Ak0l6eUKfijIfMUZ4EgX0QYo=
github.com/muesli/termenv v0.15.2 h1:GohcuySI0QmI3wN8Ok9PtKGkgkFIk7y6Vpb5PvrY+Wo=
github.com/muesli/termenv v0.15.2/go.mod h1:Epx+iuz8sNs7mNKhxzH4fWXGNpZwUaJKRS1noLXviQ8=
github.com/rivo/uniseg v0.2.0/go.mod h1:J6wj4VEh+S6ZtnVlnTBMWIodfgj8LQOQFoIToxlJtxc=
github.com/rivo/uniseg v0.4.7 h1:WUdvkW8uEhrYfLC4ZzdpI2ztxP1I582+49Oc5Mq64VQ=
github.com/rivo/uniseg v0.4.7/go.mod h1:FN3SvrM+Zdj16jyLfmOkMNblXMcoc8DfTHruCPUcx88=
github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e h1:JVG44RsyaB9T2KIHavMF/ppJZNG9ZpyihvCd0w101no=
github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e/go.mod h1:RbqR21r5mrJuqunuUZ/Dhy/avygyECGrLceyNeo4LiM=
golang.org/x/sync v0.7.0 h1:YsImfSBoP9QPYL0xyKJPq0gcaJdG3rInoqxTWbfQu9M=
golang.org/x/sync v0.7.0/go.mod h1:Czt+wKu1gCyEFDUtn0jG5QVvpJ6rzVqr5aXyt9drQfk=
golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.6.0/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.20.0 h1:Od9JTbYCk261bKm4M/mw7AklTlFYIa0bIp9BgSm1S8Y=
golang.org/x/sys v0.20.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
golang.org/x/text v0.3.8 h1:nAL+RVCQ9uMn3vJZbV+MRnydTJFPf8qqY42YiA6MrqY=
golang.org/x/text v0.3.8/go.mod h1:E6s5w1FMmriuDzIBO73fBruAKo1PCIq6d2Q6DHfQ8WQ=
			`
			}

			mainGOtmplt, err := template.New("main.go").Parse(go_main_Content)
			if err != nil {
				return fmt.Errorf("error creating the main go template: %w", err)
			}
			goModtmplt, err := template.New("go.mod").Parse(go_module_content)
			if err != nil {
				return fmt.Errorf("error creating the go.mod template: %w", err)
			}
			goSumtmplt, err := template.New("go.sum").Parse(go_sum_content)
			if err != nil {
				return fmt.Errorf("error creating the sum go template: %w", err)
			}
			maingoPath := path + "/main.go"
			Err := createFile(maingoPath, mainGOtmplt, Nostyle)
			if Err != nil {
				return fmt.Errorf("error creating the main go file: %w", Err)

			}
			modgoPath := path + "/go.mod"
			Err = createFile(modgoPath, goModtmplt, nil)

			if Err != nil {
				return fmt.Errorf("error creating the mod go file: %w", Err)
			}
			gosumPath := path + "/go.sum"
			Err = createFile(gosumPath, goSumtmplt, nil)
			if Err != nil {
				return fmt.Errorf("error creating the sum go file: %w", Err)
			}
		}
	}
	return nil
}
func main() {
	if err := CreateProjet(); err != nil {
		fmt.Println("Error creating project:", err)
	}
}
