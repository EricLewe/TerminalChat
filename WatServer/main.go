/*
 *
 * Copyright 2015, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package main

import (
    "log"
    "net"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/grpc/credentials"
    "io/ioutil"
    "google.golang.org/grpc/grpclog"
    google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
    "encoding/json"
    "flag"
    "time"
    "sync"
    "fmt"
    "math/rand"
    pb "github.com/EricLewe/TerminalChat/WatApi"
    watWBot "github.com/EricLewe/TerminalChat/WatWeatherBot"
    "google.golang.org/grpc/peer"
)

const (
    port = ":50051"
)

var (
    jsonDBFile = flag.String("json_db_file", "WatApi/data.json", "A json file containing a list of messages")
    jsonUsersFile = flag.String("json_Users_File", "WatApi/users.json", "A json file containing a list of users")
)

// server is used to implement Wat.ChatServer.
type ChatServer struct{
    savedMessages []*pb.ChatMessageReply
    savedConversations []*pb.ConversationReply
    savedUsers map[string]*pb.LoginRequest
    pipedMessages map[string][]*pb.ChatMessageReply
    subscribers map[int32][]string
    mux sync.Mutex
}

type User struct {
    username string
    password string
}
func newServer() *ChatServer {
    s := new(ChatServer)
    s.savedConversations = []*pb.ConversationReply{}
    s.pipedMessages = make(map[string][]*pb.ChatMessageReply)
    s.subscribers = make(map[int32][]string)
    s.savedUsers = make(map[string]*pb.LoginRequest)
    s.loadUsers(*jsonUsersFile)
    s.loadMessages(*jsonDBFile)
    return s
}

//returns all users who wants to get messages from a conversation
func (s *ChatServer) getSubscribers(id int32) []string {
    s.mux.Lock()
    defer s.mux.Unlock()
    return s.subscribers[id]
}

//returns adds users who wants to get messages from a conversation
func (s *ChatServer) addSubscribers(id int32, username string) {
    s.mux.Lock()

    if _, present := s.subscribers[id]; !present {
	s.subscribers[id] = []string{username}
    } else {
	s.subscribers[id] = append(s.subscribers[id], username)
    }
    defer s.mux.Unlock()
    return
}

//pops all pending messages from a user
func (s *ChatServer) getAndEmptyMessage(username string) []*pb.ChatMessageReply {
    s.mux.Lock()
    a := s.pipedMessages[username]
    delete(s.pipedMessages, username)
    defer s.mux.Unlock()
    return a
}

//adds a pending message to a user
func (s *ChatServer) addMessageToUser(username string, chatMessageReply pb.ChatMessageReply) {
    s.mux.Lock()
    if _, present := s.pipedMessages[username]; !present {
	s.pipedMessages[username] = []*pb.ChatMessageReply{&chatMessageReply}
    } else {
	s.pipedMessages[username] = append(s.pipedMessages[username], &chatMessageReply)
    }
    defer s.mux.Unlock()
    return
}

//Post method, returns the clients weather based on the peer structs ip
func (s *ChatServer) GetWeather(ctx context.Context, in *pb.WeatherRequest) (*pb.WeatherReply, error) {
    peer, _ := peer.FromContext(ctx)
    broadcast, description := watWBot.GetCurrentWeather(peer.Addr.String())
    weatherReply := pb.WeatherReply{ broadcast, description}
    return &weatherReply, nil
}

//ensures username and password is correct when a user tries to connect
func (s *ChatServer) VerifyLogin(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
    loginReply := pb.LoginReply{ "", ""}
    if user, validUserName := s.savedUsers[in.Username]; validUserName {
	if validPassword := in.Password == user.Password; validPassword {
	    loginReply.Username = in.Username
	    loginReply.MessageOfTheDay = "Welcome online " + in.Username
	}
    }

    return &loginReply, nil
}

//Get method, fetches a users sent messages and delegates them to the subscribers
func (s *ChatServer) SendMessage(ctx context.Context, in *pb.ChatMessageReply) (*pb.Request, error) {
    //Pipe this msg into all related users
    for _, subscriber := range s.getSubscribers(in.ConversationId) {
	s.addMessageToUser(subscriber, *in)
    }

    return &pb.Request{}, nil
}

//Post method, sends conversations even those the user may not acccess to (should be fixed)
func (s *ChatServer) RouteConversation(request *pb.Request, stream pb.Chat_RouteConversationServer) error {
    for _, feature := range s.savedConversations {
	if err := stream.Send(feature); err != nil {
	    return err
	}
    }
    return nil
}

//Post method, sends messages to a user
func (s *ChatServer) RouteChat(conversation *pb.ConversationRequest, stream pb.Chat_RouteChatServer) error {
    //We only what messages with specific Id, currently O(n) in worst case
    if conversation.Id > 0 {
	for _, message := range s.savedMessages {
	    if message.ConversationId == conversation.Id {
		if err := stream.Send(message); err != nil {
		    return err
		}
	    }
	}
    } else {
	for _, feature := range s.getAndEmptyMessage(conversation.Request.Username) {
	    s.savedMessages = append(s.savedMessages, feature)
	    if err := stream.Send(feature); err != nil {
		return err
	    }
	}
    }
    return nil
}

// loadMessages loads messages from a JSON file into the server struct. (should be replace with PostgresSQL)
func (s *ChatServer) loadMessages(filePath string) {
    file, err := ioutil.ReadFile(filePath)
    if err != nil {
	grpclog.Fatalf("Failed to load default features: %v", err)
    }
    if err := json.Unmarshal(file, &s.savedMessages); err != nil {
	grpclog.Fatalf("Failed to load default features: %v", err)
    }

    for _, message := range s.savedMessages {
	for _, username := range s.getSubscribers(message.ConversationId) {
	    s.addMessageToUser(username, *message)
	}
    }

}
// loadUsers loads messages from a JSON file into the server struct, also generates
// fake data regarding conversations (should be replaced with a db)
func (s *ChatServer) loadUsers(filePath string) ([]*pb.LoginRequest) {
    file, err := ioutil.ReadFile(filePath)
    if err != nil {
	grpclog.Fatalf("Failed to load default features: %v", err)
    }
    var users []*pb.LoginRequest
    if err := json.Unmarshal(file, &users); err != nil {
	grpclog.Fatalf("Failed to load default features: %v", err)
    }

    for i := 0; i < len(users); i++ {
	_, present := s.savedUsers[users[i].Username];
	if !present {
	    s.savedUsers[users[i].Username] = &pb.LoginRequest{Username: users[i].Username, Password:users[i].Password}
	    for j := 0; j < rand.Intn(6); j++ {
		s.addSubscribers(int32((i +j+1)%6), users[i].Username)
	    }

	} else {
	    fmt.Errorf("User already exists %s ", users[i].Username)
	}
    }

    //Now we create some fake data, since no Postgres yet
    timeTemp := time.Now()
    timestamp := google_protobuf.Timestamp{  int64(timeTemp.Second()), int32(timeTemp.Nanosecond())}
    for i := 0; i < 16; i++ {
	convId := int32((i))
	slice := s.subscribers[convId]
	conversationName := serialize(slice)

	features := &pb.ConversationReply{ convId,&timestamp, conversationName, &pb.ChatMessageReply{convId, "Lorem Ipsum", &timestamp, "lamacoder"}}
	s.savedConversations = append(s.savedConversations, features)
    }
    return users;
}

//serialize concatenates usernames who are in same conversation
func serialize(usernames []string) string {
    title := ""
    for i := 0; i < len(usernames); i++ {
	if i == 0 {
	    title = usernames[i]
	} else {
	    title = title + ", " + usernames[i]
	}
    }
    return title
}

//starts the server
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
    creds, err := credentials.NewServerTLSFromFile("WatApi/server.pem", "WatApi/server.key")
    var opts []grpc.ServerOption
    opts = []grpc.ServerOption{grpc.Creds(creds)}
	s := grpc.NewServer(opts...)
	pb.RegisterChatServer(s, newServer())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
