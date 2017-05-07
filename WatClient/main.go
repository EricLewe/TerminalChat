package main

import (
    "golang.org/x/net/context"
    "google.golang.org/grpc/grpclog"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc"
    "regexp"
    "strings"
    "strconv"
    "time"
    "os"
    "io"
    google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
    wcApi "github.com/EricLewe/TerminalChat/WatClientApiLib"
    pb "github.com/EricLewe/TerminalChat/WatApi"
)

const (
    address     = "localhost:50051"
    defaultName = "world"
)

type ChatClient struct {
    Username string
    conversationId int32
}

func newChatClient() *ChatClient {
    c := new(ChatClient)
    c.Username = ""
    c.conversationId = -1 //The current conversation the client subscribes on
    return c
}

func getWeather(client pb.ChatClient) pb.WeatherReply {
    weather, _ := client.GetWeather(context.Background(), &pb.WeatherRequest{})
    return *weather
}

func Login(client pb.ChatClient, username string, password string) (string, error) {
    feature, err := client.VerifyLogin(context.Background(), &pb.LoginRequest{username, password})
    if err != nil {
	grpclog.Fatalf("%v.VerifyLogin(_) = _, %v: ", client, err)
    }

    return feature.Username, nil
}

//deprecated, should be replaced by dbs constraints
func currentTime() (google_protobuf.Timestamp) {
    timeTemp := time.Now()
    timestamp := google_protobuf.Timestamp{  int64(timeTemp.Second()), int32(timeTemp.Nanosecond())}
    return timestamp
}

func SendMessageToServer(client pb.ChatClient, conversationId int32,message string, username string) {
    timestamp := currentTime();
    messageReply := &pb.ChatMessageReply{conversationId, message, &timestamp, username}
    client.SendMessage(context.Background(), messageReply)

}


func GetConversations(client pb.ChatClient, sessionId string) (conversations []*pb.ConversationReply) {

    conversation := &pb.Request{ sessionId}
    if sessionId == "" {
	return nil
    }
    stream, err := client.RouteConversation(context.Background(), conversation)
    if err != nil {
	grpclog.Fatalf("%v.getConversations(_) = _, %v: ", client, err)
    }
    for {
	conversationReply, err := stream.Recv()
	if err == io.EOF {
	    break
	}
	if err != nil {
	    grpclog.Fatalf("%v.getConversations(_) = _, %v", client, err)
	}
	conversations = append(conversations, conversationReply)
    }
    return conversations
}

//gets the messages sent by other clients from server, including itself
func GetMessagesFromClients(client pb.ChatClient, conversationId int32,sessionId string, cCollection wcApi.ControlCollection, selectedList []*pb.ConversationReply) ([]*pb.ConversationReply, []*pb.ChatMessageReply) {
    messages := GetMessagesFromConversation(client, -1, sessionId)

    selectedConversation := cCollection.SelectedConversation
    for _, message := range messages {
	if message.ConversationId == conversationId {
	    selectedConversation = append(selectedConversation, message)
	} else {
	    for i, _ := range selectedList {
		if selectedList[i].Id == message.ConversationId {
		    selectedList[i].LatestMessage.Content = message.Content
		}
	    }
	}
    }
    return selectedList, selectedConversation
}

func GetMessagesFromConversation(client pb.ChatClient, conversationId int32, sessionId string) (chatmessages []*pb.ChatMessageReply) {
    rect := &pb.ConversationRequest{conversationId, &pb.Request{sessionId}}
    stream, err := client.RouteChat(context.Background(), rect)
    if err != nil {
	grpclog.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
    }
    for {
	feature, err := stream.Recv()
	if err == io.EOF {
	    break
	}
	if err != nil {
	    grpclog.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	chatmessages = append(chatmessages, feature)
    }
    return chatmessages
}

//judges if the terminal may call the rendering of weather view or not.
func isInputWheather(in string) bool {
    var validID = regexp.MustCompile(`^!weather\s*$`)
    in = strings.ToLower(in)
    if validID.MatchString(in) {
	return true
    }
    return false
}

//judges if the terminal may join a conversation, (create conversation is not implemented yet)
func inputIsValid(in string) bool {
    var validID = regexp.MustCompile(`^((join)\s+([0-9]+$))|((create)\s+(\w+)(\s+(\w+))*)$`)
    in = strings.ToLower(in)
    if validID.MatchString(in) {
	return true
    }
    return false
}

func getConversationIdFromInput(in string) int32 {
    re := regexp.MustCompile("[0-9]+")
    id, _ := strconv.Atoi(re.FindString(in))
    return int32(id)
}

//runs a chatclient which provides data to the view.
func main() {
    creds, err := credentials.NewClientTLSFromFile("WatApi/server.pem", "Eric")
    if err != nil {
	grpclog.Fatalf("did not fix creds: %v", err)
    }
    conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
    if err != nil {
	grpclog.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewChatClient(conn)
    chatClient := newChatClient()

    username := "lamacoder"
    password := "useatownrisk"
    if len(os.Args) >= 2 {
	username = os.Args[0]
	password = os.Args[1]
    }

    for {
	chatClient.Username, _ = Login(c, username, password)

	if chatClient.Username != "" {
	    break
	}
    }


    go wcApi.InitWindow()

    cCollection := wcApi.NewControlCollection()
    cCollection.SelectedList = GetConversations(c, chatClient.Username)
    cCollection.SetChatFocus(false)
    ch := make(chan string, 1)

    go cCollection.MessagePipeline(ch)

    for {
	select {
	case terminalInput := <-ch:
	    if cCollection.ChatHasFocus {
		SendMessageToServer(c, chatClient.conversationId, terminalInput, chatClient.Username)
		cCollection.Update(GetMessagesFromClients(c, chatClient.conversationId, chatClient.Username, *cCollection, cCollection.SelectedList))
	    } else if isInputWheather(terminalInput) {
		cCollection.SetWeatherFocus(true)
	      	cCollection.UpdateWeather(getWeather(c))
	    } else if inputIsValid(terminalInput) {
		if (strings.HasPrefix(terminalInput,"join")) {
		    chatClient.conversationId = getConversationIdFromInput(terminalInput)
		    asd := GetMessagesFromConversation(c, chatClient.conversationId, chatClient.Username)
		    cCollection.Update(cCollection.SelectedList, asd)
		    cCollection.SetChatFocus(true)
		}
	    }
	default:
	    time.Sleep(time.Second * 1)
	}
	if len(cCollection.SelectedConversation) > 0 {
	    cCollection.Update(GetMessagesFromClients(c, chatClient.conversationId, chatClient.Username, *cCollection, cCollection.SelectedList))
	}
    }
}
