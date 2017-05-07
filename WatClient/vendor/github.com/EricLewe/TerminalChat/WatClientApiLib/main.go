package WatClientApiLib

import (
	"fmt"
	//uiUtil "github.com/gizak/termui"
    "github.com/gizak/termui"
    pb "github.com/EricLewe/TerminalChat/WatApi"
    //"google.golang.org/grpc/grpclog"
    //pb "myProjects/WatApi"
    "time"
    "sync"
)

var (
    cCollection ControlCollection
)
type ControlCollection struct {
    SelectedList []*pb.ConversationReply
    SelectedConversation []*pb.ChatMessageReply
    ChatHasFocus bool
    WeatherHasFocus bool
    WeatherData pb.WeatherReply
    mux sync.Mutex
}

func (cC *ControlCollection) Update(SelectedList []*pb.ConversationReply, SelectedConversation []*pb.ChatMessageReply) {
    cC.mux.Lock()

    cC.SelectedList = SelectedList
    cC.SelectedConversation = SelectedConversation

    cC.mux.Unlock()
    return
}

func (cC *ControlCollection) UpdateWeather(weatherData pb.WeatherReply) {
    cC.mux.Lock()

    cC.WeatherData = weatherData

    cC.mux.Unlock()
    return
}

func (cC *ControlCollection) SetChatFocus(ChatHasFocus bool) {
    cC.mux.Lock()

    cC.ChatHasFocus = ChatHasFocus

    defer cC.mux.Unlock()
    return
}

func (cC *ControlCollection) SetWeatherFocus(WeatherHasFocus bool) {
    cC.mux.Lock()

    cC.WeatherHasFocus = WeatherHasFocus

    defer cC.mux.Unlock()
    return
}

func (cC *ControlCollection) getSelectedList() []*pb.ConversationReply {
    cC.mux.Lock()
    defer cC.mux.Unlock()
    return cC.SelectedList
}

func (cC *ControlCollection) getSelectedConversation() []*pb.ChatMessageReply {
    cC.mux.Lock()
    defer cC.mux.Unlock()
    return cC.SelectedConversation
}

func (cC *ControlCollection) getChatHasFocus() bool {
    cC.mux.Lock()
    defer cC.mux.Unlock()
    return cC.ChatHasFocus
}

func NewControlCollection() (*ControlCollection) {
    cCollection := new(ControlCollection)
    cCollection.SelectedList = make([]*pb.ConversationReply, 0)
    cCollection.ChatHasFocus = false;
    return cCollection
}

func InitWindow() {
    err := termui.Init()
    if err != nil {
	panic(err)
    }
    defer termui.Close()

    par1 := termui.NewPar("Welcome!")
    par1.Height = 5
    par1.Width = 40
    par1.Y = 10
    termui.Render(par1)
    // handle key q pressing
    termui.Handle("/sys/kbd/q", func(termui.Event) {
	fmt.Printf("stoping!")
	// press q to quit
	termui.StopLoop()

    })

    termui.Loop()
}

func RenderWeatherInfo(weatherData pb.WeatherReply) {
    par1 := termui.NewPar(weatherData.Broadcast)
    par1.Height = 5
    par1.Width = 40
    par1.Y = 10
    par1.BorderLabel = "The weather today is: "
    termui.Render(par1)
    par2 := termui.NewPar(weatherData.Description)
    par2.Height = 5
    par2.Width = 40
    par2.X = 40 + 2
    par2.Y = 10
    par2.BorderLabel = "That is: "
    termui.Render(par2)
}

func RenderTerminal(asd string) string {
    terminalInputPar := termui.NewPar(asd)
    terminalInputPar.Height = 3
    terminalInputPar.Width = termui.TermWidth()
    terminalInputPar.SetY(termui.TermHeight() - terminalInputPar.Height)
    termui.Render(terminalInputPar)

    return ""
}

func RenderMessages(chatmessages []*pb.ChatMessageReply, offset int) {
	conversaionList := make([]termui.Par, 0)
	for i, chatmessage := range chatmessages {
	    par1 := termui.NewPar(fmt.Sprintf("Skrev: %s", chatmessage.Content))
	    par1.Height = 3
	    par1.Width = 60
	    par1.X = 60 + 5
	    par1.Y = 3 * i + offset
	    par1.BorderLabel = chatmessage.SentByUser
	    conversaionList = append(conversaionList, *par1)
	    termui.Render(par1)
	}


}

func RenderConversations(conversations []*pb.ConversationReply, offset int) {
    conversaionList := make([]termui.Par, 0)
    for i, conversation := range conversations {
	latestmessageContent := ""
	if conversation.LatestMessage != nil {
	    latestmessageContent = conversation.LatestMessage.Content
	}
	par1 := termui.NewPar(fmt.Sprintf("%d - %s", conversation.Id, latestmessageContent))
	par1.Height = 3
	par1.Width = 60
	par1.Y = 3 * i + offset
	par1.BorderLabel = conversation.Name
	conversaionList = append(conversaionList, *par1)
	termui.Render(par1)
    }


}

func (cC *ControlCollection) MessagePipeline(out chan<- string) {
    eventQueue := make(chan termui.Event)
    termui.Handle("/sys/kbd", func(event termui.Event) {
	eventQueue <- event
    })

    //mode := 0
    a := 0
    b := 0
    terminalInput := ""
    for {
	select {
	case ev := <-eventQueue:
	    offset := len("/sys/kbd/")
	    char := ev.Path[offset:]
	    switch {
	    	case ev.Path == "/sys/kbd/<up>":
		    if cC.ChatHasFocus {
			b = b - 3
		    } else {
			a = a - 3
		    }
		case ev.Path == "/sys/kbd/<down>":
		    if cC.ChatHasFocus {
			b = b + 3
		    } else {
			a = a + 3
		    }
	        case len(char) == 1:
		    terminalInput = fmt.Sprintf("%s%s", terminalInput, char)
	    	case ev.Path == "/sys/kbd/<enter>":
		    out <- terminalInput
		    terminalInput = ""
	    	case ev.Path == "/sys/kbd/<space>":
		    terminalInput = fmt.Sprintf("%s%s", terminalInput, " ")
	    case ev.Path == "/sys/kbd/<escape>":
		cC.ChatHasFocus = false
		cC.WeatherHasFocus = false
		b = 0
	    case ev.Path == "/sys/kbd/<c-8>":
		terminalInput = ""
	    }

	default:
	    if cC.WeatherHasFocus {
		termui.Clear()
		RenderWeatherInfo(cC.WeatherData)
		time.Sleep(time.Millisecond * 70)
	    } else {
		termui.Clear()
		RenderConversations(cC.SelectedList, a)
		if cC.ChatHasFocus {
		    RenderMessages(cC.SelectedConversation, b)
		}
		RenderTerminal(terminalInput)
		time.Sleep(time.Millisecond * 70)
	    }

	}
    }
}

/*func RenderSelectedListWithOffset() {
    for _, control := range cCollection.selectedList {
	par1 := termui.NewPar(fmt.Sprintf("%d - %s", conversation.Id, "Lorem Ipsum dolori! Asd"))
	par1.Height = 3
	par1.Width = 40
	par1.Y = 3 * i
	par1.BorderLabel = conversation.Name
	termui.Render(par1)
	conversaionList = append(conversaionList, *par1)
    }
}*/


//func InitWindow() {
//	g, err := gocui.NewGui(gocui.OutputNormal)
//	if err != nil {
//		log.Panicln(err)
//	}
//	defer g.Close()
//
//	g.SetManagerFunc(layout)
//
//	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
//		log.Panicln(err)
//	}
//
//	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
//		log.Panicln(err)
//	}
//}
//
//func layout(g *gocui.Gui) error {
//	maxX, maxY := g.Size()
//	if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
//		if err != gocui.ErrUnknownView {
//			return err
//		}
//		fmt.Fprintln(v, "Hello world!")
//	}
//	return nil
//}
//
//func quit(g *gocui.Gui, v *gocui.View) error {
//	return gocui.ErrQuit
//}