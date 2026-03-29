package data

type Tab struct {
	Title   string
	Content string
}

type Section struct {
	Title   string
	Color   string
	Tabs    []Tab
}

type Skill struct {
	Name       string
	LeftLabel  string
	RightLabel string
	Percentage float64
	Color      string
}

type Contact struct {
	Platform string
	Value    string
	Color    string
}

type ContentData struct {
	SidebarItems []string
	HeaderItems  []string
	Sections     map[string]Section
	Skills       []Skill
	Contacts     []Contact
	AsciiArt     string
}

func GetInitialData() ContentData {
	return ContentData{
		AsciiArt: `        ,.,
      MMMM_    ,..,
        "_ "__"MMMMM          ,...,,
 ,..., __." --"    ,.,     _-"MMMMMMM
MMMMMM"___ "_._   MMM"_."" _ """"""
 """""    "" , \_.   "_. ."
        ,., _"__ \__./ ."
       MMMMM_"  "_    ./
        ''''      (    )
 ._______________.-'____"---._.
  \                          /
   \________________________/
   (_)                    (_)`,
		Contacts: []Contact{
			{Platform: "MAIL", Value: "alex@term.shop", Color: "#00FF00"},
			{Platform: "LINK", Value: "https://www.linkedin.com/", Color: "#FFFFFF"},
			{Platform: "YT", Value: "https://www.youtube.com/@dimi244", Color: "#FF0000"},
		},
		Skills: []Skill{
			{Name: "RESILIENZA", LeftLabel: "nordvpn/tintoria", RightLabel: "selfhost smtp", Percentage: 0.7, Color: "#FFFF00"},
			{Name: "TEAMWORK", LeftLabel: "Beholder", RightLabel: "Leeeroooy jenkins", Percentage: 0.5, Color: "#FF69B4"},
			{Name: "NEGOZIAZIONE", LeftLabel: "plomo", RightLabel: "plata", Percentage: 0.8, Color: "#00FF00"},
		},
		SidebarItems: []string{"~ curriculum ~", "~ personal projects ~", "~ about me ~", "~ contact ~"},
		HeaderItems:  []string{"SSHowroom"},
		Sections: map[string]Section{
			"~ curriculum ~": {
				Title: "PROFESSIONAL CURRICULUM",
				Color: "#00d7ff",
				Tabs: []Tab{
					{
						Title: "CARRIERA",
						Content: `SENIOR GO DEVELOPER @ CLOUD-SYSTEMS (2021 - Present)
- Designed high-throughput microservices using gRPC and Kafka.
- Optimized database queries reducing latency by 40%.
- Mentored junior developers and led the Go standards committee.

GO BACKEND ENGINEER @ STREAM-FLOW (2018 - 2021)
- Developed real-time data processing pipelines in Go.
- Implemented automated testing suite with 90% coverage.
- Managed Kubernetes clusters and CI/CD pipelines.`,
					},
					{
						Title: "CERTIFICAZIONI",
						Content: `1. Gopher Certified Professional (GCP)
2. Advanced Concurrency in Go Certificate
3. Cloud-Native Go Architecture Master
4. Go Performance Tuning Specialist
5. Microservices with Go Specialist
6. Go Security & Defensive Programming
7. Full-Stack Gopher Associate`,
					},
				},
			},
			"~ personal projects ~": {
				Title: "PERSONAL PROJECTS",
				Color: "#afff00",
				Tabs: []Tab{
					{
						Title: "PROJECT 1",
						Content: `SSH-DASHBOARD
A real-time system monitoring dashboard accessible exclusively via SSH. 
Built using Wish and Bubble Tea to provide metrics without a web UI.`,
					},
					{
						Title: "PROJECT 2",
						Content: `GOPHER-GIT
A minimalist Git client implemented entirely in Go. 
Focuses on visual diffs and terminal-based branch management.`,
					},
					{
						Title: "PROJECT 3",
						Content: `CHIP-8 EMULATOR
A complete CHIP-8 emulator written in Go. 
Features a TUI-based display and supports loading custom ROMs.`,
					},
				},
			},
			"~ about me ~": {
				Title: "ABOUT ME",
				Color: "#af87ff",
				Tabs: []Tab{
					{
						Title: "PERSONAL",
						Content: `I am a terminal enthusiast who loves clean code and complex systems.
When I'm not coding, you can find me:
- Playing tactical RPGs and indie games.
- Exploring ambient and synthwave music.
- Designing weird TUIs for fun.`,
					},
				},
			},
			"~ contact ~": {
				Title: "CONNECT",
				Color: "#ff0087",
				Tabs: []Tab{
					{Title: "DIRECT", Content: "Email: alex@terminal.shop\nPGP: 0xABCD1234EF567890"},
					{Title: "SOCIAL", Content: "GitHub: github.com/alexcloudborn\nTwitter: @alexdev_cli"},
				},
			},
		},
	}
}
